package bd

import (
	"context"
	"time"

	"github.com/paolapesantez/avatweet/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ChequearSiExisteUsuario recibe un email de parámetro y chequea si ya está en la BD*/
func ChequearSiExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	//ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	db := MongoCN.Database("microblogging")
	col := db.Collection("usuarios")

	// M es una función que formatea o mapea a bson lo que recibe como json
	condicion := bson.M{"email": email}

	// en la variable usuarioEncontrado voy a modelar un usuario
	var usuarioEncontrado models.Usuario

	//FindOne me devuelve un sólo registro que cumple con la condición
	err := col.FindOne(ctx, condicion).Decode(&usuarioEncontrado)
	ID := usuarioEncontrado.ID.Hex()
	if err != nil {
		return usuarioEncontrado, false, ID
	}
	return usuarioEncontrado, true, ID
}
