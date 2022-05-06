package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	baus "github.com/r4wm/baus/internal"
)

const (
	logFile = "/tmp/baus.log"
)

func main() {
	logFile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: ", err)
	}
	defer logFile.Close()

	logger := log.New(logFile, "baus: ", log.LstdFlags)

	app := baus.App{
		Router: mux.NewRouter().StrictSlash(false),
		Logger: logger,
	}

	app.SetupRouter()
	app.ConsumeJsonBibleFile("/tmp/kjv.json")
	log.Fatal(http.ListenAndServe(":8080", app.Router))
	fmt.Println(baus.Something())
}
