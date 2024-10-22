package routers

import (
	"errors"
	"strings"
	"twitter/bd"
	"twitter/models"

	"github.com/dgrijalva/jwt-go"
)

var Email string

var IDUsuarios string

func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("MastersdelDesarrollo")

	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")

	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("Formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(t *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)

		if encontrado == true {
			Email = claims.Email
			IDUsuarios = claims.ID.Hex()
		}

		return claims, encontrado, IDUsuarios, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("Token Invalido")
	}

	return claims, false, string(""), err
}
