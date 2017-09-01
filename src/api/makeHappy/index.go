package makeHappy

import (
	"fmt"
)

var Name = "makeHappy"

type Params struct {
	Action string `json:"action"`
}

type Results struct {
	Status string `json:"status"`
}

func Do(p *Params) interface{} {
	fmt.Printf("Make action: %s\n", p.Action)

	return &Results{"BAD"}
}
