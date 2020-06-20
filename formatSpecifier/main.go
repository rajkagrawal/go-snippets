package main

import (
	"fmt"
)

//%v	the value in a default format
//when printing structs, the plus flag (%+v) adds field names
//%#v	a Go-syntax representation of the value
type S struct {
	Hello string `json:"hello"`
	Some string `json:"some,omitempty"`
}
func main() {
	var s S = S{Hello:"raj",Some:"stringvalue"}
	fmt.Printf("%v\n",s)
	fmt.Printf("%+v\n",s)
	fmt.Printf("%#v",s)

}
