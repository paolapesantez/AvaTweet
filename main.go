package main

import (
	"log"

	"github.com/paolapesantez/avatweet/bd"
	"github.com/paolapesantez/avatweet/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()

}
