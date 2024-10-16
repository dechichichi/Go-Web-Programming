package main

import (
	"fmt"
	"net/http"
)

func headers(w http.ResponseWriter, r *http.Request) {
	//这个服务器会把请求的首部打印出来
	h := r.Header
	fmt.Fprintln(w, h)
}

func main() {
	//Request结构
	//表示一个由客户端发送的HTTP请求
	//主要由 URL,Header,Body,Form,PostForm,MultipartForm等字段组成
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/headers", headers)
	server.ListenAndServe()
}
