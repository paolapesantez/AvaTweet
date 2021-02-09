package bd

import (
	"github.com/paolapesantez/avatweetServer/models"
	"golang.org/x/crypto/bcrypt"
)

/*IntentarLogin realiza el chequeo de login de un usuario contra la BD*/
func IntentarLogin(email string, password string) (models.Usuario, bool) {
	usuarioEncontrado, encontrado, _ := BuscarUsuario(email)
	if encontrado == false {
		return usuarioEncontrado, false
	}
	// Ahora comparo la password que en la BD está encriptada
	// creo una variable que sea un slice de bytes
	passwordBytes := []byte(password)
	// creo otra variable con la password que tengo en la BD para el usuario
	passwordBD := []byte(usuarioEncontrado.Password)
	// Ahora llamo a una función del package bcrypt que compara las password
	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err != nil {
		return usuarioEncontrado, false
	}
	return usuarioEncontrado, true
}
