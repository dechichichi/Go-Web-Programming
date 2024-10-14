package main

import (
	"errors"
	"net/http"
	"text/template"
)

//检查和识别用户是否登录
func session(w http.ResponseWriter,r *http.Request)(sess data.Session,err error){
	cookie,err:=r.Cookie("_cookie")
	if err==nil{
		sess=data.Session{
			Uuid:cookie.Value}
		if ok,_:=sess.Check();ok{
			err=errors.New("Invalid session")
		}
	}
	return
}

//处理器
func index(w http.ResponseWriter,r *http.Request){
	threads,err:=data.Threads();
	if err!=nil{
    _,err:=session(w,r)
	public_tmpl_files:=[]string{
		"templates/layout.html",
		"templates/public.navbar.html",
		"templates/public.index.html"
	}
	private_tmpl_files:=[]string{
		"templates/layout.html",
		"templates/private.navbar.html",
		"templates/index.html"
	}
	var templates *template.Template
	if err!=nil{
		templates=template.Must(template.ParseFiles(private_tmpl_files...))
	}else{
		templates=template.Must(template,r.ParseFiles(public_tmpl_files...))
	}
	templates.ExecuteTemplate(w,"layout",threads)
	}
}