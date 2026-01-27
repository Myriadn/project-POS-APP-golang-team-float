package repository

import (
	"project-POS-APP-golang-team-float/internal/data/entity"
)

// FindUserByEmail finds a user by email address
func (r *Repository) FindUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := r.db.Preload("Role").Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindUserByID finds a user by their ID
func (r *Repository) FindUserByID(id uint) (*entity.User, error) {
	var user entity.User
	err := r.db.Preload("Role").First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateUser creates a new user in the database
func (r *Repository) CreateUser(user *entity.User) error {
	return r.db.Create(user).Error
}

// UpdateUser updates an existing user in the database
func (r *Repository) UpdateUser(user *entity.User) error {
	return r.db.Save(user).Error
}

// UpdatePassword updates the password hash for a user
func (r *Repository) UpdatePassword(userID uint, passwordHash string) error {
	return r.db.Model(&entity.User{}).Where("id = ?", userID).Update("password_hash", passwordHash).Error
}

// EmailExists checks if an email already exists in the database
func (r *Repository) EmailExists(email string) bool {
	var count int64
	r.db.Model(&entity.User{}).Where("email = ?", email).Count(&count)
	return count > 0
}
