package interfaces

import (
	model "github.com/fazilnbr/project-workey/pkg/domain"
	"github.com/golang-jwt/jwt/v4"
)

type JWTUseCase interface {
	GenerateRefreshToken(userid int, username string, role string) (string, error)
	GenerateAccessToken(userid int, username string, role string) (string, error)
	VerifyToken(signedToken string) (bool, *model.SignedDetails)
	GetTokenFromString(signedToken string, claims *model.SignedDetails) (*jwt.Token, error)
}
