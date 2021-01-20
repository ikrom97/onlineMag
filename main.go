package main

import (
	"database/sql"
	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"onlineMag/cmd/app"
	"onlineMag/pkg/core/services"
)

func main() {
	DB, err := sql.Open("sqlite3", "mag")
	if err != nil {
		log.Fatal("Can't find sql connection", err)
	}
	router := httprouter.New()
	svc := services.NewUserSvc(DB)
	server := app.NewMainServer(DB, router, svc)
	server.Start(DB)
}
