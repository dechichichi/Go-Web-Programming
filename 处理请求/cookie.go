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
		Name: "flash",
		//使用Base64URL编码 以此满足首部对cookie值的URL编码要求
		//因为闪现消息通常会包含空格之类的 需要对cookie值进行编码
		Value: base64.RawURLEncoding.EncodeToString(msg),
	}
	http.SetCookie(w, &c)
}

func showMessage(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("flash")
	//尝试获取指定cookie
	if err != nil {
		//如果没有找到指定cookie 把变量设置成一个http.ErrNoCookie值
		fmt.Fprintln(w, "No message found")
	} else {
		//若找到指定cookie
		//1:将其MaxAge设置为-1 并且将Expires设置为一个过去的时间
		rc := http.Cookie{
			Name:    "flash",
			MaxAge:  -1,
			Expires: time.Unix(1, 0),
		}
		//2:发送同名cookie
		http.SetCookie(w, &rc)
		//上面两个操作是要完全移除这个cookie
		//在设置完新cookie后 程序会对存储在旧cookie中的消息进行解码
		//并且通过响应返回这条消息
		val, _ := base64.URLEncoding.DecodeString(c.Value)
		fmt.Fprintln(w, string(val))
	}
}
