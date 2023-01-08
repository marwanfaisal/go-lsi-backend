package repository

import (
	"go-lsi-backend/domain/source_product/model"

	"gorm.io/gorm"
)

type (
	sourceProductRepository struct {
		DB *gorm.DB
	}

	SourceProductRepositoryInterface interface {
		BulkCreate(products []model.SourceProduct) (err error)
		DeleteAll() (err error)
		List() (products []model.SourceProduct, err error)
	}
)

func NewSourceProductRepository(DB *gorm.DB) SourceProductRepositoryInterface {
	return &sourceProductRepository{
		DB: DB,
	}
}

func (r *sourceProductRepository) BulkCreate(products []model.SourceProduct) (err error) {
	if err = r.DB.Create(&products).Error; err != nil {
		return err
	}
	return
}

func (r *sourceProductRepository) DeleteAll() (err error) {
	if err = r.DB.Exec("TRUNCATE source_products RESTART IDENTITY;").Error; err != nil {
		return err
	}
	return
}

func (r *sourceProductRepository) List() (products []model.SourceProduct, err error) {
	if err = r.DB.Order("id asc").Find(&products).Error; err != nil {
		return
	}
	return
}
