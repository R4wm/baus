package main

import (
	"encoding/json"
	"flag"
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

var (
	useColor = false
	RED      = "\033[0;31m"
	NC       = "\033[0m" // No Color"

)

// Config the file with info
type Config struct {
	Host    string `json:"host"`
	Port    string `json:"port"`
	KJVPath string `json:"kjvPath"`
	LogPath string `json:"logPath"`
	Color   bool   `json:"color"`
}

func main() {
	config := flag.String("config", "", "json configuration file")
	flag.Parse()
	if *config == "" {
		log.Fatalf("config file required")
	}
	/////////////////////////
	// consume config file //
	/////////////////////////
	dat, err := os.ReadFile(*config)
	if err != nil {
		log.Fatalf("could not read config: %s\n", *config)
	}
	configData := &Config{}
	log.Println(fmt.Sprintf("configData: %#v\n", configData))
	err = json.Unmarshal(dat, &configData)
	if err != nil {
		log.Fatalf("Failed to unmarshal config from: %s\n%s\n", *config, err)
	}
	listenAddr := configData.Host + ":" + configData.Port
	logFile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: ", err)
	}
	defer logFile.Close()
	logger := log.New(logFile, "baus: ", log.LstdFlags)
	///////////////////
	// router	 //
	///////////////////
	app := baus.App{
		Router: mux.NewRouter().StrictSlash(false),
		Logger: logger,
	}
	app.SetupRouter()
	app.ConsumeJSONBibleFile(configData.KJVPath)
	log.Fatal(http.ListenAndServe(listenAddr, app.Router))
}
