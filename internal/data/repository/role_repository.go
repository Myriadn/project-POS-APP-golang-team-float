package repository

import (
	"project-POS-APP-golang-team-float/internal/data/entity"
)

// FindRoleByName finds a role by its name
func (r *Repository) FindRoleByName(name string) (*entity.Role, error) {
	var role entity.Role
	err := r.db.Where("name = ?", name).First(&role).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

// FindRoleByID finds a role by its ID
func (r *Repository) FindRoleByID(id uint) (*entity.Role, error) {
	var role entity.Role
	err := r.db.First(&role, id).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

// GetAllRoles retrieves all roles from the database
func (r *Repository) GetAllRoles() ([]entity.Role, error) {
	var roles []entity.Role
	err := r.db.Find(&roles).Error
	if err != nil {
		return nil, err
	}
	return roles, nil
}

// CreateRole creates a new role in the database
func (r *Repository) CreateRole(role *entity.Role) error {
	return r.db.Create(role).Error
}

// UpdateRole updates an existing role in the database
func (r *Repository) UpdateRole(role *entity.Role) error {
	return r.db.Save(role).Error
}

// DeleteRole deletes a role by its ID
func (r *Repository) DeleteRole(id uint) error {
	return r.db.Delete(&entity.Role{}, id).Error
}
