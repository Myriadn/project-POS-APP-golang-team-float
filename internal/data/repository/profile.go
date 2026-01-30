package repository

import (
	"context"
	"project-POS-APP-golang-team-float/internal/data/entity"
	"project-POS-APP-golang-team-float/internal/dto"

	"gorm.io/gorm"
)

type ProfileRepo struct {
	db *gorm.DB
}
type ProfileRepoInterface interface {
	UpdateProfileUser(ctx context.Context, id uint, data map[string]interface{}) error
	GetAllAdminUser(ctx context.Context, f dto.FilterRequest) ([]*entity.User, int64, error)
	UpdateAccsessControl(ctx context.Context, userID uint, permissionID []uint) error
}

func NewProfileRepo(db *gorm.DB) ProfileRepoInterface {
	return &ProfileRepo{
		db: db,
	}
}

// mengedit profile dengan beberapa data saja
func (b *ProfileRepo) UpdateProfileUser(ctx context.Context, id uint, data map[string]interface{}) error {
	result := b.db.WithContext(ctx).Model(&entity.User{}).Where("id=?", id).Updates(data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (b *ProfileRepo) GetAllAdminUser(ctx context.Context, f dto.FilterRequest) ([]*entity.User, int64, error) {
	var user []*entity.User
	var totalItems int64

	query := b.db.WithContext(ctx).Model(&entity.User{})

	if err := query.Count(&totalItems).Error; err != nil {
		return nil, 0, err
	}

	offset := (f.Page - 1) * f.Limit

	result := query.Where("role_id", 2).Preload("Role").
		Limit(f.Limit).
		Offset(offset).
		Find(&user)

	if result.Error != nil {
		return nil, 0, result.Error
	}

	return user, totalItems, nil
}

// pemblokiran akses user
func (b *ProfileRepo) UpdateAccsessControl(ctx context.Context, userID uint, permissionID []uint) error {
	return b.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Unscoped().Where("user_id = ?", userID).Delete(&entity.UserPermission{}).Error; err != nil {
			return err
		}
		if len(permissionID) == 0 {
			return nil
		}
		var newPermissions []entity.UserPermission

		for _, pid := range permissionID {
			newPermissions = append(newPermissions, entity.UserPermission{
				UserID:       userID,
				PermissionID: pid,
			})
		}
		if err := tx.Create(&newPermissions).Error; err != nil {
			return err
		}

		return nil
	})
}
