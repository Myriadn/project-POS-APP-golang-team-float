package usecase

import (
	"context"
	"project-POS-APP-golang-team-float/internal/data/entity"
	"project-POS-APP-golang-team-float/internal/data/repository"
	"project-POS-APP-golang-team-float/internal/dto"
)

type CategoryMenuUsecase struct {
	repo repository.CategoryMenuRepoInterface
}

type CategoryMenuUsecaseInterface interface {
	CreateNewCategoryUsecase(ctx context.Context, req dto.CreateNewCategoryMenuReq) (*dto.MessageResponse, error)
	UpdateCategoryMenuUsecase(ctx context.Context, id uint, req dto.UpdateCategoryMenuReq) (*dto.MessageResponse, error)
	GetDetailCategoryMenu(ctx context.Context, id uint) (*dto.DetailCategoryResponse, *dto.MessageResponse, error)
}

func NewCategoryMenuUsecase(repo repository.CategoryMenuRepoInterface) CategoryMenuUsecaseInterface {
	return &CategoryMenuUsecase{
		repo: repo,
	}
}

func (b *CategoryMenuUsecase) CreateNewCategoryUsecase(ctx context.Context, req dto.CreateNewCategoryMenuReq) (*dto.MessageResponse, error) {
	newCategoryMenu := &entity.Category{
		Name:        req.Name,
		Description: req.Description,
		Icon:        req.Icon,
	}
	err := b.repo.CreateNewCategory(ctx, newCategoryMenu)
	if err != nil {
		return nil, err
	}
	return &dto.MessageResponse{Message: "berhasil membuat category menu baru"}, nil
}

// logic bisnis untuk update data category menu
func (b *CategoryMenuUsecase) UpdateCategoryMenuUsecase(ctx context.Context, id uint, req dto.UpdateCategoryMenuReq) (*dto.MessageResponse, error) {

	updateData := make(map[string]interface{})

	if req.Name != "" {
		updateData["name"] = req.Name
	}
	if req.Description != "" {
		updateData["description"] = req.Description
	}
	if req.Icon != "" {
		updateData["icon"] = req.Icon
	}

	if len(updateData) == 0 {
		return &dto.MessageResponse{Message: "Tidak ada data yang perlu diubah"}, nil
	}

	err := b.repo.UpdateCategoryMenu(ctx, uint(id), updateData)
	if err != nil {
		return nil, err
	}

	return &dto.MessageResponse{Message: "Berhasil update data category menu"}, nil
}

// ambil detail category menu
func (b *CategoryMenuUsecase) GetDetailCategoryMenu(ctx context.Context, id uint) (*dto.DetailCategoryResponse, *dto.MessageResponse, error) {
	category, err := b.repo.GetDetailCategoryMenu(ctx, id)
	if err != nil {
		return nil, nil, err
	}
	resp := &dto.DetailCategoryResponse{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		Icon:        category.Icon,
	}
	return resp, &dto.MessageResponse{Message: "berhasil mengambil detail category menu"}, nil
}
