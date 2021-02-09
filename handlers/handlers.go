package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/paolapesantez/avatweetServer/middlew"
	"github.com/paolapesantez/avatweetServer/routers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//Manejadores seteo mi puerto, el handler y pongo a escuchar al servidor
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequearBD(routers.RegistrarUsuario)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequearBD(routers.Login)).Methods("POST")
	router.HandleFunc("/verPerfil", middlew.ChequearBD(middlew.ValidarJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificarPerfil", middlew.ChequearBD(middlew.ValidarJWT(routers.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/subirAvatar", middlew.ChequearBD(middlew.ValidarJWT(routers.SubirAvatar))).Methods("POST")
	router.HandleFunc("/subirBanner", middlew.ChequearBD(middlew.ValidarJWT(routers.SubirBanner))).Methods("POST")
	router.HandleFunc("/obtenerAvatar", middlew.ChequearBD(routers.ObtenerAvatar)).Methods("GET")
	router.HandleFunc("/obtenerBanner", middlew.ChequearBD(routers.ObtenerBanner)).Methods("GET")

	router.HandleFunc("/tweet", middlew.ChequearBD(middlew.ValidarJWT(routers.EnviarTweet))).Methods("POST")
	router.HandleFunc("/leerTweets", middlew.ChequearBD(middlew.ValidarJWT(routers.LeerTweets))).Methods("GET")
	router.HandleFunc("/eliminarTweet", middlew.ChequearBD(middlew.ValidarJWT(routers.EliminarTweet))).Methods("DELETE")

	router.HandleFunc("/crearRelacion", middlew.ChequearBD(middlew.ValidarJWT(routers.CrearRelacion))).Methods("POST")
	router.HandleFunc("/eliminarRelacion", middlew.ChequearBD(middlew.ValidarJWT(routers.EliminarRelacion))).Methods("DELETE")
	router.HandleFunc("/consultarRelacion", middlew.ChequearBD(middlew.ValidarJWT(routers.ConsultarRelacion))).Methods("GET")

	router.HandleFunc("/listarUsuarios", middlew.ChequearBD(middlew.ValidarJWT(routers.ListarUsuarios))).Methods("GET")
	router.HandleFunc("/leerTweetsSeguidores", middlew.ChequearBD(middlew.ValidarJWT(routers.LeerTweetsSeguidores))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
