package main

import "log"

type Shape interface {
	draw()
}
type Square struct {
}

func (a *Square) draw() {
	log.Println("some value in square")
}

type Rectangle struct {
}

func (a *Rectangle) draw() {
	log.Println("some value in rectange")
}
func main() {
	var s Shape
	s = getDrawableObj("square")
	s.draw()
}
func getDrawableObj(val string) Shape {
	switch val {
	case "square":
		return &Square{}
	case "rectangle":
		return &Rectangle{}
	}
	return nil
}
