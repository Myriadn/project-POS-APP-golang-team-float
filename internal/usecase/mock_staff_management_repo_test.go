package usecase

import (
	"context"
	"project-POS-APP-golang-team-float/internal/data/entity"
)

type MockStaffRepo struct {
	CreateFn func(ctx context.Context, user *entity.User) error
	UpdateFn func(ctx context.Context, id int, data map[string]interface{}) error
}

func (m *MockStaffRepo) CreateNewStaffManagement(
	ctx context.Context,
	user *entity.User,
) error {
	return m.CreateFn(ctx, user)
}
func (m *MockStaffRepo) UpdateStaffManagement(
	ctx context.Context, id int,
	data map[string]interface{},
) error {
	return m.UpdateFn(ctx, id, data)
}
