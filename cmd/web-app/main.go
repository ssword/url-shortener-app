package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/ssword/url-shortener-app/internal/controllers"
	"github.com/ssword/url-shortener-app/internal/db"
	"log"
	"net/http"
)

func main() {
	slite, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer slite.Close()

	if err := db.CreateTable(slite); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path == "/" {
			controllers.ShowIndex(writer, request)
		} else {
			controllers.Proxy(slite)(writer, request)
		}
	})
	http.HandleFunc("/shorten", controllers.Shorten(slite))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
