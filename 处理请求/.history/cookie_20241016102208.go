package main

import (
	"fmt"
	"net/http"
)

//cookie是一种存储在客户端的 体积较小的信息
//每当客户端向服务器发送HTTP请求时 cookie都会随着请求被一同发送至服务器

func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "first_cookie",
		Value:    "Go Web Programming",
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name:     "second_cookie",
		Value:    "Go Web Programming",
		HttpOnly: true,
	}
	//先用Set方法添加第一个cookie
	w.Header().Set("Set-Cookie", c1.String())
	//再用Add方法添加第二个cookie
	w.Header().Add("Set-Cookie", c2.String())
	//一种更为快捷方便的cookie设置方法
	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	h := r.Header["Cookie"]
	fmt.Fprintln(w, h)
}
