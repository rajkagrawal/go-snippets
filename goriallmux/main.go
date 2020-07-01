package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/events/{sub}", func(writer http.ResponseWriter, request *http.Request) {
		variable :=  mux.Vars(request)
		subValue := variable["sub"]
		page := variable["page"]
		fmt.Println("sub value : ", subValue)
		fmt.Println("page value : ", page)
		pageValue := request.URL.Query().Get("page")
		fmt.Println("page value : ",pageValue)
		writer.WriteHeader(http.StatusNoContent)

	})
	fmt.Println(router)
	http.ListenAndServe(":8080",router)

}
