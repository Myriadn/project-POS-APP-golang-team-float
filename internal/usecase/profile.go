package usecase

import (
	"context"
	"math"
	"project-POS-APP-golang-team-float/internal/data/repository"
	"project-POS-APP-golang-team-float/internal/dto"
	"project-POS-APP-golang-team-float/pkg/utils"
)

type ProfileUsecase struct {
	repo repository.ProfileRepoInterface
}

type ProfileUsecaseInterface interface {
	UpdateProfileUsecase(ctx context.Context, id uint, req dto.UpdateProfileReq) (*dto.MessageResponse, error)
	GetAllAdminUser(ctx context.Context, req dto.FilterRequest) ([]*dto.GetlAllAdminResponse, dto.Pagination, error)
	UpdateAccsessControl(ctx context.Context, userID uint, req dto.AccsessReq) (*dto.MessageResponse, error)
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

// mendapatkan data admin
func (b *ProfileUsecase) GetAllAdminUser(ctx context.Context, req dto.FilterRequest) ([]*dto.GetlAllAdminResponse, dto.Pagination, error) {
	if req.Page == 0 {
		req.Page = 1
	}
	if req.Limit == 0 {
		req.Limit = 6
	}

	adminUser, total, err := b.repo.GetAllAdminUser(ctx, req)
	if err != nil {
		return nil, dto.Pagination{}, err
	}

	var adminUserResponse []*dto.GetlAllAdminResponse
	for _, t := range adminUser {
		row := dto.GetlAllAdminResponse{
			ID:       t.ID,
			Email:    t.Email,
			FullName: t.FullName,
			RoleName: t.Role.Name,
		}
		adminUserResponse = append(adminUserResponse, &row)
	}
	totalPages := int(math.Ceil(float64(total) / float64(req.Limit)))

	pagination := dto.Pagination{
		CurrentPage:  req.Page,
		Limit:        req.Limit,
		TotalPages:   totalPages,
		TotalRecords: total,
	}
	return adminUserResponse, pagination, nil
}

func (b *ProfileUsecase) UpdateAccsessControl(ctx context.Context, userID uint, req dto.AccsessReq) (*dto.MessageResponse, error) {

	err := b.repo.UpdateAccsessControl(ctx, userID, req.PermissionID)
	if err != nil {
		return nil, err
	}
	return &dto.MessageResponse{Message: "berhasil memblokir accsess admin"}, nil
}
