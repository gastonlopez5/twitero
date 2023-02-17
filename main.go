package main

import (
	"log"

	"githut.com/gastonlopez5/twitero/bd"
	"githut.com/gastonlopez5/twitero/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}

	handlers.Manejadores()
}
