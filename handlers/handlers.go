package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/paolapesantez/avatweet/middlew"
	"github.com/paolapesantez/avatweet/routers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//Manejadores seteo mi puerto, el handler y pongo a escuchar al servidor
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequearBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequearBD(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middlew.ChequearBD(middlew.ValidarJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificarperfil", middlew.ChequearBD(middlew.ValidarJWT(routers.ModificarPerfil))).Methods("PUT")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
