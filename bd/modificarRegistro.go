package bd

import (
	"context"
	"time"

	"github.com/paolapesantez/avatweet/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*ModificarRegistro permite modificar el perfil del usuario*/
func ModificarRegistro(usuario models.Usuario, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("microblogging")
	col := db.Collection("usuarios")

	/* supongo que me van a enviar de a un campo a modificar a la vez,
	   por eso me fijo si lo que viene tiene valor (largo mayor a cero)
	   Creamos un mapa de interfaces para armar el registro de actualización a la bd
	   Le pongo la información que hay que modificar*/
	registro := make(map[string]interface{})

	if len(usuario.Nombre) > 0 {
		registro["nombre"] = usuario.Nombre
	}
	if len(usuario.Apellidos) > 0 {
		registro["apellidos"] = usuario.Apellidos
	}
	registro["fechaNacimiento"] = usuario.FechaNacimiento
	if len(usuario.Avatar) > 0 {
		registro["avatar"] = usuario.Avatar
	}
	if len(usuario.Banner) > 0 {
		registro["banner"] = usuario.Banner
	}
	if len(usuario.Biografia) > 0 {
		registro["biografia"] = usuario.Biografia
	}
	if len(usuario.Ubicacion) > 0 {
		registro["ubicacion"] = usuario.Ubicacion
	}
	if len(usuario.SitioWeb) > 0 {
		registro["sitioWeb"] = usuario.SitioWeb
	}
	updtString := bson.M{
		"$set": registro,
	}
	objID, _ := primitive.ObjectIDFromHex(ID)

	// ahora tengo que realizar un filtro con el ID
	filtro := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filtro, updtString)
	if err != nil {
		return false, err
	}
	return true, nil

}
