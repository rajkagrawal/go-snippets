package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

func main() {
	has := sha256.New()
	parent := "go/sha256/"
	file, err := os.Open(parent+"raj.txt")
	if err != nil {
		panic(err )
	}
	if _, err := io.Copy(has,file );err != nil {
		panic(err)
	}

	fmt.Println("\n",has.Sum(nil))
	has.Size()
	fmt.Printf("%x",has.Sum(nil))
}
