package main

import (
	"go-lsi-backend/db"
	"go-lsi-backend/domain/destination_product"
	destinationProductRepo "go-lsi-backend/domain/destination_product/repository"
	"go-lsi-backend/domain/source_product"
	sourceProductRepo "go-lsi-backend/domain/source_product/repository"
	"log"
)

func main() {
	db.Open()

	sourceProductService := source_product.NewSourceProductService(sourceProductRepo.NewSourceProductRepository(db.SourceDB))
	destinationProductService := destination_product.NewDestinationProductService(destinationProductRepo.NewSourceProductRepository(db.DestinationDB))

	if err := sourceProductService.DummyBulkCreate(500); err != nil {
		log.Fatal(err)
	}
	if err := destinationProductService.DummyBulkCreate(500); err != nil {
		log.Fatal(err)
	}
}
