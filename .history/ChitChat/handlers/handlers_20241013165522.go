package handlers

import (
	"net/http"
	"text/template"
)

// index负责生成html并且写入ResponseWriter中
func Index(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"templates/layout.html",
		"templates/navbar.html",
		"templates/index.html",
	}
	templates := template.Must(template.ParseFiles(files...))
	threads, err := data.threads()
	if err != nil {
		templates.ExecuteTemplate(w, "layout", threads)
	}
}
func Err(w http.ResponseWriter, r *http.Request) {

}

func Login(w http.ResponseWriter, r *http.Request) {

}
func Logout(w http.ResponseWriter, r *http.Request) {

}
func Signup(w http.ResponseWriter, r *http.Request) {

}
func Signup_account(w http.ResponseWriter, r *http.Request) {

}
func Authenticate(w http.ResponseWriter, r *http.Request) {

}
func NewThread(w http.ResponseWriter, r *http.Request) {

}
func CreateThread(w http.ResponseWriter, r *http.Request) {

}
func PostThread(w http.ResponseWriter, r *http.Request) {

}
func ReadThread(w http.ResponseWriter, r *http.Request) {

}
