package usecase

import (
	"context"
	"math"
	"project-POS-APP-golang-team-float/internal/data/entity"
	"project-POS-APP-golang-team-float/internal/data/repository"
	"project-POS-APP-golang-team-float/internal/dto"
)

type ProductMenuUsecase struct {
	repo repository.ProductMenuRepoInterface
}

type ProductMenuUsecaseInterface interface {
	CreateNewProductUsecase(ctx context.Context, req dto.CreateNewProductMenuReq) (*dto.MessageResponse, error)
	UpdateProductMenuUsecase(ctx context.Context, id uint, req dto.UpdateProductMenuReq) (*dto.MessageResponse, error)
	GetDetailProductMenu(ctx context.Context, id uint) (*dto.DetailProductResponse, *dto.MessageResponse, error)
	GetAllProductMenu(ctx context.Context, req dto.FilterRequest) ([]*dto.AllProductResponse, dto.Pagination, error)
	DeleteProductMenu(ctx context.Context, id uint) (*dto.MessageResponse, error)
}

func NewProductMenuUsecase(repo repository.ProductMenuRepoInterface) ProductMenuUsecaseInterface {
	return &ProductMenuUsecase{
		repo: repo,
	}
}

// logic membuat product baru
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
		Unit:         req.Unit,
		Status:       req.Status,
	}
	err := b.repo.CreateNewProduct(ctx, newProductMenu)
	if err != nil {
		return nil, err
	}
	return &dto.MessageResponse{Message: "berhasil membuat produk menu baru"}, nil
}

// logic bisnis untuk update data product menu
func (b *ProductMenuUsecase) UpdateProductMenuUsecase(ctx context.Context, id uint, req dto.UpdateProductMenuReq) (*dto.MessageResponse, error) {

	updateData := make(map[string]interface{})

	if req.Name != "" {
		updateData["name"] = req.Name
	}
	if req.Description != "" {
		updateData["description"] = req.Description
	}
	if req.CategotyID != 0 {
		updateData["category_id"] = req.CategotyID
	}
	if req.Stock != 0 {
		updateData["stock"] = req.Stock
	}
	if req.Image != "" {
		updateData["image"] = req.Image
	}
	if req.Price != 0 {
		updateData["price"] = req.Price
	}
	if req.Availability != "" {
		updateData["availability"] = req.Availability
	}
	if req.MenuType != "" {
		updateData["menu_type"] = req.MenuType
	}
	if req.Unit != "" {
		updateData["unit"] = req.Unit
	}
	if req.Status != "" {
		updateData["status"] = req.Status
	}

	if len(updateData) == 0 {
		return &dto.MessageResponse{Message: "Tidak ada data yang perlu diubah"}, nil
	}

	err := b.repo.UpdateProductMenu(ctx, uint(id), updateData)
	if err != nil {
		return nil, err
	}

	return &dto.MessageResponse{Message: "Berhasil update data product menu"}, nil
}

// ambil detail product menu
func (b *ProductMenuUsecase) GetDetailProductMenu(ctx context.Context, id uint) (*dto.DetailProductResponse, *dto.MessageResponse, error) {
	product, err := b.repo.GetDetailProductMenu(ctx, id)
	if err != nil {
		return nil, nil, err
	}
	resp := &dto.DetailProductResponse{
		ID:           product.ID,
		Name:         product.Name,
		Description:  product.Description,
		Price:        product.Price,
		Stock:        product.Stock,
		CategotyName: product.Category.Name,
		Image:        product.Image,
		Availability: product.Availability,
	}
	return resp, &dto.MessageResponse{Message: "berhasil mengambil detail product menu"}, nil
}

// logic bisnis untuk mengambil semua product menu
func (b *ProductMenuUsecase) GetAllProductMenu(ctx context.Context, req dto.FilterRequest) ([]*dto.AllProductResponse, dto.Pagination, error) {
	if req.Page == 0 {
		req.Page = 1
	}
	if req.Limit == 0 {
		req.Limit = 6
	}

	product, total, err := b.repo.GetAllProductMenu(ctx, req)
	if err != nil {
		return nil, dto.Pagination{}, err
	}

	var productResponse []*dto.AllProductResponse
	for _, t := range product {
		row := dto.AllProductResponse{
			ID:           t.ID,
			Name:         t.Name,
			Price:        t.Price,
			Stock:        t.Stock,
			CategotyName: t.Category.Name,
			Image:        t.Image,
			Availability: t.Availability,
		}
		productResponse = append(productResponse, &row)
	}
	totalPages := int(math.Ceil(float64(total) / float64(req.Limit)))

	pagination := dto.Pagination{
		CurrentPage:  req.Page,
		Limit:        req.Limit,
		TotalPages:   totalPages,
		TotalRecords: total,
	}
	return productResponse, pagination, nil
}

// delete product menu
func (b *ProductMenuUsecase) DeleteProductMenu(ctx context.Context, id uint) (*dto.MessageResponse, error) {

	err := b.repo.DeleteProductMenu(ctx, uint(id))
	if err != nil {
		return nil, err
	}

	return &dto.MessageResponse{Message: "Berhasil delete data product menu"}, nil
}
