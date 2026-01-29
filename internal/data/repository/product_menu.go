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
