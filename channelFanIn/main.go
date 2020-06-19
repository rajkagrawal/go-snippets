package main

import (
	"fmt"
	"sync"
)

func main() {

	cons1 := consume()
	cons2 := consume()
	//FanIn :
	//A function can read from multiple inputs and proceed until all are closed by multiplexing
	//the input channels onto a single channel that's closed when all the inputs are closed

	//Below done is kept just in case consumer only reads few of the values from outpout channel
	//if a stage fails to consume all the inbound values, the goroutines attempting to send those values will block indefinitely:
	//https://blog.golang.org/pipelines
	done := make(chan struct{})
	defer close(done)
	for cha := range merge(done, cons1, cons2) {
		fmt.Println(cha)
	}
}

func merge(done chan struct{}, input ...<-chan int) <-chan int {
	output := make(chan int)
	wg := sync.WaitGroup{}
	fu := func(inp <-chan int) {
		//for inpuVal := range inp {
		//	output <- inpuVal
		//}

		//for {
		//	time.Sleep(4*time.Second)
		//	select {
		//	case inpuVal := <- inp:
		//		output <- inpuVal
		//	case <-done :
		//		fmt.Println("this is retruning")
		//		return
		//	}
		//}

		defer wg.Done()
		for n := range inp {
			select {
			case output <- n:
			case <-done:
				return
			}

		}
	}
	for _, inputChan := range input {
		wg.Add(1)
		go fu(inputChan)
	}
	go func() {
		wg.Wait()
		close(output)
	}()

	return output

}

func consume() <-chan int {
	ch := make(chan int, 0)
	go func() {
		for i := 0; i < 19; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}
