package handler

import (
	"fmt"
	"net/http"

	"github.com/fazilnbr/project-workey/pkg/config"
	"github.com/fazilnbr/project-workey/pkg/domain"
	services "github.com/fazilnbr/project-workey/pkg/usecase/interface"
	"github.com/fazilnbr/project-workey/pkg/utils"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	adminUseCase  services.AdminUseCase
	workerUseCase services.WorkerUseCase
	userUseCase   services.UserUseCase
	jwtUseCase    services.JWTUseCase
	authUseCase   services.AuthUseCase
	cfg           config.Config
}

func NewAuthHandler(
	adminUseCase services.AdminUseCase,
	workerUseCase services.WorkerUseCase,
	userusecase services.UserUseCase,
	jwtUseCase services.JWTUseCase,
	authUseCase services.AuthUseCase,
	cfg config.Config,

) AuthHandler {
	return AuthHandler{
		adminUseCase:  adminUseCase,
		workerUseCase: workerUseCase,
		userUseCase:   userusecase,
		jwtUseCase:    jwtUseCase,
		authUseCase:   authUseCase,
		cfg:           cfg,
	}
}

// @Summary SignUp for users
// @ID SignUp authentication
// @Tags User Authentication
// @Produce json
// @Tags User Authentication
// @Param WorkerLogin body domain.User{username=string,password=string} true "Worker Login"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/sent-otp [post]
func (cr *AuthHandler) UserSendOTP(ctx *gin.Context) {
	fmt.Printf("\n\nuser  :  \n\n")
	var newUser domain.Signup

	err := ctx.Bind(&newUser)
	if err != nil {
		response := utils.ErrorResponse("Failed to create user", err.Error(), nil)
		ctx.Writer.Header().Set("Content-Type", "application/json")
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*ctx, response)
		return
	}
	phoneNumber := fmt.Sprintf(newUser.CountryCode + newUser.PhoneNumber)
	_, err = cr.userUseCase.RegisterAndVarify(ctx, phoneNumber)

	if err != nil {
		response := utils.ErrorResponse("Failed to create user", err.Error(), nil)
		ctx.Writer.Header().Set("Content-Type", "application/json")
		ctx.Writer.WriteHeader(http.StatusUnprocessableEntity)
		utils.ResponseJSON(*ctx, response)
		return
	}

	response := utils.SuccessResponse(true, "SUCCESS", nil)
	ctx.Writer.Header().Set("Content-Type", "application/json")
	ctx.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*ctx, response)
}

// @Summary SignUp for users
// @ID SignUp authentication
// @Tags User Authentication
// @Produce json
// @Tags User Authentication
// @Param WorkerLogin body domain.User{username=string,password=string} true "Worker Login"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/login [post]
func (cr *AuthHandler) UserLogin(ctx *gin.Context) {
	var newUser domain.Signup

	err := ctx.Bind(&newUser)
	if err != nil {
		response := utils.ErrorResponse("Failed to create user", err.Error(), nil)
		ctx.Writer.Header().Set("Content-Type", "application/json")
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*ctx, response)
		return
	}
	phoneNumber := fmt.Sprintf(newUser.CountryCode + newUser.PhoneNumber)
	userId, err := cr.userUseCase.RegisterAndVarify(ctx, phoneNumber)

	if err != nil {
		response := utils.ErrorResponse("Failed to create user", err.Error(), nil)
		ctx.Writer.Header().Set("Content-Type", "application/json")
		ctx.Writer.WriteHeader(http.StatusUnprocessableEntity)
		utils.ResponseJSON(*ctx, response)
		return
	}

	accessToken, err := cr.jwtUseCase.GenerateAccessToken(userId, "", "user")
	if err != nil {
		response := utils.ErrorResponse("Failed to generate access token", err.Error(), nil)
		ctx.Writer.Header().Add("Content-Type", "application/json")
		ctx.Writer.WriteHeader(http.StatusUnauthorized)
		utils.ResponseJSON(*ctx, response)
		return
	}

	refreshToken, err := cr.jwtUseCase.GenerateRefreshToken(userId, "", "user")

	if err != nil {
		response := utils.ErrorResponse("Failed to generate refresh token please login again", err.Error(), nil)
		ctx.Writer.Header().Add("Content-Type", "application/json")
		ctx.Writer.WriteHeader(http.StatusUnauthorized)
		utils.ResponseJSON(*ctx, response)
		return
	}

	userResponse := domain.UserResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	response := utils.SuccessResponse(true, "SUCCESS", userResponse)
	ctx.Writer.Header().Set("Content-Type", "application/json")
	ctx.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*ctx, response)
}
