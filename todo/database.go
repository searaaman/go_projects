package main
import (
	"fmt"
	"database/sql"
)

func connectDB(){
	var err error
	constr:="host=localhost port=5432 user=postgres password=12345678 dbname=todoapp sslmode=disable"
	db,err=sql.Open("postgres",constr)

	if err!=nil{
		panic(err)
	}
	err=db.Ping()
	if err!=nil{
		panic(err)
	}
	fmt.Println("connected to db succesfully")
}