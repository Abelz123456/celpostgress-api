package middleware

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	jWTPrivateToken = "SecrteTokenSecrteToken"
	// ip              = "192.168.0.107"
)

type Maker struct {
	token *jwt.Token
	// response common.WebResponse
}

func NewJwt() (*Maker, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	maker := &Maker{
		token: token,
	}

	return maker, nil
}

func (maker *Maker) CreateToken(emailName string) (string, error) {
	claims := maker.token.Claims.(jwt.MapClaims)
	claims["username"] = emailName
	claims["level"] = "application"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, err := maker.token.SignedString([]byte("secret"))
	return t, err

}

func (maker *Maker) VerifyToken(accessToken string) (*jwt.Token, error) {

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	// _, err := utils.ParseToken(accessToken)

	// if err != nil {
	// 	switch err.(*jwt.ValidationError).Errors {
	// 	case jwt.ValidationErrorExpired:
	// 		maker.response = common.WebResponse{
	// 			Code:   200,
	// 			Status: "OK",
	// 			Data:   "siteResponse",
	// 		}
	// 	default:
	// 		maker.response = common.WebResponse{
	// 			Code:   200,
	// 			Status: "OK",
	// 			Data:   "siteResponse",
	// 		}
	// 	}
	// }
	return token, err

}

func (maker *Maker) GetClaims(tokenString string) JwtClaims {
	claims := &JwtClaims{}

	_, err := getTokenFromString(tokenString, claims)
	if err == nil {
		return *claims
	}
	return *claims
}
func getTokenFromString(tokenString string, claims *JwtClaims) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(jWTPrivateToken), nil
	})
}
