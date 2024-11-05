package main

import (
	"log"
	"net/http"
	"time"
)

func greetHandler(w http.ResponseWriter, req *http.Request) {
	startTime := time.Now()
	log.Println("server: greet handler started")

	select {
	case <-time.After(10 * time.Second):
		w.Write([]byte("greeting\n"))
	case <-req.Context().Done():
		log.Println("server: context cancelled")
		http.Error(w, req.Context().Err().Error(), http.StatusInternalServerError)
	}
	log.Println("server: greet handler ended in", time.Since(startTime))
}

func sayHi(w http.ResponseWriter, req *http.Request) {
	log.Println("We are sayHi function...")
	starttime := time.Now()

	select {
	case <-time.After(10 * time.Second):
		w.Write([]byte(`{"message": "Hello, World!"}`))
	case <-req.Context().Done():
		log.Println("Context error is : ", req.Context().Err())
		http.Error(w, req.Context().Err().Error(), http.StatusInternalServerError)

	}
	log.Println("server: greet handler ended in", time.Since(starttime))
}

func main() {
	http.HandleFunc("/greet", greetHandler)
	http.HandleFunc("/hi", sayHi)
	log.Fatal(http.ListenAndServe(":8090", nil))
}
