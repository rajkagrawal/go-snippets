package main

import (
	"errors"
	"fmt"
	"strconv"
)

// Earlier when we embed other error while returning
// fmt.Errorf("%v %v","some context", err) the err details were removed and only string part of error is only retained

// But after 1.13 we can retain the error value using %w for existing error or wrapping that error in Unwrap method
// fmt.Errorf("decompress %v: %w", name, err)

type SomeErr struct {
	ss  string
	Err error
}
type SecondErr struct {
	ss  string
	ii  int
	Err error
}

func (a *SomeErr) Error() string {
	return a.ss
}
func (a *SomeErr) Unwrap() error {
	return a.Err
}

var ErrRaj = errors.New("hello raj this is sentinel err")
var Xyz = errors.New("this is xys")

func (a *SecondErr) Error() string {
	return a.ss + strconv.Itoa(a.ii)
}
func (a *SecondErr) Unwrap() error {
	return a.Err
}

func main() {
	fa := getError()
	//Only if unwrap method is present on both SomeErr and SecondErr
	//errros.Is works mostly on sentinel value
	if errors.Is(fa, ErrRaj) {
		fmt.Println("this is in as")
	}
	// Unwrap method needs to be present on SomeErr which wraps this SecondErr
	var x *SecondErr
	if errors.As(fa, &x) {
		fmt.Println(x.Error())
	}
	var a *SomeErr
	if errors.As(fa, &a) {
		fmt.Println(a.Error())
	}
	fmt.Println("checking fmt errorf with %w")
	errrr := getFmtError()

	// WE can get the details of addition information error as it was embedded with %w
	// If not embedded with %w we would only get textual inforamtion discarding other information
	var xy *SecondErr
	if errors.As(errrr, &xy) {
		fmt.Println(x.Error())
	}

}

func getFmtError() error {
	err := &SecondErr{ss: "second tyupe", ii: 12, Err: ErrRaj}
	//Here if we add %w then we can query this error object later
	return fmt.Errorf("%s : %w", "thisisfmterr", err)
}

func getError() error {
	err := fu()
	return &SomeErr{"sds", &SecondErr{ss: "raj", ii: 223, Err: err}}

}
func fu() error {
	return ErrRaj
}
