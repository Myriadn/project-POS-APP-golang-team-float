package usecase

import (
	"context"
	"project-POS-APP-golang-team-float/internal/data/entity"
	"project-POS-APP-golang-team-float/internal/data/repository"
	"project-POS-APP-golang-team-float/internal/dto"
	"project-POS-APP-golang-team-float/pkg/utils"
)

type StaffManagementUsecase struct {
	repo repository.StaffManagementRepoInterface
}

type StaffManagementUsecaseInterface interface {
	CreateNewStaffManagementUsecase(ctx context.Context, req dto.CreateNewStaffManagementReq) (*dto.MessageResponse, error)
}

func NewStaffManagementUsecase(repo repository.StaffManagementRepoInterface) StaffManagementUsecaseInterface {
	return &StaffManagementUsecase{
		repo: repo,
	}
}

func (b *StaffManagementUsecase) CreateNewStaffManagementUsecase(ctx context.Context, req dto.CreateNewStaffManagementReq) (*dto.MessageResponse, error) {
	passwordHash, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}
	newStaff := &entity.User{
		Email:             req.Email,
		Username:          req.Username,
		PasswordHash:      passwordHash,
		FullName:          req.FullName,
		Phone:             req.Phone,
		RoleID:            req.RoleID,
		ProfilePicture:    req.ProfilePicture,
		Salary:            req.Salary,
		DateOfBirth:       req.DateOfBirth,
		ShiftStart:        req.ShiftStart,
		ShiftEnd:          req.ShiftEnd,
		Address:           req.Address,
		AdditionalDetails: req.AdditionalDetails,
	}
	err = b.repo.CreateNewStaffManagement(ctx, newStaff)
	if err != nil {
		return nil, err
	}
	return &dto.MessageResponse{Message: "berhasil membuat data staff atau admin"}, nil
}
