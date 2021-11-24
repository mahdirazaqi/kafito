package main

import (
	"flag"
	"log"

	"github.com/mahdirazaqi/kafito/config"
	"github.com/mahdirazaqi/kafito/database"
)

var c = flag.String("c", "./config.json", "Config File Path")

func main() {
	flag.Parse()

	if err := config.Load(*c); err != nil {
		log.Fatal(err)
	}

	if err := database.Connect(config.C.Mongo.Host, config.C.Mongo.DB, config.C.Mongo.User, config.C.Mongo.Password); err != nil {
		log.Fatal(err)
	}
}
