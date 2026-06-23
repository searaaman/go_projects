package main
import (
	"encoding/json"
	"net/http"
	"fmt"
	"strings"
	"strconv"
)
var Todos []Todo
var currentID int

type Todo struct {
	ID 		  int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func todosHandler(w http.ResponseWriter, r *http.Request){
	if r.Method=="GET"{
		getTodoHandler(w,r)
		return
	}
	if r.Method=="POST"{
		createTodoHandler(w,r)
		return
	}
	if r.Method=="DELETE"{
		deleteTodoHandler(w,r)
		return
	}
	if r.Method=="PUT"{
		updateTodoHandler(w,r)
		return
	}
	http.Error(w,"Method not allowed",http.StatusMethodNotAllowed)

}

func getTodoHandler(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(Todos)
}

func createTodoHandler(w http.ResponseWriter ,r *http.Request){

	if r.Method!= "POST"{
	http.Error(w,"only post is allowed",http.StatusMethodNotAllowed)
	return
	}
	var todo Todo
	err:=json.NewDecoder(r.Body).Decode(&todo)
	if err !=nil{
		http.Error(w,"invalid json",http.StatusBadRequest)
		return
	}
	currentID++
	todo.ID=currentID
	Todos=append(Todos,todo)
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(todo)
}

func deleteTodoHandler(w http.ResponseWriter, r *http.Request){
	path:=r.URL.Path
	parts:=strings.Split(path,"/")
	ID,err:=strconv.Atoi(parts[2])
	if err!=nil{
		http.Error(w, "invalid id ", http.StatusBadRequest)
		return
	}
	for index,todo:=range Todos{
		if todo.ID==ID{
			Todos =append(Todos[:index],Todos[index+1:]...)
			w.Write([]byte("Todo deleted succesfully"))
			return

		}
	}
	http.Error(w,"Todo not found",http.StatusNotFound)
	return
	
	

}

func updateTodoHandler(w http.ResponseWriter,r *http.Request){
	path:=r.URL.Path
	parts:=strings.Split(path,"/")
	ID,err:=strconv.Atoi(parts[2])
	if err!=nil{
		http.Error(w,"Invalid ID",http.StatusNotFound)
		return
	}
	var updatedtodo Todo
	err=json.NewDecoder(r.Body).Decode(&updatedtodo)
	if err !=nil{
		http.Error(w,"INVLID JSON",http.StatusBadRequest)
		return
	}

	for index,todo:=range Todos{
		if todo.ID==ID{
			updatedtodo.ID=ID
			Todos[index]=updatedtodo
		w.Header().Set("Content-Type","application/json")
		json.NewEncoder(w).Encode(updatedtodo)

		return
		}
	}
	http.Error(w,"Todo not found",http.StatusNotFound)

}

func main(){

	http.HandleFunc("/todos/",todosHandler)
	

	Todos =[]Todo{
		{ID: 1, Title: "Learn Go", Completed: false},
		{ID: 2, Title: "Build Todo API", Completed: true},
	}
	currentID=2
	
	fmt.Println("Server started and running on port 8080")
	fmt.Println(http.ListenAndServe(":8080",nil))
}