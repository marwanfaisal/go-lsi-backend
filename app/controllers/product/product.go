package product

import (
	"go-lsi-backend/db"
	"go-lsi-backend/domain/destination_product"
	destinationProductModel "go-lsi-backend/domain/destination_product/model"
	"go-lsi-backend/domain/source_product"
	sourceProductModel "go-lsi-backend/domain/source_product/model"
	"go-lsi-backend/lib/response"
	"log"
	"net/http"

	destinationProductRepo "go-lsi-backend/domain/destination_product/repository"
	sourceProductRepo "go-lsi-backend/domain/source_product/repository"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

func Sync(c *gin.Context) {
	var sourceProducts []sourceProductModel.SourceProduct
	var destinationProducts []destinationProductModel.DestinationProduct
	var err error

	go func() {
		sourceProductService := source_product.NewSourceProductService(sourceProductRepo.NewSourceProductRepository(db.SourceDB))
		destinationProductService := destination_product.NewDestinationProductService(destinationProductRepo.NewSourceProductRepository(db.DestinationDB))

		sourceProducts, err = sourceProductService.List()
		if err != nil {
			log.Fatal(err)
		}

		copier.Copy(&destinationProducts, &sourceProducts)
		destinationProductService.BulkUpdate(destinationProducts)
	}()

	response.Json(c, http.StatusOK, map[string]any{"message": "Success, process running on the background"})
	return
}

func Source(c *gin.Context) {
	sourceProductService := source_product.NewSourceProductService(sourceProductRepo.NewSourceProductRepository(db.SourceDB))
	sourceProducts, err := sourceProductService.List()
	if err != nil {
		response.Json(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Json(c, http.StatusOK, sourceProducts)
}

func Destination(c *gin.Context) {
	destinationProductService := destination_product.NewDestinationProductService(destinationProductRepo.NewSourceProductRepository(db.DestinationDB))
	destinationProducts, err := destinationProductService.List()
	if err != nil {
		response.Json(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Json(c, http.StatusOK, destinationProducts)
}
