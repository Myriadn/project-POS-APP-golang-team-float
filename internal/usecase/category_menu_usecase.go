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
