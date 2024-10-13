package main
import(
	"net/http"
)

func main(){
	mux:=http.NewServeMux()
	files:=http.FileSever(http.Dir("/public"))
	mux.Handle("/static/",http.StripPrefix("/static/",files))
	mux.HandleFunc("/",index)
	server:=&http.Server{
		Addr:"0.0.0.0:8080",
		Handle:mux,
	}
	server.ListenAndServe()
}