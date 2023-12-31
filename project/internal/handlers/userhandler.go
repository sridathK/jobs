package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"project/internal/auth"
	"project/internal/middlewear"
	"project/internal/model"
	"project/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

// type handler struct {
// 	a  *auth.Auth
// 	us services.UsersService
// 	cs services.CompanyService
// }

// func NewHandler(a *auth.Auth, us services.UsersService, cs services.CompanyService) (*handler, error) {
// 	if us == nil {
// 		return nil, errors.New("service implementation not given")
// 	}

// 	return &handler{a: a, us: us, cs: cs}, nil

// }

type UserHandler struct {
	a  *auth.Auth
	us services.UsersService
}

func NewUserHandler(a *auth.Auth, us services.UsersService) (UserHandlerInt, error) {
	if us == nil {
		return nil, errors.New("service implementation not given")
	}

	return &UserHandler{a: a, us: us}, nil

}

type UserHandlerInt interface {
	userSignup(c *gin.Context)
	userLogin(c *gin.Context)
	userForgetPassword(c *gin.Context)
	userUpdatePassword(c *gin.Context)
}

func (h *UserHandler) userSignup(c *gin.Context) {
	ctx := c.Request.Context()

	traceId, ok := ctx.Value(middlewear.TraceIdKey).(string)
	if !ok {
		log.Error().Str("traceId", traceId).Msg("trace id not found in userSignin handler")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	var userCreate model.UserSignup
	body := c.Request.Body
	err := json.NewDecoder(body).Decode(&userCreate)
	if err != nil {
		log.Error().Err(err).Msg("error in decoding")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}

	validate := validator.New()
	err = validate.Struct(&userCreate)
	if err != nil {
		log.Error().Err(err).Msg("error in validating ")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "invalid input"})
		return
	}
	us, err := h.us.UserSignup(userCreate)
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId).Msg("user signup problem")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "user signup failed"})
		return
	}
	c.JSON(http.StatusOK, us)

}

func (h *UserHandler) userLogin(c *gin.Context) {
	ctx := c.Request.Context()

	traceId, ok := ctx.Value(middlewear.TraceIdKey).(string)
	if !ok {
		log.Error().Str("traceId", traceId).Msg("trace id not found in userSignin handler")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}
	var userLogin model.UserLogin
	body := c.Request.Body
	err := json.NewDecoder(body).Decode(&userLogin)
	if err != nil {
		log.Error().Err(err).Msg("error in decoding")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	validate := validator.New()
	err = validate.Struct(&userLogin)
	if err != nil {
		log.Error().Err(err).Msg("error in validating ")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "invalid input"})
		return
	}

	regClaims, err := h.us.Userlogin(userLogin)
	if err != nil {
		log.Error().Err(err).Msg("error in Loginin ")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": "invalid input"})
		return
	}

	token, err := h.a.GenerateToken(regClaims)
	if err != nil {
		log.Error().Err(err).Msg("error in Gneerating toek ")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return

	}

	c.JSON(http.StatusOK, token)

}

func (h *UserHandler) userForgetPassword(c *gin.Context) {
	ctx := c.Request.Context()

	traceId, ok := ctx.Value(middlewear.TraceIdKey).(string)
	if !ok {
		log.Error().Str("traceId", traceId).Msg("trace id not found in userSignin handler")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	var userForgetPassword model.UserForgetPassword
	body := c.Request.Body
	err := json.NewDecoder(body).Decode(&userForgetPassword)
	if err != nil {
		log.Error().Err(err).Msg("error in decoding")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	validate := validator.New()
	err = validate.Struct(&userForgetPassword)
	if err != nil {
		log.Error().Err(err).Msg("error in validating ")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "invalid input"})
		return
	}
	string, err := h.us.UserForgetPassword(userForgetPassword)
	if err != nil {
		if string == "wrong creds given" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
			return
		}
		if string == "" {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
			return
		}
	}
	c.JSON(http.StatusOK, string)

}

func (h *UserHandler) userUpdatePassword(c *gin.Context) {
	ctx := c.Request.Context()

	traceId, ok := ctx.Value(middlewear.TraceIdKey).(string)
	if !ok {
		log.Error().Str("traceId", traceId).Msg("trace id not found in userSignin handler")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	var userUpdatePassword model.UserUpdatePassword
	body := c.Request.Body
	err := json.NewDecoder(body).Decode(&userUpdatePassword)
	if err != nil {
		log.Error().Err(err).Msg("error in decoding")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	validate := validator.New()
	err = validate.Struct(&userUpdatePassword)
	if err != nil {
		log.Error().Err(err).Msg("error in validating ")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "invalid input"})
		return
	}
	string, err := h.us.UserUpdatePassword(userUpdatePassword)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	c.JSON(http.StatusOK, string)

}
