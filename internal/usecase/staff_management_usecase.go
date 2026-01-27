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
	UpdateStaffManagementUsecase(ctx context.Context, id uint, req dto.CreateNewStaffManagementReq) (*dto.MessageResponse, error)
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

func (b *StaffManagementUsecase) UpdateStaffManagementUsecase(ctx context.Context, id uint, req dto.CreateNewStaffManagementReq) (*dto.MessageResponse, error) {
	passwordHash, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}
	newStaff := map[string]interface{}{
		"email":              req.Email,
		"username":           req.Username,
		"password_hash":      passwordHash,
		"full_name":          req.FullName,
		"phone":              req.Phone,
		"role_id":            req.RoleID,
		"profile_picture":    req.ProfilePicture,
		"salary":             req.Salary,
		"date_of_birth":      req.DateOfBirth,
		"shift_start":        req.ShiftStart,
		"shift_end":          req.ShiftEnd,
		"address":            req.Address,
		"additional_details": req.AdditionalDetails,
	}

	// 2. Cek String: Hanya masukkan jika tidak kosong ("")
	if req.Email != "" {
		newStaff["email"] = req.Email
	}
	if req.Username != "" {
		newStaff["username"] = req.Username
	}
	if req.FullName != "" {
		newStaff["full_name"] = req.FullName
	}
	if req.Phone != "" {
		newStaff["phone"] = req.Phone
	}
	if req.ProfilePicture != "" {
		newStaff["profile_picture"] = req.ProfilePicture
	}
	if req.ShiftStart != "" {
		newStaff["shift_start"] = req.ShiftStart
	}
	if req.ShiftEnd != "" {
		newStaff["shift_end"] = req.ShiftEnd
	}
	if req.Address != "" {
		newStaff["address"] = req.Address
	}
	if req.AdditionalDetails != "" {
		newStaff["additional_details"] = req.AdditionalDetails
	}
	// Asumsi DateOfBirth adalah string (misal "1990-01-01")
	if !req.DateOfBirth.IsZero() {
		newStaff["date_of_birth"] = req.DateOfBirth
	}

	if req.RoleID != 0 {
		newStaff["role_id"] = req.RoleID
	}
	if req.Salary != 0 {
		newStaff["salary"] = req.Salary
	}

	if req.Password != "" {
		passwordHash, err := utils.HashPassword(req.Password)
		if err != nil {
			return nil, err
		}
		newStaff["password_hash"] = passwordHash
	}
	err = b.repo.UpdateStaffManagement(ctx, id, newStaff)
	if err != nil {
		return nil, err
	}
	return &dto.MessageResponse{Message: "berhasil update data staff atau admin"}, nil
}
