package repository

import (
	"context"
	"database/sql"
	"project-POS-APP-golang-team-float/internal/data/entity"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type suiteStaffManagement struct {
	suite.Suite
	mock   sqlmock.Sqlmock
	repo   StaffManagementRepoInterface
	sqlDB  *sql.DB
	gormDB *gorm.DB
}

// mock db
func MockDBStaffManagement(t *testing.T) (*gorm.DB, sqlmock.Sqlmock, *sql.DB) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	mockDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn:       db,
		DriverName: "postgres",
	}), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	assert.NoError(t, err)

	return mockDB, mock, db
}
func (s *suiteStaffManagement) SetupTest() {
	s.gormDB, s.mock, s.sqlDB = MockDBStaffManagement(s.T())

	s.repo = NewStaffManagementRepo(s.gormDB)
}

// tutup koneksi
func (s *suiteStaffManagement) TearDownSuite() {
	s.sqlDB.Close()
}

func (s *suiteStaffManagement) TestCreateNewStaffManagement() {
	id := 1

	staffDummy := &entity.User{
		Email:             "cakra@gmail.com",
		Username:          "Cakra",
		PasswordHash:      "halo123",
		FullName:          "Cakra Candra",
		Phone:             "123456781234",
		RoleID:            3,
		ProfilePicture:    "profile picture",
		Salary:            0,
		DateOfBirth:       nil,
		ShiftStart:        "09.00",
		ShiftEnd:          "16.00",
		Address:           "jalan buah",
		AdditionalDetails: "additional details",
		IsActive:          true,
	}
	//test ketika succsess
	s.Run("success", func() {
		staff := *staffDummy

		s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "users"`)).
			WithArgs(
				staff.Email,
				staff.Username,
				staff.PasswordHash,
				staff.FullName,
				staff.Phone,
				staff.RoleID,
				staff.ProfilePicture,
				staff.Salary,
				staff.DateOfBirth,
				staff.ShiftStart,
				staff.ShiftEnd,
				staff.Address,
				staff.AdditionalDetails,
				staff.IsActive,
				sqlmock.AnyArg(), // created_at
				sqlmock.AnyArg(), // updated_at
				sqlmock.AnyArg(), // deleted_at
			).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(id))

		err := s.repo.CreateNewStaffManagement(context.Background(), &staff)

		s.NoError(err)
		s.Equal(uint(id), staff.ID)
		s.NoError(s.mock.ExpectationsWereMet())
	})
	//test ketika gagal
	s.Run("failed", func() {
		staff := *staffDummy

		s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "users"`)).
			WillReturnError(gorm.ErrRecordNotFound)

		err := s.repo.CreateNewStaffManagement(context.Background(), &staff)

		s.Error(err)
		s.Equal(uint(0), staff.ID)
		s.NoError(s.mock.ExpectationsWereMet())
	})
}

func (s *suiteStaffManagement) TestUpdateStaffManagement() {
	id := uint(1)

	staffDummy := map[string]interface{}{
		"email":              "cakra@gmail.com",
		"username":           "Cakra",
		"password_hash":      "halo123",
		"full_name":          "Cakra Candra",
		"phone":              "123456781234",
		"role_id":            3,
		"profile_picture":    "profile picture",
		"salary":             0,
		"date_of_birth":      nil,
		"shift_start":        "09.00",
		"shift_end":          "16.00",
		"address":            "jalan buah",
		"additional_details": "additional details",
		"is_active":          true,
	}
	//test ketika succsess
	s.Run("success", func() {

		s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "users"`)).
			WillReturnResult(sqlmock.NewResult(0, 1))

		err := s.repo.UpdateStaffManagement(context.Background(), id, staffDummy)

		s.NoError(err)
		s.NoError(s.mock.ExpectationsWereMet())
	})
	//test ketika gagal
	s.Run("failed", func() {
		id := uint(2)

		s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "users"`)).
			WillReturnError(gorm.ErrRecordNotFound)

		err := s.repo.UpdateStaffManagement(context.Background(), id, staffDummy)

		s.Error(err)
		s.NoError(s.mock.ExpectationsWereMet())
	})
}

func (s *suiteStaffManagement) TestGetDetailStaffManagement() {

	s.Run("success", func() {
		id := 1
		now := time.Now()
		shiftStart := time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC) // Jam 09.00
		shiftEnd := time.Date(2024, 1, 1, 16, 0, 0, 0, time.UTC)  // Jam 16.00
		DateOfBirth := time.Date(2024, 1, 1, 16, 0, 0, 0, time.UTC)
		// Data Dummy
		rows := sqlmock.NewRows([]string{
			"id", "email", "username", "password_hash", "full_name", "phone", "role_id", "profile_picture", "salary", "date_of_birth", "shift_start", "shift_end",
			"address", "additional_details", "is_active", "created_at", "updated_at", "deleted_at",
		}).AddRow(id, "cakra@gmail.com", "Cakra", "halo123", "Cakra Candra", "123456781234", 3, "profile picture", 0, DateOfBirth, shiftStart, shiftEnd, "jalan buah", "additional details", true, now, now, nil)

		//menggunakan regexp.QuoteMeta untuk membaca *
		s.mock.ExpectQuery(`FROM "users"`).WithArgs(id).
			WillReturnRows(rows)

		user, err := s.repo.GetDetailStaffManagement(context.Background(), uint(id))

		//lakukan validasi
		s.NoError(err)
		s.NotNil(user)
		s.Equal(uint(id), user.ID)
		s.Equal("cakra@gmail.com", user.Email)
		s.Equal("Cakra", user.Username)

		// Cek apakah semua urutan mock terpenuhi
		s.NoError(s.mock.ExpectationsWereMet())
	})

	s.Run("not_found", func() {
		id := 2
		s.mock.ExpectQuery(`FROM "users"`).WithArgs(id, sqlmock.AnyArg()).
			WillReturnError(gorm.ErrRecordNotFound)

		user, err := s.repo.GetDetailStaffManagement(context.Background(), uint(id))

		s.Error(err)
		s.Nil(user)

		s.NoError(s.mock.ExpectationsWereMet())
	})
}

// fungsi untuk test semua fungsi testing yang di kumpulkan di suite
func TestNewStaffManagement(t *testing.T) {
	suite.Run(t, new(suiteStaffManagement))
}
