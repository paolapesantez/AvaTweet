package bd

import (
	"context"
	"time"

	"github.com/paolapesantez/avatweet-server/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertarUsuario es la parada final con la BD para insertarlos datos del usuario*/
func InsertarUsuario(usuario models.Usuario) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("microblogging")
	col := db.Collection("usuarios")
	usuario.Password, _ = EncriptarPassword(usuario.Password)
	result, err := col.InsertOne(ctx, usuario)
	if err != nil {
		return "", false, err
	}
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
