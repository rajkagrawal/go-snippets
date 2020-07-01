package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

const (
	ClaimStatusFulfilled ClaimStatus = iota + 1
	ClaimStatusRejected
)

type ClaimStatus int

func (s ClaimStatus) String() string {
	switch s {
	case 1:
		return "FULFILLED"
	case 2:
		return "REJECTED"
	}
	return "ERROR"
}
func (s ClaimStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, s)), nil
}

func (s *ClaimStatus) UnmarshalJSON(buf []byte) error {
	fmt.Println("raj",string(buf))
	val,_ := strconv.Unquote(string(buf))
	switch val  {
	case "FULFILLED":
		*s = ClaimStatus(ClaimStatusFulfilled)
		return nil
	case "REJECTED":
		dummy := ClaimStatus(ClaimStatusRejected)
		fmt.Println("hello raj rejct",ClaimStatusFulfilled)
		*s = dummy
		return nil
	}
	fmt.Println("hello after default raj")

	*s = 0
	return nil
}

type Claim struct {
	// a bunch of other stuff
	Status ClaimStatus `json:"status"`
	// a bunch of other stuff
}

func main() {
	obj := []byte(`{"status": "FULFILLED"}`)

	fmt.Printf("byte object: %s\n", obj)

	var claim Claim
	err := json.Unmarshal(obj, &claim)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return
	}

	fmt.Printf("got claim: %+v", claim)
}

