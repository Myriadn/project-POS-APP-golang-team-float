package usecase

import (
	"context"
	"project-POS-APP-golang-team-float/internal/data/entity"
)

type MockStaffRepo struct {
	CreateFn    func(ctx context.Context, user *entity.User) error
	UpdateFn    func(ctx context.Context, id uint, data map[string]interface{}) error
	GetDetailFn func(ctx context.Context, id uint) (*entity.User, error)
	DeleteFn    func(ctx context.Context, id uint) error
}

func (m *MockStaffRepo) CreateNewStaffManagement(ctx context.Context, user *entity.User) error {
	return m.CreateFn(ctx, user)
}
func (m *MockStaffRepo) UpdateStaffManagement(ctx context.Context, id uint, data map[string]interface{}) error {
	return m.UpdateFn(ctx, id, data)
}

func (m *MockStaffRepo) GetDetailStaffManagement(ctx context.Context, id uint) (*entity.User, error) {
	return m.GetDetailFn(ctx, id)
}

func (m *MockStaffRepo) DeleteStaffManagement(ctx context.Context, id uint) error {
	return m.DeleteFn(ctx, id)
}
