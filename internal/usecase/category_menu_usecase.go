package usecase

import (
	"context"
	"math"
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
	GetAllCategoryMenu(ctx context.Context, req dto.FilterRequest) ([]*dto.AllCategoryMenuResponse, dto.Pagination, error)
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

func (b *CategoryMenuUsecase) GetAllCategoryMenu(ctx context.Context, req dto.FilterRequest) ([]*dto.AllCategoryMenuResponse, dto.Pagination, error) {
	if req.Page == 0 {
		req.Page = 1
	}
	if req.Limit == 0 {
		req.Limit = 6
	}

	category, total, err := b.repo.GetAllCategoryMenu(ctx, req)
	if err != nil {
		return nil, dto.Pagination{}, err
	}

	var categoryResponse []*dto.AllCategoryMenuResponse
	for _, t := range category {
		row := dto.AllCategoryMenuResponse{
			ID:         t.ID,
			Name:       t.Name,
			TotalItems: int64(len(t.Products)),
			Icon:       t.Icon,
		}
		categoryResponse = append(categoryResponse, &row)
	}
	totalPages := int(math.Ceil(float64(total) / float64(req.Limit)))

	pagination := dto.Pagination{
		CurrentPage:  req.Page,
		Limit:        req.Limit,
		TotalPages:   totalPages,
		TotalRecords: total,
	}
	return categoryResponse, pagination, nil
}
