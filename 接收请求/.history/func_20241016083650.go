package main

import (
	"fmt"
	"net/http"
)

// 处理器函数:与处理器有相同行为的函数 与ServeHTTP方法有相同的签名
// 即:他们接受ResponseWriter 和指向Request结构的指针作为参数
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

func Serve() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)
	server.ListenAndServe()
}
