package main
import (
	"net/http"
	_"github.com/lib/pq"
	"strconv"
	"strings"
	"fmt"
	"encoding/json"

)

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
	rows,err:=db.Query("SELECT * FROM todos")
	if err!=nil{
		http.Error(w,"Failed To Fetch Tables ",http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var todos []Todo
	for rows.Next(){
		var todo Todo

		err:=rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.Completed,
		)
		if err!=nil{
			http.Error(w,"invalid",http.StatusInternalServerError)
		}
		todos=append(todos,todo)
	}
	json.NewEncoder(w).Encode(todos)
}

func createTodoHandler(w http.ResponseWriter ,r *http.Request){

	var todo Todo
	err:=json.NewDecoder(r.Body).Decode(&todo)
	if err !=nil{
		http.Error(w,"invalid format",http.StatusBadRequest)
		return
	}

	query:="INSERT INTO todos(title,completed) VALUES($1,$2)"
	_,err=db.Exec(query,todo.Title,todo.Completed)

	if err!=nil{
		fmt.Println(err)
		http.Error(w,"failed to create todo",http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Todo Created successfully"))
}

 

func updateTodoHandler(w http.ResponseWriter,r *http.Request){
	path:=r.URL.Path
	parts:=strings.Split(path,"/")
	ID,err:=strconv.Atoi(parts[2])
	var updatedtodo Todo
	err=json.NewDecoder(r.Body).Decode(&updatedtodo)
	if err !=nil{
		http.Error(w,"INVLID JSON",http.StatusBadRequest)
		return
	}
	query:="UPDATE TODOS SET TITLE =$1 ,COMPLETED =$2  WHERE ID =$3"
	result,err:=db.Exec(query,
		updatedtodo.Title,
		updatedtodo.Completed,
		ID,)
	
	
	if err!=nil{
		http.Error(w,"Failed to update todo",http.StatusInternalServerError)
		return
	}
	rowsAffected,err:=result.RowsAffected()
	if err!=nil{
		http.Error(w,"FAILED TO Determine Rows Affected",http.StatusBadRequest)
		return
	}
	if rowsAffected==0{
		http.Error(w,"Todo not found",http.StatusNotFound)
		return
	}

	
	w.Write([]byte("Todo updated successfully"))

}
func deleteTodoHandler(w http.ResponseWriter, r *http.Request){
	path:=r.URL.Path
	parts:=strings.Split(path,"/")
	ID,err:=strconv.Atoi(parts[2])
	query:="DELETE FROM TODOS WHERE ID=$1"
	result,err:=db.Exec(query,ID)
	if err!=nil{
		http.Error(w, "Failed to delete todo ", http.StatusBadRequest)
		return
	}
	rowsAffected,err:=result.RowsAffected()
	if err!=nil{
		http.Error(w,"Failed to determine affected rows",http.StatusBadRequest)
		return
	}
	if rowsAffected==0{
		http.Error(w,"Todo not found",http.StatusNotFound)
		return
	}
	w.Write([]byte("Todo deleted  successfully"))
	return
	
	

}
