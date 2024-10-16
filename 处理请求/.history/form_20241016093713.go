package main

import (
	"fmt"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	//使用ParseForm对请求进行语法分析
	//再访问Form字段获取具体表单
	r.ParseForm()
	fmt.Fprintln(w, r.Form)
}

func Form() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
