package repository

import (
	"context"
	"project-POS-APP-golang-team-float/internal/data/entity"

	"gorm.io/gorm"
)

type CategoryMenuRepo struct {
	db *gorm.DB
}
type CategoryMenuRepoInterface interface {
	CreateNewCategory(ctx context.Context, category *entity.Category) error
}

func NewCategoryMenuRepo(db *gorm.DB) CategoryMenuRepoInterface {
	return &CategoryMenuRepo{
		db: db,
	}
}

// membuat category menu baru
func (b *CategoryMenuRepo) CreateNewCategory(ctx context.Context, category *entity.Category) error {
	result := b.db.WithContext(ctx).Create(category)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
