package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/paolapesantez/avatweet-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*ListarUsuariosCriterio - Lee los usuarios registrados en el sistema,
si se recibe "R" en quienes trae sólo los que se relacionan conmigo
  Tiene muchos parámetros:
  - el ID, del usuario que está leyendo los usuarios
  - la página porque si hay muchos usuarios hay que paginar
  - el search que es una palabra para la búsqueda(todos los usuarios que contengan la palabra search
  - el tipo de búsqueda que vamos a hacer, si listamos todos, sólo los que nos siguen, o los que seguimos
  devuelve un slice de tipo models usuario, ya que devolvemos n usuarios que responden a la búsqueda
  y un booleano */
func ListarUsuariosCriterio(ID string, page int64, search string, tipo string) ([]*models.Usuario, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("microblogging")
	col := db.Collection("usuarios")

	//creamos la variable para el resultado, que va a ser un slice de usuarios
	var results []*models.Usuario

	findOptions := options.Find()
	// ahora le damos las propiedades a nuestro find
	//defino el salto para posicionarme en los resultados
	findOptions.SetSkip((page - 1) * 20)
	//devuelvo de a 20 resultados
	findOptions.SetLimit(20)

	/*Ahora haremos la condicion con un mapeo de bson,
	  con el nombre de tipo regex (expresión regular)
	  ?i hace que no distinga entre mayusculas y minusculas */
	query := bson.M{
		"nombre": bson.M{"$regex": `(?i)` + search},
	}
	//Ahora ejecutamos el find que devuelve el resultado en un cursor
	cursor, err := col.Find(ctx, query, findOptions)
	if err != nil {
		// si no fue satisfactoria la búsqueda en la BD
		fmt.Println(err.Error())
		//Devuelve results vacio y false
		return results, false
	}

	//Vamos a definir dos variables booleanas
	var encontrado, incluir bool

	// El next me permite avanzar al siguiente registro
	for cursor.Next(ctx) {

		//defino un usuario para trabajar con el elemento que puedo después incluir en results
		var usuario models.Usuario
		//grabamos lo del cursor en el modelo de usuario para leer los campos
		err := cursor.Decode(&usuario)
		if err != nil {
			//encontró un error
			fmt.Println(err.Error())
			return results, false
		}

		//creo una variable relacion para consultar sobre la relacion con el usuario
		var relacion models.Relacion
		relacion.UsuarioID = ID
		relacion.UsuarioRelacionID = usuario.ID.Hex()

		//por cada iteracion tengo que ver si al usuario lo debo incluir o no en el resultado
		incluir = false

		encontrado, err = BuscarRelacion(relacion)

		if tipo == "new" && encontrado == false {
			//lo tengo que incluir en la lista, porque no lo encontró en la relación
			incluir = true
		}
		//si tipo es follow sólo quiero listar los que yo estoy siguiendo
		if tipo == "follow" && encontrado == true {
			incluir = true
		}
		if relacion.UsuarioRelacionID == ID {
			//sería el caso de que me estoy siguiendo a mí mismo
			incluir = false
		}
		if incluir == true {
			//hago un blanqueo de los campos que no me interesa incluir
			//solo quiero el avatar, el nombre, los apellidos y la fecha de nacimiento
			usuario.Password = ""
			usuario.Biografia = ""
			usuario.SitioWeb = ""
			usuario.Ubicacion = ""
			usuario.Banner = ""
			usuario.Email = ""

			//agrego con append el modelo de usuario usuario en results
			results = append(results, &usuario)
		}
	}

	//me fijo si hubo un error del cursor
	err = cursor.Err()
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}
	//si no hubo error, cierro el cursor y retorno la lista de resultados
	cursor.Close(ctx)
	return results, true
}
