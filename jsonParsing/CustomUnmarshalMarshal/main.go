package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Person struct {
	Name string `json:"name"`
	BirthTime time.Time `json:"birth"`
}

func (a *Person) UnmarshalJSON(data []byte) error {
	type AnotherPerson Person
	sss :=  struct {
		*AnotherPerson
		T int64 `json:"birth"`
	}{
		AnotherPerson: (*AnotherPerson)(a),
	}
	err := json.Unmarshal(data,&sss)

	if err != nil {
		fmt.Println("error in mfirst unmarshalling",err)
		return nil
	}

	a.BirthTime = time.Unix(sss.T,0)	//*a = Person(*sss.P)
	//a.BirthTime = time.Unix(sss.T,0)
	return nil
}

func (a *Person) MarshalJSON() ([]byte,error) {
	type AnotherPerson Person
	sss :=  struct {
		*AnotherPerson
		T int64 `json:"birth"`
	}{
		AnotherPerson: (*AnotherPerson)(a),
		T : a.BirthTime.Unix(),
	}
	return json.Marshal(sss)

}

func main() {
fmt.Println( time.Now().Unix())
str :=  `{"name":"raj","birth":1593627762}`
var p *Person
json.Unmarshal([]byte(str),&p)
fmt.Println(p)
val , _ := json.Marshal(p)
fmt.Println(string(val))

}
