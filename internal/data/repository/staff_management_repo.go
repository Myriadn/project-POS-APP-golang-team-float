package repository

import (
	"context"
	"errors"
	"project-POS-APP-golang-team-float/internal/data/entity"

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
}

func NewStaffManagementRepo(db *gorm.DB) StaffManagementRepoInterface {
	return &StaffManagementRepo{
		db: db,
	}
}

func (b *StaffManagementRepo) CreateNewStaffManagement(ctx context.Context, user *entity.User) error {
	result := b.db.WithContext(ctx).Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (b *StaffManagementRepo) UpdateStaffManagement(ctx context.Context, id uint, data map[string]interface{}) error {
	result := b.db.WithContext(ctx).Model(&entity.User{}).Where("id=?", id).Updates(data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (b *StaffManagementRepo) GetDetailStaffManagement(ctx context.Context, id uint) (*entity.User, error) {
	var user entity.User
	result := b.db.WithContext(ctx).Preload("Role").First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

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
