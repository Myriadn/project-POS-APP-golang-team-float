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
	UpdateCategoryMenu(ctx context.Context, id uint, data map[string]interface{}) error
	GetDetailCategoryMenu(ctx context.Context, id uint) (*entity.Category, error)
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

// mengedit category menu dengan beberapa data saja
func (b *CategoryMenuRepo) UpdateCategoryMenu(ctx context.Context, id uint, data map[string]interface{}) error {
	result := b.db.WithContext(ctx).Model(&entity.Category{}).Where("id=?", id).Updates(data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// mendapatkan detail category menu
func (b *CategoryMenuRepo) GetDetailCategoryMenu(ctx context.Context, id uint) (*entity.Category, error) {
	var category entity.Category
	result := b.db.WithContext(ctx).First(&category, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &category, nil
}
