package middlewares

import (
	"net/http"

	"github.com/mxaxaxbx/twitter-clone/bd"
)

func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConection() == 0 {
			http.Error(w, "Sin conexion a la BD", 500)
			return
		}

		next.ServeHTTP(w, r)
	}
}
