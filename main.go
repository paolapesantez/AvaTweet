package main

import (
	"log"

	"github.com/paolapesantez/avatweet-server/bd"
	"github.com/paolapesantez/avatweet-server/handlers"
)

func main() {
	if bd.ChequearConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()

}
