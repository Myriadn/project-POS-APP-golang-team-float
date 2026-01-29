package usecase

import (
	"context"
	"project-POS-APP-golang-team-float/internal/data/entity"
	"project-POS-APP-golang-team-float/internal/data/repository"
	"project-POS-APP-golang-team-float/internal/dto"
)

type ProductMenuUsecase struct {
	repo repository.ProductMenuRepoInterface
}

type ProductMenuUsecaseInterface interface {
	CreateNewProductUsecase(ctx context.Context, req dto.CreateNewProductMenuReq) (*dto.MessageResponse, error)
}

func NewProductMenuUsecase(repo repository.ProductMenuRepoInterface) ProductMenuUsecaseInterface {
	return &ProductMenuUsecase{
		repo: repo,
	}
}

func (b *ProductMenuUsecase) CreateNewProductUsecase(ctx context.Context, req dto.CreateNewProductMenuReq) (*dto.MessageResponse, error) {
	newProductMenu := &entity.Product{
		CategoryID:   req.CategotyID,
		Name:         req.Name,
		Description:  req.Description,
		Price:        req.Price,
		Stock:        req.Stock,
		Image:        req.Image,
		Availability: req.Availability,
		MenuType:     req.MenuType,
	}
	err := b.repo.CreateNewProduct(ctx, newProductMenu)
	if err != nil {
		return nil, err
	}
	return &dto.MessageResponse{Message: "berhasil membuat produk menu baru"}, nil
}
