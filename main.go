package main

import (
	"apirest-gorm/database"
	"apirest-gorm/handlers"
	"apirest-gorm/repository"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	log.SetFlags(0)
	database.Connect()
	repository.Init(database.DB)

	mux := mux.NewRouter()	
	mux.HandleFunc("/api/user/", handlers.GetAllUsers).Methods("GET")	
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.GetOneUser).Methods("GET")	
	mux.HandleFunc("/api/user/", handlers.CreateUser).Methods("POST")
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT")
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.RemoveUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", mux))
}
