package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	service "github.com/fazilnbr/project-workey/pkg/usecase/interface"
	"github.com/fazilnbr/project-workey/pkg/utils"
	response "github.com/fazilnbr/project-workey/pkg/utils"
	"github.com/gin-gonic/gin"
)

type Middleware interface {
	AthoriseJWT(*gin.Context)
}

type middlewar struct {
	jwtUseCase service.JWTUseCase
}

// AthoriseJWT implements Middileware

func (cr *middlewar) AthoriseJWT(c *gin.Context) {
	//getting from header
	autheader := c.Request.Header["Authorization"]
	auth := strings.Join(autheader, " ")
	bearerToken := strings.Split(auth, " ")

	// fmt.Print("\n\n\nff\n\n\ntoken : ", len(bearerToken), ":  :", bearerToken)

	if len(bearerToken) != 2 {
		err := errors.New("request does not contain an access token")
		response := response.ErrorResponse("Failed to create user", err.Error(), nil)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnauthorized)

		utils.ResponseJSON(*c, response)
		c.Abort()
		return
	}

	authtoken := bearerToken[1]
	// fmt.Print(authtoken)
	ok, claims := cr.jwtUseCase.VerifyToken(authtoken)
	source := fmt.Sprint(claims.Source)

	fmt.Printf("\n\nok : %v\n\n", time.Now().Unix()-claims.ExpiresAt)
	fmt.Printf("\n\nsorce : %v\n\n", source)

	if !ok {
		err := errors.New("your access token is not valid")
		response := response.ErrorResponse("Error", err.Error(), source)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnauthorized)
		utils.ResponseJSON(*c, response)
		c.Abort()
		return
	}

	if source != "accesstoken" {
		err := errors.New("The token not an access token")
		response := response.ErrorResponse("Error", err.Error(), source)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnauthorized)
		utils.ResponseJSON(*c, response)
		c.Abort()
		return
	}

	user_email := fmt.Sprint(claims.UserName)
	id := fmt.Sprint(claims.UserId)
	fmt.Printf("\n\nid : %v\n\n", id)

	c.Writer.Header().Set("email", user_email)
	c.Writer.Header().Set("id", id)

}

func NewUserMiddileware(jwtUserUseCase service.JWTUseCase) Middleware {
	return &middlewar{
		jwtUseCase: jwtUserUseCase,
	}
}
func NewWorkerMiddileware(jwtWorkerUsecase service.JWTUseCase) Middleware {
	return &middlewar{
		jwtUseCase: jwtWorkerUsecase,
	}
}
func NewAdminMiddileware(jwtAdminUseCase service.JWTUseCase) Middleware {
	return &middlewar{
		jwtUseCase: jwtAdminUseCase,
	}
}
