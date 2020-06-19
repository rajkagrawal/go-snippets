package main

import "fmt"

type Observer interface {
	update(score int)
}
type Subject struct {
	 cricketScore int
	 observers []Observer
}

//Methods like deregister can also be put
func NewSubject() *Subject{
	return &Subject{}
}

func (s Subject) stream(i int) {
	for _, val := range s.observers{
		val.update(i)
	}
}
type ObserverTelevision struct {
	s *Subject
}
func (a *ObserverTelevision) update(i int ) {
	fmt.Println("score :",i)
}
type ObserverRadio struct {
	s *Subject
}
func (a *ObserverRadio) update(i int ) {
	fmt.Println("audio score :",i)
}
func (s *Subject) AddObserver(in Observer) {
	s.observers  = append(s.observers, in )
}
func NewObserverRadio(s *Subject) *ObserverRadio {
	radioObserver := &ObserverRadio{s : s }
	s.AddObserver(radioObserver)
	return radioObserver
}

func NewObserverTelevision(s *Subject) *ObserverTelevision {
	televisionObserver := &ObserverTelevision{s : s }
	s.AddObserver(televisionObserver)
	return televisionObserver
}

func main() {
	sub := NewSubject()
	 NewObserverTelevision(sub)
	 NewObserverRadio(sub)
	sub.stream(5)
}



