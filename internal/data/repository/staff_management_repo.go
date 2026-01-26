package repository

import (
	"context"
	"project-POS-APP-golang-team-float/internal/data/entity"

	"gorm.io/gorm"
)

type StaffManagementRepo struct {
	db *gorm.DB
}

func NewStaffManagementRepo(db *gorm.DB) *StaffManagementRepo {
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
