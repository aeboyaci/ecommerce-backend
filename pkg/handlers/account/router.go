package account

import (
	"ecommerce-backend/pkg/common/logger"
	"ecommerce-backend/pkg/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type router struct {
	controller controller
}

func RegisterRouter(apiRouter *gin.RouterGroup, protectedApiRouter *gin.RouterGroup) {
	r := router{
		controller: newController(),
	}

	authenticationRouter := apiRouter.Group("/account")
	authenticationRouter.POST("/sign-in", r.signIn)
	authenticationRouter.POST("/sign-up", r.signUp)

	protectedApiRouter.GET("/me", r.getUserInformation)
}

func (r router) signIn(ctx *gin.Context) {
	var user SignInDTO
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "cannot parse request body. invalid request body sent",
		})
		return
	}

	token, err := r.controller.signIn(user)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	ctx.SetCookie("token", token, 60*60*24, "/", "localhost", false, true)

	logger.GetInstance().Debug(fmt.Sprintf("user signed in with username: %s", user.Username))
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    "user successfully signed in",
	})
}

func (r router) signUp(ctx *gin.Context) {
	var user models.User
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "cannot parse request body. invalid request body sent",
		})
		return
	}

	err := r.controller.signUp(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	logger.GetInstance().Debug(fmt.Sprintf("user signed up with username: %s", user.Username))
	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    "user successfully created",
	})
}

func (r router) getUserInformation(ctx *gin.Context) {
	username := ctx.GetString("username")
	user, err := r.controller.getUserInformation(username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	logger.GetInstance().Debug(fmt.Sprintf("user information: %+v", user))
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    user,
	})
}
