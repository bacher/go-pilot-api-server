package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-errors/errors"
	"io/ioutil"
	"log"
	"miniteam/src/api/getInitialData"
	"miniteam/src/api/makeHappy"
	"net/http"
	"os"
	"regexp"
)

const PORT = ":8080"

var apiRx *regexp.Regexp

var ml *log.Logger = log.New(os.Stderr, "", 0)

func main() {
	apiRx = regexp.MustCompile(`^/api/([A-Za-z]+).json$`)

	Connect()
	Users.Load()

	srv := &http.Server{Addr: PORT}

	//go func() {
	//	time.Sleep(2 * time.Second)
	//
	//	if err := srv.Shutdown(nil); err != nil {
	//		panic(err)
	//	}
	//}()

	http.HandleFunc("/", handler)

	ml.Println("Server is starting at", PORT)

	err := srv.ListenAndServe()

	if err == nil || err == http.ErrServerClosed {
		ml.Println("Server closed")
	} else {
		ml.Printf("HttpServer listen error: %s\n", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	apiMatch := apiRx.FindStringSubmatch(r.URL.Path)

	if len(apiMatch) == 0 {
		w.WriteHeader(404)

		fmt.Fprint(w, "Sorry, bad api")
		return
	}

	defer func() {
		if err := recover(); err != nil {
			ml.Println("Recovered after api call fail:", errors.Wrap(err, 3).ErrorStack())
			returnFail(w)
		}
	}()

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		returnBadRequest(w)
		return
	}

	var params []*json.RawMessage

	if err := json.Unmarshal(reqBody, &params); err != nil {
		returnBadRequest(w)
		return
	}

	var sessionId string
	if err := json.Unmarshal(*params[0], &sessionId); err != nil {
		returnBadRequest(w)
		return
	}

	var data interface{}

	switch apiMatch[1] {
	case getInitialData.Name:
		p := &getInitialData.Params{}
		if !tt(params, p, w) {
			return
		}
		data = getInitialData.Do(p)
	case makeHappy.Name:
		p := &makeHappy.Params{}
		if !tt(params, p, w) {
			return
		}
		data = makeHappy.Do(p)
	}

	switch d := data.(type) {
	case []byte:
		w.Write(d)
	case string:
		fmt.Fprintln(w, d)
	default:
		bytes, err := json.Marshal(data)
		if err != nil {
			returnFail(w)
			return
		}
		w.Write(bytes)
	}
}

func returnBadRequest(w http.ResponseWriter) {
	w.WriteHeader(400)
	fmt.Fprintln(w, "Bad Request")
}

func returnFail(w http.ResponseWriter) {
	w.WriteHeader(500)
	fmt.Fprintln(w, "Internal Error")
}

func tt(data []*json.RawMessage, params interface{}, w http.ResponseWriter) bool {
	if len(data) >= 2 {
		if json.Unmarshal(*data[1], params) != nil {
			returnBadRequest(w)
			return false
		}
	}

	return true
}
