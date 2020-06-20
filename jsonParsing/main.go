package main

import (
	"encoding/json"
	"fmt"
	"strings"
)
type S struct {
	Hello string `json:"hello"`
	Some string `json:"some,omitempty"`
}
func main() {
	var s S
	sss := strings.NewReader(`{"hello":"raj"}`)
	err := json.NewDecoder(sss).Decode(&s)
	fmt.Println(err)
	fmt.Println(s)
	x , err := json.Marshal(s)
	fmt.Println(string(x))

}
