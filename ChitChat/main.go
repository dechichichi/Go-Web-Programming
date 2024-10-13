package main

import (
	"chitchat/handlers"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	//当服务器收到一个以/static/开头的URL请求时 以上两行代码会移除URL中的/static前缀
	//然后在public目录中查找被请求的文件
	//eg: http://localhost/static/css/bootstrap.min.css
	//将会在public目录查找 /css/bootstrap.min.css 文件并返回客户端
	mux.HandleFunc("/", handlers.Index)
	mux.HandleFunc("/err", handlers.Err)
	mux.HandleFunc("/login", handlers.Login)
	mux.HandleFunc("/logout", handlers.Logout)
	mux.HandleFunc("/signup", handlers.Signup)
	mux.HandleFunc("/signup_account", handlers.Signup)
	mux.HandleFunc("/authenticate", handlers.Authenticate)
	mux.HandleFunc("/thread/new", handlers.NewThread)
	mux.HandleFunc("/thread/create", handlers.CreateThread)
	mux.HandleFunc("/thread/post", handlers.PostThread)
	mux.HandleFunc("/thread/read", handlers.ReadThread)
	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
