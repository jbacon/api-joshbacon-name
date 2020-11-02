package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// Employees struct
type employees []employee

// Employee struct
type employee struct {
	Gender string `json:"gender"`
	ID     int    `json:"id"`
}

func employeesHandler(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, employees, err := getEmployees(db)
		if err != nil {
			log.Printf("Something went wrong querying database: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Something went wrong"))
			return
		}
		json, err := json.MarshalIndent(employees, "", "	")
		if err != nil {
			log.Printf("Something went wrong building json: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Something went wrong"))
			return
		}
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	})
}

func getEmployees(db *sql.DB) ([]string, employees, error) {
	stmt, err := db.Prepare("SELECT gender, id FROM employees")
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to prepare query string to sql statement: %v", err)
	}
	queryResult, err := stmt.Query()
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to execute sql statement: %v", err)
	}
	defer queryResult.Close()
	columnNames, err := queryResult.Columns()
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to get columns: %v", err)
	}
	employees := employees{}
	for queryResult.Next() {
		employee := employee{}
		err := queryResult.Scan(&employee.Gender, &employee.ID)
		if err != nil {
			return nil, nil, fmt.Errorf("Failed to get scan database row into employee struct: %v", err)
		}
		employees = append(employees, employee)
	}
	return columnNames, employees, nil
}

func main() {
	db, err := sql.Open("sqlite3", "./employees.db")
	if err != nil {
		log.Fatalf("Could not open db: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalf("Missing required environment variable PORT")
	}

	serveMux := http.NewServeMux()

	serveMux.Handle("/employees", loggingMiddleware(employeesHandler(db)))

	httpServer := &http.Server{
		Addr:    ":" + port,
		Handler: serveMux,
	}
	log.Println("Listening...")
	err = httpServer.ListenAndServe()
	if err != nil {
		log.Fatalf("Server failed to listen and server: %v", err)
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request received")
		next.ServeHTTP(w, r)
		log.Println("Request handled")
	})
}
