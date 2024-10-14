package main

import (
	"fmt"
	"net/http"
)

// 一个处理器就是一个拥有ServeHTTP方法的接口
type MyHandler struct{}

// ServeHTTP方法 接受两个参数 第一个是一个ResponseWriter接口 第二个是一个指向http.Request的指针
func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello,world!")
}

func main() {
	//这个处理器代替默认的多路处理器
	handler := MyHandler{}
	//创建一个服务器 (网络地址,负责处理请求的处理器)
	//若网络地址参数为空 默认80端口
	//若处理器参数为nil 则默认使用DefaultServeMux
	http.ListenAndServe("", nil)
	//下面这个代码和上面的唯一区别是下面这个可以通过Server结构对服务器进行更多的配置
	server := http.Server{
		Addr:    "127.0.0.1:8080", //与上行的第一个参数一样
		Handler: &handler,         //与上行的第二个参数一样
		//下面参数可选(与上面的唯一区别)
	}
	//将HTTP通信放到SSL之上进行 HTTPS服务
	server.ListenAndServeTLS("cert.pem", "key.pem") //cert.pem SSL证书 key.pem 服务器私钥
}
