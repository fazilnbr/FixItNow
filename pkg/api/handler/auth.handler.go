package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	_ "github.com/fazilnbr/project-workey/cmd/api/docs"
	"github.com/fazilnbr/project-workey/pkg/config"
	"github.com/fazilnbr/project-workey/pkg/domain"
	services "github.com/fazilnbr/project-workey/pkg/usecase/interface"
	"github.com/fazilnbr/project-workey/pkg/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
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

var (
	oauthConfGl = &oauth2.Config{
		ClientID:     "",
		ClientSecret: "",
		RedirectURL:  "http://localhost:8080/user/callback-gl",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
	oauthStateStringGl = ""
)

func (cr *AuthHandler) InitializeOAuthGoogle() {
	oauthConfGl.ClientID = cr.cfg.ClientID
	oauthConfGl.ClientSecret = cr.cfg.ClientSecret
	oauthStateStringGl = cr.cfg.OauthStateString
}

// @Summary Authenticate With Google
// @ID Authenticate With Google
// @Tags User Authentication
// @Security BearerAuth
// @Produce json
// @Success 200 {object} utils.Response{}
// @Failure 422 {object} utils.Response{}
// @Router /user/login-gl [get]
func (cr *AuthHandler) GoogleAuth(c *gin.Context) {
	HandileLogin(c, oauthConfGl, oauthStateStringGl)
}

func HandileLogin(c *gin.Context, oauthConf *oauth2.Config, oauthStateString string) error {
	URL, err := url.Parse(oauthConf.Endpoint.AuthURL)
	if err != nil {
		return err
	}
	parameters := url.Values{}
	parameters.Add("client_id", oauthConf.ClientID)
	parameters.Add("scope", strings.Join(oauthConf.Scopes, " "))
	parameters.Add("redirect_uri", oauthConf.RedirectURL)
	parameters.Add("response_type", "code")
	parameters.Add("state", oauthStateString)
	URL.RawQuery = parameters.Encode()
	url := URL.String()
	c.Redirect(http.StatusTemporaryRedirect, url)
	return nil

}

func (cr *AuthHandler) CallBackFromGoogle(ctx *gin.Context) {
	ctx.Request.ParseForm()
	state := ctx.Request.FormValue("state")

	if state != oauthStateStringGl {
		ctx.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	code := ctx.Request.FormValue("code")

	if code == "" {
		ctx.JSON(http.StatusBadRequest, "Code Not Found to provide AccessToken..\n")

		reason := ctx.Request.FormValue("error_reason")
		if reason == "user_denied" {
			ctx.JSON(http.StatusBadRequest, "User has denied Permission..")
		}
	} else {
		token, err := oauthConfGl.Exchange(oauth2.NoContext, code)
		if err != nil {
			return
		}
		resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + url.QueryEscape(token.AccessToken))
		if err != nil {
			ctx.Redirect(http.StatusTemporaryRedirect, "/")
			return
		}
		defer resp.Body.Close()

		respons, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			ctx.Redirect(http.StatusTemporaryRedirect, "/")
			return
		}
		type data struct {
			Id             string
			Email          string
			Verified_email bool
			Picture        string
			// data           string
		}
		var gdata data
		json.Unmarshal(respons, &gdata)

		if !gdata.Verified_email {
			response := utils.ErrorResponse("Failed to Login ", "Your email is not varified by google ", nil)
			ctx.Writer.Header().Add("Content-Type", "application/json")
			ctx.Writer.WriteHeader(http.StatusUnauthorized)
			utils.ResponseJSON(*ctx, response)
			return
		}

		userId, err := cr.userUseCase.RegisterAndVarifyWithEmail(ctx, gdata.Email)
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
}

// @Summary Send OTP for Users
// @ID sendOtp
// @Tags User Authentication
// @Produce json
// @Param WorkerLogin body domain.Signup{} true "Worker Login"
// @Success 200 {object} utils.Response{}
// @Failure 422 {object} utils.Response{}
// @Router /user/sent-otp [post]
func (cr *AuthHandler) UserSendOTP(ctx *gin.Context) {
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
	err = cr.authUseCase.SendOTP(ctx, phoneNumber)

	if err != nil {
		response := utils.ErrorResponse("Error while sending OTP to user", err.Error(), nil)
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
// @Param WorkerLogin body domain.Signup{} true "Worker Login"
// @Success 200 {object} utils.Response{}
// @Failure 422 {object} utils.Response{}
// @Router /user/signup-and-login [post]
func (cr *AuthHandler) UserRegisterAndLogin(ctx *gin.Context) {
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
	err = cr.authUseCase.VarifyOTP(ctx, phoneNumber, newUser.Otp)
	if err != nil {
		response := utils.ErrorResponse("Invalid OTP", err.Error(), nil)
		ctx.Writer.Header().Set("Content-Type", "application/json")
		ctx.Writer.WriteHeader(http.StatusUnprocessableEntity)
		utils.ResponseJSON(*ctx, response)
		return
	}
	userId, err := cr.userUseCase.RegisterAndVarifyWithNumber(ctx, phoneNumber)
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
