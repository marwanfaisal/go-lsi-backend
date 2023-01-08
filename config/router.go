package config

import (
	"go-lsi-backend/app/controllers/product"

	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()
	router.Use(corsConfig)

	v1 := router.Group("v1")
	{
		products := v1.Group("product")
		{
			products.PUT("sync", product.Sync)
			products.GET("destination", product.Destination)
			products.GET("source", product.Source)
		}
	}

	router.Run(":8000")
}
