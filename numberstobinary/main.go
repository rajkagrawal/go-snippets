package main

import (
	"fmt"
	"math"
)

func main() {
	for i:=1;i<int(math.Pow((float64(4)),float64(2)));i++{
		fmt.Print(fmt.Sprintf("%04b\n",i))
	}
}
