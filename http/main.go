package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
)

type Employee struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Profession     string `json:"profession"`
	Salary         int    `json:"salary"`
	ExperienceYear int    `json:"experience_year"`
}


func getEmployeesHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		rows, err := db.Query("SELECT id, name, profession, salary, experience_year FROM employee")
		if err != nil {
			http.Error(w, "Ошибка БД", 500)
			return
		}
		defer rows.Close()

		var employees []Employee

		for rows.Next() {
			var e Employee
			err := rows.Scan(&e.ID, &e.Name, &e.Profession, &e.Salary, &e.ExperienceYear)
			if err != nil {
				http.Error(w, "Ошибка чтения", 500)
				return
			}
			employees = append(employees, e)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(employees)
	}
}


func getEmployeeByID(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		idStr := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Неверный id", 400)
			return
		}

		var e Employee

		err = db.QueryRow(
			"SELECT id, name, profession, salary, experience_year FROM employee WHERE id=$1",
			id,
		).Scan(&e.ID, &e.Name, &e.Profession, &e.Salary, &e.ExperienceYear)

		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Не найден", 404)
			} else {
				http.Error(w, "Ошибка БД", 500)
			}
			return
		}

		json.NewEncoder(w).Encode(e)
	}
}


func createEmployeeHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			http.Error(w, "Метод не поддерживается", 405)
			return
		}

		var e Employee

		err := json.NewDecoder(r.Body).Decode(&e)
		if err != nil {
			http.Error(w, "Неверный JSON", 400)
			return
		}

		if e.Name == "" || e.Profession == "" {
			http.Error(w, "Пустые поля", 400)
			return
		}

		_, err = db.Exec(
			"INSERT INTO employee(name, profession, salary, experience_year) VALUES ($1,$2,$3,$4)",
			e.Name, e.Profession, e.Salary, e.ExperienceYear,
		)

		if err != nil {
			http.Error(w, "Ошибка вставки", 500)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{
			"message": "Сотрудник добавлен",
		})
	}
}

func main() {

	connStr := "host=localhost port=5432 user=postgres password=123 dbname=company sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()


	http.HandleFunc("/api/getallemployees", getEmployeesHandler(db))
	http.HandleFunc("/api/getone", getEmployeeByID(db))
	http.HandleFunc("/api/createemployee", createEmployeeHandler(db))

	log.Println("Сервер на :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}