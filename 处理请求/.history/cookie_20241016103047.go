package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"
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

func setMessage(w http.ResponseWriter, r *http.Request) {
	msg := []byte("Hello,world")
	c := http.Cookie{
		Name:  "flash",
		Value: base64.RawURLEncoding.EncodeToString(msg),
	}
	http.SetCookie(w, &c)
}

func showMessage(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("flash")
	if err != nil {
		panic(err)
	} else {
		rc := http.Cookie{
			Name:    "flash",
			MaxAge:  -1,
			Expires: time.Unix(1, 0),
		}
		http.SetCookie(w, &rc)
		val, _ := base64.URLEncoding.DecodeString(c.Value)
		fmt.Fprintln(w, string(val))
	}
}
