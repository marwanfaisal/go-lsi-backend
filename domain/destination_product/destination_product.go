package destination_product

import (
	"fmt"
	"go-lsi-backend/domain/destination_product/model"
	"go-lsi-backend/domain/destination_product/repository"
	"math/rand"
	"time"
)

type (
	destinationProductService struct {
		repository repository.DestinationProductRepositoryInterface
	}

	DestinationProductServiceInterface interface {
		DummyBulkCreate(total int) (err error)
		BulkUpdate(products []model.DestinationProduct) (err error)
		List() (products []model.DestinationProduct, err error)
	}
)

func NewDestinationProductService(repository repository.DestinationProductRepositoryInterface) DestinationProductServiceInterface {
	return &destinationProductService{
		repository: repository,
	}
}

func (s *destinationProductService) DummyBulkCreate(total int) (err error) {
	var products []model.DestinationProduct
	for i := 0; i < total; i++ {
		rand.Seed(time.Now().UnixNano())
		products = append(products, model.DestinationProduct{
			ProductName:  fmt.Sprintf("Product-%d", i+1),
			Qty:          0,
			SellingPrice: 0,
			PromoPrice:   0,
		})
	}

	if err = s.repository.DeleteAll(); err != nil {
		return err
	}

	if err = s.repository.BulkCreate(products); err != nil {
		return err
	}

	return
}

func (s *destinationProductService) BulkUpdate(products []model.DestinationProduct) (err error) {
	if err = s.repository.BulkUpdate(products); err != nil {
		return err
	}

	return
}

func (s *destinationProductService) List() (products []model.DestinationProduct, err error) {
	if products, err = s.repository.List(); err != nil {
		return
	}
	return
}
