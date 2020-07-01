package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"net/http/pprof"
	"runtime"
	"strconv"
	"sync"
	"time"
)
//curl -o heap.out -XGET  http://localhost:8080/profile/heap
//curl -o cpu.out -XGET  http://localhost:8080/profile/cpu?seconds=12
//curl -o trace.out -XGET  http://localhost:8080/profile/trace?seconds=12
//go tool trace trace.out
//pprof cpu.out  we can also use go tool pprof but prefer to use google pprof
//pprof heap.out
//top -cum , list <function name> etc are some commands


func main() {
	router := mux.NewRouter()
	profilingRouter := router.PathPrefix("/profile").Subrouter()
	profilingRouter.Handle("/heap",pprof.Handler("heap"))
	profilingRouter.HandleFunc("/cpu",pprof.Profile)
	profilingRouter.HandleFunc("/trace",pprof.Trace)
	router.HandleFunc("/raj", func(writer http.ResponseWriter, request *http.Request) {

		fibQueryVal := request.URL.Query().Get("fib")
		fib,_  := strconv.Atoi(fibQueryVal)
		fmt.Fprintf(writer,"The fib value : %d",getFib(fib))

	})
	//router.Handler("/",pprof.Handler("heap"))
	http.ListenAndServe(":8080",router)
}
var arrdsay = make([]int,0)
func getFib(n int)int {
	runtime.GOMAXPROCS(runtime.NumCPU())
	a := 0
	b := 1
	wg := &sync.WaitGroup{}
	wg.Add(n)
//	time.Sleep(200*time.Millisecond)
	for i := 1;i<=n ; i++ {
		fmt.Println("this is in for loop")
		go func() {
			a = a + b
			a, b = b, a
			for j := 10; j < 1000; j++ {
				ti := rand.Intn(893282389) * rand.Intn(32093292309) * j
				arrdsay = append(arrdsay, ti)
			}
			time.Sleep(1000*time.Millisecond)
			fmt.Println("has the done being called")
			wg.Done()
		}()

	}
	fmt.Println("wainting here ",time.Now())
	wg.Wait()
	fmt.Println("aftere waint ",time.Now())
	return 1
}