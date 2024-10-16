package main

import "net/http"

func WriteExample(w http.ResponseWriter, r *http.Request) {
	str := "<html><head>head><title>Go Web</title></head><body><h1>Hello World</h1><body></html>"
	w.Write([]byte(str))
}

func Func_Write() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/write", WriteExample)
	server.ListenAndServe()
}
