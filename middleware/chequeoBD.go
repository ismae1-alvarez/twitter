package middleware

import (
	"net/http"
	"twitter/bd"
)

/*CequeoBD es el middleware que permite ver el estado de la base de datos*/

func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Coxion perdida con la base de datos", 500)
			return
		}

		next.ServeHTTP(w, r)
	}
}
