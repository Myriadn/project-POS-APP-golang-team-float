package repository

import (
	"context"
	"errors"
	"project-POS-APP-golang-team-float/internal/data/entity"
	"project-POS-APP-golang-team-float/internal/dto"

	"gorm.io/gorm"
)

type StaffManagementRepo struct {
	db *gorm.DB
}
type StaffManagementRepoInterface interface {
	CreateNewStaffManagement(ctx context.Context, user *entity.User) error
	UpdateStaffManagement(ctx context.Context, id uint, data map[string]interface{}) error
	GetDetailStaffManagement(ctx context.Context, id uint) (*entity.User, error)
	DeleteStaffManagement(ctx context.Context, id uint) error
	GetAllStaffManagement(ctx context.Context, f dto.GetStaffManagementFilterRequest) ([]*entity.User, int64, error)
}

func NewStaffManagementRepo(db *gorm.DB) StaffManagementRepoInterface {
	return &StaffManagementRepo{
		db: db,
	}
}

// membuat akun staff baru
func (b *StaffManagementRepo) CreateNewStaffManagement(ctx context.Context, user *entity.User) error {
	result := b.db.WithContext(ctx).Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// mengedit data staff
func (b *StaffManagementRepo) UpdateStaffManagement(ctx context.Context, id uint, data map[string]interface{}) error {
	result := b.db.WithContext(ctx).Model(&entity.User{}).Where("id=?", id).Updates(data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// mendapatkan detail informasi staff
func (b *StaffManagementRepo) GetDetailStaffManagement(ctx context.Context, id uint) (*entity.User, error) {
	var user entity.User
	result := b.db.WithContext(ctx).Preload("Role").First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// deete akun staff
func (b *StaffManagementRepo) DeleteStaffManagement(ctx context.Context, id uint) error {
	result := b.db.WithContext(ctx).Delete(&entity.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("data staff tidak ditemukan")
	}
	return nil
}

// mendapatkan daftar data staff
func (b *StaffManagementRepo) GetAllStaffManagement(ctx context.Context, f dto.GetStaffManagementFilterRequest) ([]*entity.User, int64, error) {
	var users []*entity.User
	var totalItems int64

	query := b.db.WithContext(ctx).Model(&entity.User{})

	if err := query.Count(&totalItems).Error; err != nil {
		return nil, 0, err
	}
	switch f.SortBy {
	case "full_name_asc":
		query = query.Order("full_name asc")
	case "full_name_desc":
		query = query.Order("full_name desc")
	case "email_asc":
		query = query.Order("email asc")
	case "email_desc":
		query = query.Order("email desc")
	default:
		query = query.Order("created_at desc")
	}

	offset := (f.Page - 1) * f.Limit

	result := query.Preload("Role").
		Limit(f.Limit).
		Offset(offset).
		Find(&users)

	if result.Error != nil {
		return nil, 0, result.Error
	}

	return users, totalItems, nil
}
