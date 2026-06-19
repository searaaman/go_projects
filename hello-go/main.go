package main
import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter,r *http.Request){
	if r.Method=="GET" && r.URL.Path =="/"{

	
	fmt.Fprintln(w,r.URL.Path)
	}
}

func healthHandler(w http.ResponseWriter , r *http.Request){
	if 	r.Method=="GET" && r.URL.Path=="/health"{
	fmt.Fprintln(w,"status up")
	}
}
func greetHandler(w http.ResponseWriter,r*http.Request){
	name:=r.URL.Query().Get("name")
	if r.Method=="GET"&& r.URL.Path=="/greet"{
		fmt.Fprintln(w,"HELLO "+ name +"!!")
	}
}

func main(){
	http.HandleFunc("/",homeHandler)
	http.HandleFunc("/health",healthHandler)
	fmt.Println("Server started and running on port 8080")
	http.ListenAndServe(":8080",nil)
}