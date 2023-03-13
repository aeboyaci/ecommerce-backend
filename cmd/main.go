package main

import (
	"ecommerce-backend/pkg/common/database"
	"ecommerce-backend/pkg/common/logger"
	"ecommerce-backend/pkg/handlers/authentication"
	"github.com/gin-gonic/gin"
)

func main() {
	log, err := logger.Initialize()
	if err != nil {
		panic(err)
	}

	err = database.Initialize()
	if err != nil {
		log.Error(err.Error())
	}

	r := gin.Default()

	apiRouter := r.Group("/api")
	authentication.RegisterRouter(apiRouter)

	r.Run("0.0.0.0:8080")
}
