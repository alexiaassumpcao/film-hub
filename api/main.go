package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type Film struct {
	Name        string  `json:"name"`
	ReleaseDate string  `json:"release_date"`
	Ranking     float32 `json:"ranking"`
	RunningTime int     `json:"running_time"`
}
type User struct {
	Name           string `json:"name"`
	AccessQuestion string `json:"access_question"`
	AccessAnswer   string `json:"access_answer"`
	FavoriteFilm   string `json:"favorite_film"`
}

const (
	HOST     = "localhost"
	PORT     = 5432
	USER     = "postgres"
	PASSWORD = "21071292"
	DBNAME   = "film-hub-db"
)

func OpenConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", HOST, PORT, USER, PASSWORD, DBNAME)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	return db
}
func GETHandler(w http.ResponseWriter, r *http.Request) {
	db := OpenConnection()
	q := `

select
	name
from
	films
;`
	rows, err := db.Query(q)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var films []Film
	for rows.Next() {
		var film Film
		rows.Scan(&film.Name)
		films = append(films, film)
	}
	filmsBytes, _ := json.MarshalIndent(films, "", "\t")
	w.Header().Set("Content-Type", "application/json")
	w.Write(filmsBytes)
	defer db.Close()
}

func POSTHandler(w http.ResponseWriter, r *http.Request) {
	db := OpenConnection()
	var f Film
	err := json.NewDecoder(r.Body).Decode(&f)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sqlStatement := `INSERT INTO films (release_date, name, ranking, running_time) VALUES 
	($1, $2, $3, $4);`
	_, err = db.Exec(sqlStatement, f.ReleaseDate, f.Name, f.Ranking, f.RunningTime)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	defer db.Close()
}

func main() {
	http.HandleFunc("/", GETHandler)
	http.HandleFunc("/insert", POSTHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
