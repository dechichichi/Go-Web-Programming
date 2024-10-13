package main
import(
	"net/http"
)

func main(){
	mux:=http.NewServeMux()
	files:=http.FileSever(http.Dir("/public"))
	mux.Handle
}