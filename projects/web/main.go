package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Info struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Info{Name: "John", Age: 25})
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome man")
	})
	r.HandleFunc("/about", aboutHandler)
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
