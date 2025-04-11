package main

import (
	"log"

	"github.com/starnuik/avito_pvz/pkg/app"
)

func main() {
	app, err := app.New()
	if err != nil {
		log.Fatal(err)
	}
	app.Run()
}
