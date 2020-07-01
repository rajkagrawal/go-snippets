package main

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
)
type S struct {
	Hello string `json:"hello"`
	Some string `json:"some,omitempty"`
}

type movie struct {
	Genre []string `json:"genre"`
	Title  string `json:"title"`
	Year int

}
func main() {
	var s S
	sss := strings.NewReader(`{"hello":"raj"}`)
	err := json.NewDecoder(sss).Decode(&s)
	fmt.Println(err)
	fmt.Println(s)
	x , err := json.Marshal(s)
	fmt.Println(string(x))
	str := `{ "genre":["action","romantic"],"title":"interstellar","year":2014}`
	var m movie
	err = json.Unmarshal([]byte(str),&m)
	fmt.Println(err , m)
}
