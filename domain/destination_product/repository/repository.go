package repository

import (
	"go-lsi-backend/domain/destination_product/model"

	"gorm.io/gorm"
)

type (
	destinationProductRepository struct {
		DB *gorm.DB
	}

	DestinationProductRepositoryInterface interface {
		BulkCreate(products []model.DestinationProduct) (err error)
		DeleteAll() (err error)
		BulkUpdate(products []model.DestinationProduct) (err error)
		List() (products []model.DestinationProduct, err error)
	}
)

func NewSourceProductRepository(DB *gorm.DB) DestinationProductRepositoryInterface {
	return &destinationProductRepository{
		DB: DB,
	}
}

func (r *destinationProductRepository) BulkCreate(products []model.DestinationProduct) (err error) {
	if err = r.DB.Create(&products).Error; err != nil {
		return err
	}
	return
}

func (r *destinationProductRepository) DeleteAll() (err error) {
	if err = r.DB.Exec("TRUNCATE destination_products RESTART IDENTITY;").Error; err != nil {
		return err
	}

	return
}

func (r *destinationProductRepository) BulkUpdate(products []model.DestinationProduct) (err error) {
	tx := r.DB.Begin()
	for _, product := range products {
		if err = tx.Select("Qty", "SellingPrice", "PromoPrice").Save(&product).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error

}

func (r *destinationProductRepository) List() (products []model.DestinationProduct, err error) {
	if err = r.DB.Order("id asc").Find(&products).Error; err != nil {
		return
	}
	return
}
