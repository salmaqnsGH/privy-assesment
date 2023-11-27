package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

type Employee struct {
	ID     *int   `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Salary int    `json:"salary"`
}

func (a *App) Initialize(user, password, dbname string) {
	connString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

	var err error
	a.DB, err = sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
	a.initializaROutes()
}

func (a *App) initializaROutes() {
	a.Router.HandleFunc("/employees", a.getEmployees).Methods("GET")
}

func (a *App) getEmployees(w http.ResponseWriter, r *http.Request) {
	employees, err := getEmployees(a.DB)
	if err != nil {
		return
	}

	response, _ := json.Marshal(employees)
	w.WriteHeader(200)
	w.Write(response)
}

func main() {
	a := App{}

	a.Initialize(
		"root",
		"secret",
		"privy",
	)
	log.Fatal(http.ListenAndServe(":8080", a.Router))
}

func getEmployees(db *sql.DB) ([]Employee, error) {
	query := `SELECT id, name,age, salary FROM employees`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	employees := []Employee{}

	for rows.Next() {
		var employee Employee

		err := rows.Scan(&employee.ID, &employee.Name, &employee.Age, &employee.Salary)
		if err != nil {
			return nil, err
		}

		employees = append(employees, employee)
	}

	return employees, nil
}
