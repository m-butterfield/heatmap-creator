package main

import (
	"github.com/m-butterfield/heatmap-creator/server/app/controllers"
	"log"
)

func main() {
	if err := controllers.Run("8000"); err != nil {
		log.Fatal(err)
	}
}
