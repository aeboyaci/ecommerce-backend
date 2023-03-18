package address

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

func RegisterRouter(protectedApiRouter *gin.RouterGroup) {
	r := router{
		controller: newController(),
	}

	addressRouter := protectedApiRouter.Group("/addresses")
	addressRouter.POST("", r.addNewAddress)
}

func (r router) addNewAddress(ctx *gin.Context) {
	userId := ctx.GetString("userId")

	var address models.Address
	if err := ctx.BindJSON(&address); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "cannot parse request body. invalid request body sent",
		})
		return
	}

	err := r.controller.addNewAddress(userId, address)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	logger.GetInstance().Debug(fmt.Sprintf("address created for %s", userId))
	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    "address successfully created",
	})
}
