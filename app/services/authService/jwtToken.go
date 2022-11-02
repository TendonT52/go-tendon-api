package authService

// import (
// 	"time"

// 	"github.com/golang-jwt/jwt/v4"
// 	"github.com/spf13/viper"
// )

// func generateJWT(jwtId, userId string, curriculumId []string) (string, error) {
// 	exp := time.Now().Add(viper.GetDuration("token.accessTokenExpire")) 
// 	claim := models.AccessToken{
// 		jwt.RegisteredClaims{
// 			ID: jwtId,
// 			Subject : userId,
// 			Audience: curriculumId,
// 			Issuer: viper.GetString("app.name"),
// 			ExpiresAt: jwt.NewNumericDate(exp),
// 		},
// 	}
// 	jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
// }