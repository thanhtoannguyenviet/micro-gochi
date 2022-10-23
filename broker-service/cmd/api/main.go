package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

const webPort = "80"

var counts int64

type Config struct {
	DB *sql.DB
}

func main() {
	log.Println("Starting broker service")
	app := Config{}
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
