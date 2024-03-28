package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Guarda la conexion
var DB *sql.DB


func Connecta() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	dns := fmt.Sprintf("%v:%v@(%v:%v)/%v",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	//Abrir conexión a la database
	connection, err := sql.Open("mysql", dns)
	if err != nil {
		log.Fatal(err.Error())
	}
	DB = connection

	//Verificar conexión a la db
	Ping()

	log.Printf("%+v\n", "conexion exitosa")
	
}

// Verificar la conexion
func Ping() {
	if err := DB.Ping(); err != nil {
		log.Fatal(err.Error())
	}
}

// Cerrar la Conexion
func Close() {
	DB.Close()
}

func MyCreateTable(schema string) {
	DB.Exec(schema)
}

// Ferificar si una tabla existe o no
func ExistsTable(tableName string) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
	rows, err := DB.Query(sql)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	return rows.Next()
}

// Crear una tabla en la base de datos
func CreateTable(schema, name string) {

	if !ExistsTable(name) {
		_, err := DB.Exec(schema)
		if err != nil {
			fmt.Println(err)
		}
	}
}

// Eliminara Tabla
func TruncateTable(tableName string) {
	sql := fmt.Sprintf("TRUNCATE TABLE %s", tableName)
	log.Printf("%+v\n",sql )
	DB.Exec(sql)
}

// Polimorfismo a Exec
func Exec(query string, args ...interface{}) (sql.Result, error) {
	Connecta()
	result, err := DB.Exec(query, args...)
	Close()
	if err != nil {
		fmt.Println(err)
	}
	
	return result, err
}

// Polimorfismo a Query
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	Connecta()
	rows, err := DB.Query(query, args...)
	Close()
	if err != nil {
		fmt.Println(err)
	}
	return rows, err
}
