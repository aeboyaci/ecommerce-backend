package bootstrap

import (
	"ecommerce-backend/pkg/common/database"
	"ecommerce-backend/pkg/common/env"
	"ecommerce-backend/pkg/common/logger"
	"ecommerce-backend/pkg/handlers/account"
	"ecommerce-backend/pkg/middlewares"
	"github.com/gin-gonic/gin"
)

func Initialize() {
	log, err := logger.Initialize()
	if err != nil {
		panic(err)
	}

	err = env.Load()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Info("read environment variables")

	err = database.Initialize()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Info("connected to the database")
}

func RegisterRouters(r *gin.Engine) {
	apiRouter := r.Group("/api")
	protectedApiRouter := apiRouter.Group("")
	protectedApiRouter.Use(middlewares.EnforceAuthentication())

	account.RegisterRouter(apiRouter, protectedApiRouter)
}
