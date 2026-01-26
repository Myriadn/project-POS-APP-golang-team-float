package repository

import (
	"context"
	"project-POS-APP-golang-team-float/internal/data/entity"

	"gorm.io/gorm"
)

type StaffManagementRepo struct {
	db *gorm.DB
}
type StaffManagementRepoInterface interface {
	CreateNewStaffManagement(ctx context.Context, user *entity.User) error
	UpdateStaffManagement(ctx context.Context, id int, data map[string]interface{}) error
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
func (b *StaffManagementRepo) UpdateStaffManagement(ctx context.Context, id int, data map[string]interface{}) error {
	result := b.db.WithContext(ctx).Model(&entity.User{}).Where("id=?", id).Updates(data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
