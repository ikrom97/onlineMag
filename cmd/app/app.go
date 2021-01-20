package app

import (
	"database/sql"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"onlineMag/db"
	"onlineMag/pkg/core/services"
)

type MainServer struct {
	Db      *sql.DB
	router  *httprouter.Router
	usersvc *services.UserSvc
}

func NewMainServer(Db *sql.DB, router *httprouter.Router, usersvc *services.UserSvc) *MainServer {
	return &MainServer{Db: Db, router: router, usersvc: usersvc}
}
func (server *MainServer) Start(Db *sql.DB) {
	err := db.DatabaseInit(Db)
	if err != nil {
		log.Fatal("Can't init database err = ", err)
	}
	server.InitRoutes()
}
func (server *MainServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	server.router.ServeHTTP(writer, request)
}
