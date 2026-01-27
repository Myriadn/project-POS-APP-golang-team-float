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
	UpdateStaffManagementUsecase(ctx context.Context, id uint, req dto.UpdateStaffManagementReq) (*dto.MessageResponse, error)
	GetDetailStaffManagement(ctx context.Context, id uint) (*dto.DetailStaffResponse, *dto.MessageResponse, error)
	DeleteStaffManagementUsecase(ctx context.Context, id uint) (*dto.MessageResponse, error)
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

// logic bisnis untuk update dagta staff
func (b *StaffManagementUsecase) UpdateStaffManagementUsecase(ctx context.Context, id uint, req dto.UpdateStaffManagementReq) (*dto.MessageResponse, error) {

	updateData := make(map[string]interface{})

	if req.Email != "" {
		updateData["email"] = req.Email
	}
	if req.Username != "" {
		updateData["username"] = req.Username
	}
	if req.FullName != "" {
		updateData["full_name"] = req.FullName
	}
	if req.Phone != "" {
		updateData["phone"] = req.Phone
	}
	if req.ProfilePicture != "" {
		updateData["profile_picture"] = req.ProfilePicture
	}
	if req.Address != "" {
		updateData["address"] = req.Address
	}
	if req.ShiftStart != "" {
		updateData["shift_start"] = req.ShiftStart
	}
	if req.ShiftEnd != "" {
		updateData["shift_end"] = req.ShiftEnd
	}
	if req.AdditionalDetails != "" {
		updateData["additional_details"] = req.AdditionalDetails
	}

	if req.DateOfBirth != nil {
		updateData["date_of_birth"] = req.DateOfBirth
	}

	if req.RoleID != 0 {
		updateData["role_id"] = req.RoleID
	}
	if req.Salary != 0 {
		updateData["salary"] = req.Salary
	}

	if req.Password != "" {
		hash, err := utils.HashPassword(req.Password)
		if err != nil {
			return nil, err
		}
		updateData["password_hash"] = hash
	}

	if len(updateData) == 0 {
		return &dto.MessageResponse{Message: "Tidak ada data yang perlu diubah"}, nil
	}

	err := b.repo.UpdateStaffManagement(ctx, uint(id), updateData)
	if err != nil {
		return nil, err
	}

	return &dto.MessageResponse{Message: "Berhasil update data staff"}, nil
}

func (b *StaffManagementUsecase) GetDetailStaffManagement(ctx context.Context, id uint) (*dto.DetailStaffResponse, *dto.MessageResponse, error) {
	userEntity, err := b.repo.GetDetailStaffManagement(ctx, id)
	if err != nil {
		return nil, nil, err
	}
	resp := &dto.DetailStaffResponse{
		ID:             userEntity.ID,
		Email:          userEntity.Email,
		FullName:       userEntity.FullName,
		Phone:          userEntity.Phone,
		RoleName:       userEntity.Role.Name,
		Salary:         userEntity.Salary,
		ShiftStart:     userEntity.ShiftStart,
		ShiftEnd:       userEntity.ShiftEnd,
		Address:        userEntity.Address,
		ProfilePicture: userEntity.ProfilePicture,
	}
	return resp, &dto.MessageResponse{Message: "berhasil mengambil detail staff"}, nil
}

func (b *StaffManagementUsecase) DeleteStaffManagementUsecase(ctx context.Context, id uint) (*dto.MessageResponse, error) {

	err := b.repo.DeleteStaffManagement(ctx, uint(id))
	if err != nil {
		return nil, err
	}

	return &dto.MessageResponse{Message: "Berhasil delete data staff"}, nil
}
