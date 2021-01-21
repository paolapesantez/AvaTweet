package middlew

import (
	"net/http"

	"github.com/paolapesantez/avatweet/bd"
)

//ChequearBD para ...
func ChequearBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequearConnection() == 0 {
			http.Error(w, "Conexion perdida con la base de datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
