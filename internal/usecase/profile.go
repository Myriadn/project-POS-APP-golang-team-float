package usecase

import (
	"context"
	"project-POS-APP-golang-team-float/internal/data/repository"
	"project-POS-APP-golang-team-float/internal/dto"
	"project-POS-APP-golang-team-float/pkg/utils"
)

type ProfileUsecase struct {
	repo repository.ProfileRepoInterface
}

type ProfileUsecaseInterface interface {
	UpdateProfileUsecase(ctx context.Context, id uint, req dto.UpdateProfileReq) (*dto.MessageResponse, error)
}

func NewProfileUsecase(repo repository.ProfileRepoInterface) ProfileUsecaseInterface {
	return &ProfileUsecase{
		repo: repo,
	}
}

// logic bisnis untuk update profile
func (b *ProfileUsecase) UpdateProfileUsecase(ctx context.Context, id uint, req dto.UpdateProfileReq) (*dto.MessageResponse, error) {

	updateData := make(map[string]interface{})

	if req.Username != "" {
		updateData["username"] = req.Username
	}
	if req.Email != "" {
		updateData["email"] = req.Email
	}
	if req.ProfilePicture != "" {
		updateData["profile_picture"] = req.ProfilePicture
	}
	if req.Address != "" {
		updateData["address"] = req.Address
	}

	if req.NewPassword != "" && req.ConfirmPassword != "" {
		if req.NewPassword == req.ConfirmPassword {
			password, _ := utils.HashPassword(req.NewPassword)
			updateData["password_hash"] = password
		}
	}
	if len(updateData) == 0 {
		return &dto.MessageResponse{Message: "Tidak ada data yang perlu diubah"}, nil
	}

	err := b.repo.UpdateProfileUser(ctx, uint(id), updateData)
	if err != nil {
		return nil, err
	}

	return &dto.MessageResponse{Message: "Berhasil update data profile"}, nil
}
