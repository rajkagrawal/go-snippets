package main

import (
	"fmt"
	"sync"

)
var s sync.Mutex
type obj struct {
  s sync.Mutex
}

func (a *obj) Write() {
	a.s.Lock()
}

var a *obj
var b = &sync.Once{}
func NewObj() *obj {
	if a == nil {
		b.Do(func() {
			fmt.Println("hello raj")
			a = new(obj)
		})
	}
	return a
}

func main() {
	fmt.Println(NewObj())
	fmt.Println(NewObj())
	fmt.Println(NewObj())
	fmt.Println(NewObj())

}