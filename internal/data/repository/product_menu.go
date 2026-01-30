package repository

import (
	"context"
	"errors"
	"project-POS-APP-golang-team-float/internal/data/entity"
	"project-POS-APP-golang-team-float/internal/dto"

	"gorm.io/gorm"
)

type ProductMenuRepo struct {
	db *gorm.DB
}
type ProductMenuRepoInterface interface {
	CreateNewProduct(ctx context.Context, product *entity.Product) error
	UpdateProductMenu(ctx context.Context, id uint, data map[string]interface{}) error
	GetDetailProductMenu(ctx context.Context, id uint) (*entity.Product, error)
	GetAllProductMenu(ctx context.Context, f dto.FilterRequest) ([]*entity.Product, int64, error)
	DeleteProductMenu(ctx context.Context, id uint) error
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

// mendapatkan detail product menu
func (b *ProductMenuRepo) GetDetailProductMenu(ctx context.Context, id uint) (*entity.Product, error) {
	var product entity.Product
	result := b.db.WithContext(ctx).Preload("Category").First(&product, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

// mendapatkan semua product
func (b *ProductMenuRepo) GetAllProductMenu(ctx context.Context, f dto.FilterRequest) ([]*entity.Product, int64, error) {
	var product []*entity.Product
	var totalItems int64

	query := b.db.WithContext(ctx).Model(&entity.Product{})

	if err := query.Count(&totalItems).Error; err != nil {
		return nil, 0, err
	}

	offset := (f.Page - 1) * f.Limit
	if f.MenuType != "" {
		query = query.Where("products.menu_type ILIKE ?", "%"+f.MenuType+"%")
	}
	if f.Category != "" {
		query = query.Where("products.category ILIKE ?", "%"+f.Category+"%")
	}
	if f.Status != "" {
		query = query.Where("products.status ILIKE ?", "%"+f.Status+"%")
	}
	if f.Stock != "" {
		query = query.Where("products.availability ILIKE ?", "%"+f.Stock+"%")
	}
	if f.Value != "" {
		query = query.Where("products.unit ILIKE ?", "%"+f.Value+"%")
	}
	if f.Piece != "" {
		query = query.Where("products.stock = ?", f.Piece)
	}
	if f.PriceMin != "" {
		query = query.Where("products.price >= ?", f.PriceMin)
	}
	if f.PriceMax != "" {
		query = query.Where("products.price <= ?", f.PriceMax)
	}

	result := query.Preload("Category").
		Limit(f.Limit).
		Offset(offset).
		Find(&product)

	if result.Error != nil {
		return nil, 0, result.Error
	}

	return product, totalItems, nil
}

// delete product menu
func (b *ProductMenuRepo) DeleteProductMenu(ctx context.Context, id uint) error {
	result := b.db.WithContext(ctx).Delete(&entity.Product{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("data product menu tidak ditemukan")
	}
	return nil
}
