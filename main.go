package main

import (
	"log"

	"github.com/mxaxaxbx/twitter-clone/bd"
	"github.com/mxaxaxbx/twitter-clone/handlers"
)

func main() {
	if bd.ChequeoConection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()
}
