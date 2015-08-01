//this is a PoC, en experiment to structure a REST API Serviec backed by a Database
//currently mongo
package main

import (
	"log"
	"time"

	"github.com/switching/go-web-poc/dbService"
	"github.com/switching/go-web-poc/handlers"
	"github.com/zenazn/goji"
	"gopkg.in/mgo.v2"
)

func main() {
	url := "127.0.0.1"

	timeOut := time.Duration(2) * time.Second
	log.Printf("connecting to mongo=%v, timeout=%v", url, timeOut)
	mgoMasterSession, err := mgo.DialWithTimeout(url, timeOut)

	if err != nil {
		log.Fatalf("error connecting to mongo=%v, error=%v", url, err)
	}

	dbName := "dev"

	// Setup endpoints
	// add middleware here
	db := dbService.New(mgoMasterSession, dbName, "users")
	uh := handlers.NewUserHandler(db)
	goji.Post("/register", uh.Create)

	goji.Serve()
}
