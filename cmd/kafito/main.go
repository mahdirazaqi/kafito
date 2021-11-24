package main

import (
	"flag"
	"log"

	"github.com/mahdirazaqi/kafito/config"
)

var c = flag.String("c", "./config.json", "Config File Path")

func main() {
	flag.Parse()

	if err := config.Load(*c); err != nil {
		log.Fatal(err)
	}
}
