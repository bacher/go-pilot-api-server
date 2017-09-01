package getInitialData

import "fmt"

var Name = "getInitialData"

type Params struct {
	Name string `json:"name"`
}

var data map[string]string

func Do(p *Params) interface{} {
	if data == nil {
		data = make(map[string]string)

		data["field1"] = "Who are you?"
		data["field2"] = "What is your name?"
	}

	fmt.Println(p.Name)

	return data
}
