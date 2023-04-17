package handler

import (
	"net/http"
	"strconv"

	"github.com/fazilnbr/project-workey/pkg/domain"
	services "github.com/fazilnbr/project-workey/pkg/usecase/interface"
	"github.com/fazilnbr/project-workey/pkg/utils"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUseCase services.UserUseCase
}

// @Summary Get User Profile
// @ID GetUserProfile
// @Tags User Profile Management
// @Produce json
// @Security BearerAuth
// @Success 200 {object} utils.Response{}
// @Failure 422 {object} utils.Response{}
// @Router /user/profile [post]
func (c *UserHandler) GetUserProfile(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Writer.Header().Get("id"))

	profile, err := c.userUseCase.GetProfile(ctx, id)

	if err != nil {
		response := utils.ErrorResponse("Failed to Update User Email", err.Error(), nil)
		ctx.Writer.Header().Set("Content-Type", "application/json")
		ctx.Writer.WriteHeader(http.StatusUnprocessableEntity)
		utils.ResponseJSON(*ctx, response)
		return
	}
	response := utils.SuccessResponse(true, "SUCCESS", profile)
	ctx.Writer.Header().Set("Content-Type", "application/json")
	ctx.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*ctx, response)
}

// @Summary Add User Profile And Update Mail
// @ID AddProfileAndUpdateMail
// @Tags User Profile Management
// @Produce json
// @Security BearerAuth
// @Param userData body domain.UserData{} true "User Data"
// @Success 200 {object} utils.Response{}
// @Failure 422 {object} utils.Response{}
// @Router /user/profile [post]
func (c *UserHandler) AddProfileAndUpdateMail(ctx *gin.Context) {
	var userData domain.UserData
	id, _ := strconv.Atoi(ctx.Writer.Header().Get("id"))

	err := ctx.Bind(&userData)
	if err != nil {
		response := utils.ErrorResponse("Failed to Fetch Data", err.Error(), nil)
		ctx.Writer.Header().Set("Content-Type", "application/json")
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*ctx, response)
		return
	}
	userData.UserId = id

	err = c.userUseCase.UpdateMail(ctx, userData.Email, userData.UserId)

	if err != nil {
		response := utils.ErrorResponse("Failed to Update User Email", err.Error(), nil)
		ctx.Writer.Header().Set("Content-Type", "application/json")
		ctx.Writer.WriteHeader(http.StatusUnprocessableEntity)
		utils.ResponseJSON(*ctx, response)
		return
	}

	err = c.userUseCase.AddProfile(ctx, userData)

	if err != nil {
		response := utils.ErrorResponse("Failed to Add User Profile", err.Error(), nil)
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

func NewUserHandler(userUseCase services.UserUseCase) UserHandler {
	return UserHandler{
		userUseCase: userUseCase,
	}
}
