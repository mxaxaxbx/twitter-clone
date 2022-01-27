package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/mxaxaxbx/twitter-clone/middlewares"
	"github.com/mxaxaxbx/twitter-clone/routes"
	"github.com/rs/cors"
)

/* manejadores, seteo de rutas y manejo de errores */
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlewares.ChequeoBD(routes.Registro)).Methods("POST")
	router.HandleFunc("/login", middlewares.ChequeoBD(routes.Login)).Methods("POST")
	router.HandleFunc("/verPerfil", middlewares.ChequeoBD(middlewares.ValidoJWT(routes.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificarPerfil", middlewares.ChequeoBD(middlewares.ValidoJWT(routes.ModificarPerfil))).Methods("PUT")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8000"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
