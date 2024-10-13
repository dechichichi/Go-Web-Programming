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
func err(w http.ResponseWriter, r *http.Request) {

}

func login(w http.ResponseWriter, r *http.Request) {

}
func logout(w http.ResponseWriter, r *http.Request) {

}
func signup(w http.ResponseWriter, r *http.Request) {

}
func signup_account(w http.ResponseWriter, r *http.Request) {

}
func authenticate(w http.ResponseWriter, r *http.Request) {

}
func newThread(w http.ResponseWriter, r *http.Request) {

}
func createThread(w http.ResponseWriter, r *http.Request) {

}
func postThread(w http.ResponseWriter, r *http.Request) {

}
func readThread(w http.ResponseWriter, r *http.Request) {

}
