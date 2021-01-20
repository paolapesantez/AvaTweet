package routers

import (
	"encoding/json"
	"net/http"

	"github.com/paolapesantez/avatweet/bd"
	"github.com/paolapesantez/avatweet/models"
)

/*Registro es la función para crear en la BD el registro de usuario*/
func Registro(w http.ResponseWriter, r *http.Request) {

	var usuario models.Usuario
	err := json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil {
		http.Error(w, "Error en los datos recibidos: "+err.Error(), 400)
		return
	}
	/*Si no hubo error con el Body hago unas validaciones*/
	if len(usuario.Email) == 0 {
		http.Error(w, "El email de usuario es requerido ", 400)
		return
	}
	if len(usuario.Password) < 6 {
		http.Error(w, "Debe especificar una contraseña de al menos seis caracteres", 400)
		return
	}
	_, encontrado, _ := bd.ChequeoYaExisteUsuario(usuario.Email)
	if encontrado == true {
		http.Error(w, "Ya existe un usuario registrado con ese Email", 400)
		return
	}
	_, status, err := bd.InsertoRegistro(usuario)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar realizar el registro de usuario"+err.Error(), 400)
		return
	}
	/*Si llegó hasta acá todo anduvo bien*/
	if status == false {
		http.Error(w, "No se ha logrado insertar el registro del usuario", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
