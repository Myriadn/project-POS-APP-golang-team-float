package repository

import (
	"context"
	"project-POS-APP-golang-team-float/internal/data/entity"

	"gorm.io/gorm"
)

type ProductMenuRepo struct {
	db *gorm.DB
}
type ProductMenuRepoInterface interface {
	CreateNewProduct(ctx context.Context, product *entity.Product) error
	UpdateProductMenu(ctx context.Context, id uint, data map[string]interface{}) error
}

func NewProductMenuRepo(db *gorm.DB) ProductMenuRepoInterface {
	return &ProductMenuRepo{
		db: db,
	}
}

// membuat product menu baru
func (b *ProductMenuRepo) CreateNewProduct(ctx context.Context, product *entity.Product) error {
	result := b.db.WithContext(ctx).Create(product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// mengedit product menu dengan beberapa data saja
func (b *ProductMenuRepo) UpdateProductMenu(ctx context.Context, id uint, data map[string]interface{}) error {
	result := b.db.WithContext(ctx).Model(&entity.Product{}).Where("id=?", id).Updates(data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
