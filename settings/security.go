package settings

import (
	"fly-go/models"
	"fmt"
	"net/http"

	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

var SECRET = []byte("R@gnarok2")

// gera token
func GenerateJWT(passageiro models.Passageiro) (string, error) {

	claims := models.Claim{
		Passageiro: passageiro,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour + 1).Unix(),
			Issuer:    "Teste jwt",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := token.SignedString(SECRET)
	if err != nil {
		fmt.Println("não se pode montar o token")
		fmt.Println(err)
		return "", err
	}

	return result, nil
}

// valida se o token passado no header é válido
func ValidateToken(w http.ResponseWriter, r *http.Request) {

	token, err := request.ParseFromRequestWithClaims(r, request.OAuth2Extractor, &models.Claim{},
		func(token *jwt.Token) (interface{}, error) {
			return SECRET, nil
		})
	if err != nil {
		switch err.(type) {
		case *jwt.ValidationError:
			vErr := err.(*jwt.ValidationError)
			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				fmt.Println("token expirado")
				return
			case jwt.ValidationErrorSignatureInvalid:
				fmt.Println("Dados invalido, token informado não coincide")
				return
			default:
				fmt.Println("Token invalido")
				return

			}
		default:
			fmt.Println("Token informado invalido")
			return

		}
	}

	if token.Valid {
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprint(w, "bem vindo ao sistema")
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Token não autorizado")
	}
}
