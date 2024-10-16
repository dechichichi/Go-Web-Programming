package main

import (
	"net/http"
)

func WriteExample(w http.ResponseWriter, r *http.Request) {
	str := "<html><head><title>Go Web</title></head><body><h1>Hello World</h1></body></html>"
	w.Write([]byte(str))
}
