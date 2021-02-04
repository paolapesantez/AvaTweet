package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/paolapesantez/avatweet/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*BuscarRelacion - consulta relacion entre dos usuarios*/
func BuscarRelacion(relacion models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("microblog")
	col := db.Collection("relacion")

	/*Mapeo a bson la relación relacion que viene como parámetro para
	  buscar en la base de datos*/
	condicion := bson.M{
		"usuarioid":         relacion.UsuarioID,
		"usuariorelacionid": relacion.UsuarioRelacionID,
	}

	//defino un modelo para contener el resultado de la consulta
	var resultado models.Relacion
	err := col.FindOne(ctx, condicion).Decode(&resultado)

	// Imprimo por pantalla (saldrá en la terminal) el resultado de la consulta
	fmt.Println(resultado)

	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	return true, nil
}
