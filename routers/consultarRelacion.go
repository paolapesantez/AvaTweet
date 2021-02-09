package routers

import (
	"encoding/json"
	"net/http"

	"github.com/paolapesantez/avatweet-server/bd"
	"github.com/paolapesantez/avatweet-server/models"
)

/*ConsultarRelacion chequea si hay relación entre dos usuarios*/
func ConsultarRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	var relacion models.Relacion

	relacion.UsuarioID = IDUsuario
	relacion.UsuarioRelacionID = ID

	//Defino una variable resp en donde voy a tener la respuesta de si hay relación o no entre los usuarios
	var resp models.RespuestaConsultaRelacion

	status, err := bd.BuscarRelacion(relacion)

	// No pierdo tiempo mostrando mensajes de error y me focalizo en responder true o false
	if err != nil || status == false {
		resp.Status = false
	} else {
		resp.Status = true
	}

	// Ahora seteo el Header indicando que la respuesta va a ser un json
	w.Header().Set("Content-type", "application/json")
	// escribo el header ahora con el status created
	w.WriteHeader(http.StatusCreated)
	// ahora hacemos nuestro encode para pasar a json nuestro modelo
	json.NewEncoder(w).Encode(resp)
}
