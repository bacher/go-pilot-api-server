package getInitialData

import (
	"bytes"
)

var Name = "getInitialData"

type Params struct {
	Name string `json:"name"`
}

type Response struct {
	Name  string `json:"name"`
	Age   int8   `json:"age"`
	Rules []Rule `json:"rules"`
}

type Rule struct {
	Id   int32  `json:"id"`
	Text string `json:"text"`
}

//var data map[string]string

func Do(p *Params) interface{} {
	var buf bytes.Buffer

	for i := 0; i < 100; i++ {
		buf.WriteString("kek")
	}

	//mmm := &map[string]string{
	//	"field1": fmt.Sprintf("%s, Who Am I?", p.Name),
	//	"field2": buf.String(),
	//}

	resp := &Response{
		Name: p.Name,
		Age:  31,
		Rules: []Rule{
			{1, "Hello"},
			{2, p.Name},
			{3, "Goodbye"},
		},
	}

	return resp
}
