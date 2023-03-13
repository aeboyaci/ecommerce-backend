package main

import (
	"ecommerce-backend/pkg/bootstrap"
	"github.com/gin-gonic/gin"
)

func main() {
	bootstrap.Initialize()

	r := gin.Default()
	bootstrap.RegisterRouters(r)
	r.Run("0.0.0.0:8080")
}
