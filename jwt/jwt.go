package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/mxaxaxbx/twitter-clone/models"
)

func GeneroJWT(t models.Usuario) (string, error) {
	miclave := []byte("miclave_secreta")

	payload := jwt.MapClaims{
		"email":            t.Email,
		"nombre":           t.Nombre,
		"apellidos":        t.Apellidos,
		"fecha_nacimiento": t.FechaNacimiento,
		"biografia":        t.Biografia,
		"sitioweb":         t.SitioWeb,
		"_id":              t.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miclave)

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
