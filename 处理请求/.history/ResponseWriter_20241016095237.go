package main

import (
	"log"
	"net/http"
)

func WriteExample(w http.ResponseWriter, r *http.Request) {
	str := "<html><head><title>Go Web</title></head><body><h1>Hello World</h1></body></html>"
	w.Write([]byte(str))
}

func Func_Write() {
	server := http.Server{
		Addr: "127.0.0.1:8082",
	}
	http.HandleFunc("/write", WriteExample)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
