package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go" //Creo un alias para manejarlo más fácil
	"github.com/paolapesantez/avatweet-server/models"
)

func GenerarJWT(usuario models.Usuario) (string, error) {
	miClave := []byte("SkillFactoryGo_Avalith")

	payload := jwt.MapClaims{
		"_id":             usuario.ID.Hex(),
		"email":           usuario.Email,
		"nombre":          usuario.Nombre,
		"apellidos":       usuario.Apellidos,
		"fechaNacimiento": usuario.FechaNacimiento,
		"biografia":       usuario.Biografia,
		"ubicacion":       usuario.Ubicacion,
		"sitioWeb":        usuario.SitioWeb,
		"exp":             time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
