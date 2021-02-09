package routers

import (
	"encoding/json"
	"net/http"

	"github.com/paolapesantez/avatweet-server/bd"
	"github.com/paolapesantez/avatweet-server/models"
)

/*ModificarPerfil - modifica el perfil de usuario*/
func ModificarPerfil(w http.ResponseWriter, r *http.Request) {

	var usuario models.Usuario

	err := json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil {
		// es un Json mal construído
		http.Error(w, "Datos incorrectos "+err.Error(), 400)
		return
	}

	var status bool
	status, err = bd.ModificarRegistro(usuario, IDUsuario)
	//IDUsuario es la variable global que seteamos antes con el ID
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar modificar el registro. Intente nuevamente "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado modificar elregistro del usuario "+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
