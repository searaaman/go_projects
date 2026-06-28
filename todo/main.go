package main
import (
	"net/http"
	"fmt"
	"database/sql"
	_"github.com/lib/pq"
)

var db *sql.DB
func main(){
	connectDB()

	http.HandleFunc("/todos/",todosHandler)
	
	
	fmt.Println("Server started and running on port 8080")
	fmt.Println(http.ListenAndServe(":8080",nil))
}