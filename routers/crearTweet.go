package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/paolapesantez/avatweet/bd"
	"github.com/paolapesantez/avatweet/models"
)

/*CrearTweet permite grabar el tweet en la BD */
func CrearTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet

	// decodificamos el body y armamos un registro
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.TweetUser{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}
	// Para insertarlo en la base de datos necesitamos mapearlo a un bson
	_, status, err := bd.InsertarTweet(registro)

	//si hay un error
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar insertar el registro, intentelo nuevamente.  "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado insertar el tweet.", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}