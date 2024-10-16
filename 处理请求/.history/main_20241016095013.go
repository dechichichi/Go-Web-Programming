package main

import (
	"fmt"
	"net/http"
)

// 请求首部
func headers(w http.ResponseWriter, r *http.Request) {
	//这个服务器会把请求的首部打印出来
	//h := r.Header
	//如果想要获取某个特定的首部
	//h := r.Header["Sec-Fetch-Dest"]
	//或者
	h := r.Header.Get("Sec-Fetch-Dest")
	//直接引用Header得到一个字符串切片
	//在Header上调用Get返回字符串
	fmt.Fprintln(w, h)
}

// 请求主体
func body(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	fmt.Fprintln(w, string(body))
}

func main() {
	//Request结构
	//表示一个由客户端发送的HTTP请求
	//主要由 URL,Header,Body,Form,PostForm,MultipartForm等字段组成
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/body", body)
	server.ListenAndServe()
	Form()
	Func_Write()
}
