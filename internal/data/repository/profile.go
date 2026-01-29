package repository

import (
	"context"
	"project-POS-APP-golang-team-float/internal/data/entity"

	"gorm.io/gorm"
)

type ProfileRepo struct {
	db *gorm.DB
}
type ProfileRepoInterface interface {
	UpdateProfileUser(ctx context.Context, id uint, data map[string]interface{}) error
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
