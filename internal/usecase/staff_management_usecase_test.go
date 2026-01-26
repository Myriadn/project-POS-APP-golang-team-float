package usecase

import (
	"context"
	"project-POS-APP-golang-team-float/internal/data/entity"
	"project-POS-APP-golang-team-float/internal/dto"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewStaffManagementUsecase(t *testing.T) {

	req := dto.CreateNewStaffManagementReq{
		Email:    "cakra@gmail.com",
		Username: "Cakra",
		Password: "password123",
		RoleID:   3,
	}

	t.Run("success", func(t *testing.T) {
		mockRepo := &MockStaffRepo{
			CreateFn: func(ctx context.Context, user *entity.User) error {

				//asset logic
				assert.Equal(t, req.Email, user.Email)
				assert.NotEmpty(t, user.PasswordHash)
				assert.NotEqual(t, req.Password, user.PasswordHash)

				return nil
			},
		}

		uc := NewStaffManagementUsecase(mockRepo)

		res, err := uc.CreateNewStaffManagementUsecase(context.Background(), req)

		assert.NoError(t, err)
		assert.Equal(t, "berhasil membuat data staff atau admin", res.Message)
	})
	t.Run("failed_repo_error", func(t *testing.T) {
		mockRepo := &MockStaffRepo{
			CreateFn: func(ctx context.Context, user *entity.User) error {
				return assert.AnError //simulasi error
			},
		}

		uc := NewStaffManagementUsecase(mockRepo)

		res, err := uc.CreateNewStaffManagementUsecase(context.Background(), req)

		assert.Error(t, err)
		assert.Nil(t, res)
	})
}
