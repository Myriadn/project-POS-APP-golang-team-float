package usecase

import (
	"context"
	"project-POS-APP-golang-team-float/internal/data/entity"
)

type MockStaffRepo struct {
	CreateFn func(ctx context.Context, user *entity.User) error
}

func (m *MockStaffRepo) CreateNewStaffManagement(
	ctx context.Context,
	user *entity.User,
) error {
	return m.CreateFn(ctx, user)
}
