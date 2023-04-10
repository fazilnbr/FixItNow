package usecase

import (
	"fmt"
	"log"
	"os"
	"time"

	model "github.com/fazilnbr/project-workey/pkg/domain"
	services "github.com/fazilnbr/project-workey/pkg/usecase/interface"
	"github.com/golang-jwt/jwt/v4"
)

type JWTUseCase struct {
	SecretKey string
}

// GenerateRefreshToken implements interfaces.JWTUsecase
func (j *JWTUseCase) GenerateRefreshToken(userid int, username string, role string) (string, error) {
	claims := &model.SignedDetails{
		UserId:   userid,
		UserName: username,
		Source:   "refreshtoken",
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * 12 * 7).Unix(),
		},
	}
	// fmt.Printf("\n\nrefresh time : %v\n\n", time.Hour*12*7)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(j.SecretKey))

	if err != nil {
		log.Println(err)
	}

	return signedToken, err
}

// GenerateRefreshToken implements interfaces.JWTUsecase
func (j *JWTUseCase) GenerateAccessToken(userid int, username string, role string) (string, error) {

	claims := &model.SignedDetails{
		UserId:   userid,
		UserName: username,
		Source:   "accesstoken",
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Minute * time.Duration(5)).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(j.SecretKey))

	if err != nil {
		log.Println(err)
	}

	return signedToken, err
}

// // GetTokenFromString implements interfaces.JWTUseCase
func (j *JWTUseCase) GetTokenFromString(signedToken string, claims *model.SignedDetails) (*jwt.Token, error) {
	return jwt.ParseWithClaims(signedToken, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(j.SecretKey), nil
	})

}

// VerifyToken implements interfaces.JWTUseCase
func (j *JWTUseCase) VerifyToken(signedToken string) (bool, *model.SignedDetails) {
	claims := &model.SignedDetails{}
	token, err := j.GetTokenFromString(signedToken, claims)

	if err != nil {
		return false, claims
	}

	if token.Valid {
		if e := claims.Valid(); e == nil {
			return true, claims
		}
	}
	return false, claims
}

func NewJWTUserService() services.JWTUseCase {
	return &JWTUseCase{
		SecretKey: os.Getenv("USER_KEY"),
	}
}
