package source_product

import (
	"fmt"
	"go-lsi-backend/domain/source_product/model"
	"go-lsi-backend/domain/source_product/repository"
	"math/rand"
	"time"
)

type (
	sourceProductService struct {
		repository repository.SourceProductRepositoryInterface
	}

	SourceProductServiceInterface interface {
		DummyBulkCreate(total int) (err error)
		List() (products []model.SourceProduct, err error)
	}
)

func NewSourceProductService(repository repository.SourceProductRepositoryInterface) SourceProductServiceInterface {
	return &sourceProductService{
		repository: repository,
	}
}

func (s *sourceProductService) DummyBulkCreate(total int) (err error) {
	var products []model.SourceProduct
	for i := 0; i < total; i++ {
		rand.Seed(time.Now().UnixNano())
		randNumber := rand.Intn(10-1) + 1
		products = append(products, model.SourceProduct{
			ProductName:  fmt.Sprintf("Product-%d", i+1),
			Qty:          randNumber,
			SellingPrice: float64(20000),
			PromoPrice:   float64(5000),
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

func (s *sourceProductService) List() (products []model.SourceProduct, err error) {
	if products, err = s.repository.List(); err != nil {
		return
	}
	return
}
