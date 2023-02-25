package main

import (
	"github.com/sk10az/go_anime_crud/internal/app"
	"log"
)

func main() {
	app, err := app.New()
	if err != nil {
		log.Fatal("App can't be run")
	}

	log.Fatal(app.Run())
}
