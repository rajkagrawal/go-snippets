package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"

	"log"
	"net/http"
	"time"
)

func myfun()  {
	if r := recover();r!=nil {
		fmt.Println(r)
	}
}
func main() {
	go func() {
	muxer := mux.NewRouter()
	muxer.HandleFunc("/get", func(writer http.ResponseWriter, request *http.Request) {
		typeStruct := struct {
			Raj string `json:"raj"`
		}{}
		time.Sleep(10*time.Second)
		fmt.Println("hello raj")
			d := json.NewDecoder(request.Body).Decode(&typeStruct)
			if d != nil {
				fmt.Println("error deconding the reques")
			}
			fmt.Println(typeStruct)
			fmt.Fprintf(writer,"hello %s","raj")
	}).Methods(http.MethodPost)
		http.ListenAndServe(":8081",muxer)
	}()
	ctx := context.Background()
	ctx = context.WithValue(ctx,"raj","kumar")
	ctxWithTimeout, cancelFunc := context.WithTimeout(ctx,2*time.Second)
	defer cancelFunc()
	send(ctxWithTimeout)
	ctxWithTimeoutSec, cancelFuncSec := context.WithTimeout(ctx,3*time.Second)
	defer cancelFuncSec()
	sendSecond(ctxWithTimeoutSec)
	fmt.Println("hello raj")
	client := http.Client{}
	typeStruct := struct {
		Raj string `json:"raj"`
	}{
		Raj : "hello raj ",
	}
	by , err := json.Marshal(typeStruct)
	defer myfun()
	if err != nil {
		panic("error marshalling the value ")
	}
	ct := context.Background()
	request, err := http.NewRequest(http.MethodPost,"http://localhost:8081/get", bytes.NewReader(by)  )
	if err != nil {
		panic("err in requeist formation")
	}
	ct , cf := context.WithTimeout(ct,5*time.Second)
	defer cf()
	request = request.WithContext(ct)
	res, err :=  client.Do(request)
	if err != nil {

		panic("error calling the client"+err.Error())
	}
	time.Sleep(20*time.Second)
	defer res.Body.Close()
	if res.StatusCode  == 200 {
		var x interface{}
		json.NewDecoder(res.Body).Decode(x)
	}else{
		log.Println("erro in getting response bod ", res.Body, res.StatusCode)
	}
}

func send(ctx context.Context)  {
		fmt.Println(ctx.Value("raj"))
		select {
		case <-time.NewTimer(5*time.Second).C:
			fmt.Println("hello raj")
		case <-ctx.Done():
			fmt.Println("context done")
		}
}

func sendSecond(ctx context.Context){
	for {
		x := time.Now()
		time.Sleep(9*time.Second)
		fmt.Println(time.Since(x))
		break
	}
}