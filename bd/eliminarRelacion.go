package bd

import (
	"context"
	"time"

	"github.com/paolapesantez/avatweetServer/models"
)

/*EliminarRelacion borra la relación en la bd */
func EliminarRelacion(relacion models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("microblogging")
	col := db.Collection("relaciones")

	_, err := col.DeleteOne(ctx, relacion)
	if err != nil {
		return false, err
	}
	return true, nil
}
