package helper

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(id int, email, secret string) (string, error) {
	claims := jwt.MapClaims{}
	claims["exp"] = time.Now().Add((30 * time.Minute)).Unix() // kapan token gabisa dipake
	claims["iat"] = time.Now().Unix()                         // kapan token dibuat
	claims["nbf"] = time.Now().Unix()                         // kapan token bisa digunakan == iat
	claims["email"] = email
	claims["id"] = id
	claims["aud"] = "fadhil.service"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
