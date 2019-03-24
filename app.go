package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {
	log.Println("Calling initialize function")
	connectionString := fmt.Sprintf("%s:%s@/%s", user, password, dbname)
	var err error
	a.DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = a.DB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
	a.Router = mux.NewRouter()
	log.Println("Done calling")
}

func (a *App) Run(addr string) {
	a.Router.HandleFunc("/users", a.GetUsers).Methods("GET")
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

var users []user

func (a *App) GetUsers(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println(params)
	users = append(users, user{1, "Chetan", 36})
	users = append(users, user{2, "Sonali", 36})
	users = append(users, user{3, "Sofie", 4})
	users, error := getUsers(a.DB, 0, 10)
	if error != nil {
		respondWithError(w, 500, error.Error())
		return
	}
	respondWithJSON(w, 200, users)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}
