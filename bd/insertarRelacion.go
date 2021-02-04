package bd

import (
	"context"
	"time"

	"github.com/paolapesantez/avatweet/models"
)

/*InsertarRelacion - graba la relación en la bd*/
func InsertarRelacion(relacion models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("microblogging")
	col := db.Collection("relaciones")

	//Directamente trato de insertar en la base de datos el modelo que recibo
	_, err := col.InsertOne(ctx, relacion)
	if err != nil {
		return false, err
	}
	return true, nil
}
