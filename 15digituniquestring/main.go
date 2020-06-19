package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	stringVal := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYS123456789"
	rand.Seed(time.Now().UnixNano())
	randShuffle(stringVal)
	manualShuffle(stringVal)
}
func manualShuffle(stringVal string)  {
	runeVal := []rune(stringVal)
	for i:=len(runeVal)-1;i>=0;i--{
		j := rand.Intn(i+1)
		runeVal[i], runeVal[j] = runeVal[j], runeVal[i]
	}

	fmt.Println(string(runeVal[:15]))
}
func randShuffle(stringVal string) {
	runeVal := []rune(stringVal)
	rand.Shuffle(len(runeVal), func(i, j int) {
		runeVal[i], runeVal[j] = runeVal[j], runeVal[i]
	})
	fmt.Println(string(runeVal[:15]))
}
