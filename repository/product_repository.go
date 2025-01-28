package repository

import (
	"time"

	"github.com/AsrofunNiam/learn-grpc/helper"
	"github.com/AsrofunNiam/learn-grpc/model"
	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll(db *gorm.DB, filters *map[string]string) model.Products
	FindByID(db *gorm.DB, id uint) model.Product
	Create(db *gorm.DB, Product *model.Product) (*model.Product, error)
	Update(db *gorm.DB, Product *model.Product) *model.Product
	Delete(db *gorm.DB, id, deletedByID uint)
}

type ProductRepositoryImpl struct {
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (repository *ProductRepositoryImpl) FindAll(db *gorm.DB, filters *map[string]string) model.Products {
	products := model.Products{}
	currentDate := time.Now().Format("2006-01-02")
	tx := db.Model(&model.Product{})

	err := tx.Preload("ProductPrice", "start_date <= ? AND end_date >= ?", currentDate, currentDate).Preload("Company").Find(&products).Error
	helper.PanicIfError(err)

	return products
}

func (repository *ProductRepositoryImpl) FindByID(db *gorm.DB, id uint) model.Product {
	var product model.Product
	err := db.First(&product, id).Error
	helper.PanicIfError(err)
	return product
}

func (repository *ProductRepositoryImpl) Create(db *gorm.DB, product *model.Product) (*model.Product, error) {
	err := db.Create(&product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (repository *ProductRepositoryImpl) Update(db *gorm.DB, product *model.Product) *model.Product {
	err := db.Updates(&product).First(&product).Error
	helper.PanicIfError(err)

	return product
}

func (repository *ProductRepositoryImpl) Delete(db *gorm.DB, id, deletedByID uint) {
	err := db.First(&model.Product{}, id).Error
	helper.PanicIfError(err)

	// soft delete
	err = db.Updates(&model.Product{
		Model:       gorm.Model{ID: uint(id)},
		DeletedByID: deletedByID,
	}).Error

	helper.PanicIfError(err)
}
