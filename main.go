package main

import (
	"log"

	"github.com/paolapesantez/avatweetServer/bd"
	"github.com/paolapesantez/avatweetServer/handlers"
)

func main() {
	if bd.ChequearConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()

}
