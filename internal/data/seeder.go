package data

import (
	"gorm.io/gorm"

	"project-POS-APP-golang-team-float/internal/data/entity"
	"project-POS-APP-golang-team-float/pkg/utils"
)

func Seed(db *gorm.DB) error {
	if err := seedRoles(db); err != nil {
		return err
	}
	if err := seedCategories(db); err != nil {
		return err
	}
	if err := seedPaymentMethods(db); err != nil {
		return err
	}
	if err := seedTables(db); err != nil {
		return err
	}
	if err := seedSuperAdmin(db); err != nil {
		return err
	}
	if err := seedPermission(db); err != nil {
		return err
	}
	if err := seedRolePermissions(db); err != nil {
		return err
	}
	return nil
}

func seedRoles(db *gorm.DB) error {
	roles := []entity.Role{
		{Name: "superadmin", Description: "Super Administrator dengan akses penuh"},
		{Name: "admin", Description: "Administrator dengan akses terbatas"},
		{Name: "staff", Description: "Staff dengan akses operasional dasar"},
	}
	for _, role := range roles {
		var existing entity.Role
		if db.Where("name = ?", role.Name).First(&existing).RowsAffected == 0 {
			if err := db.Create(&role).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

func seedCategories(db *gorm.DB) error {
	categories := []entity.Category{
		{Name: "All", Description: "Semua kategori menu"},
		{Name: "Pizza", Description: "Berbagai macam pizza", Icon: "/icons/pizza.png"},
		{Name: "Burger", Description: "Burger dengan berbagai pilihan", Icon: "/icons/burger.png"},
		{Name: "Chicken", Description: "Ayam goreng dan panggang", Icon: "/icons/chicken.png"},
		{Name: "Bakery", Description: "Roti dan kue", Icon: "/icons/bakery.png"},
		{Name: "Beverage", Description: "Minuman dingin dan panas", Icon: "/icons/beverage.png"},
		{Name: "Seafood", Description: "Hidangan laut", Icon: "/icons/seafood.png"},
	}
	for _, cat := range categories {
		var existing entity.Category
		if db.Where("name = ?", cat.Name).First(&existing).RowsAffected == 0 {
			if err := db.Create(&cat).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

func seedPaymentMethods(db *gorm.DB) error {
	methods := []entity.PaymentMethod{
		{Name: "Cash", IsActive: true},
		{Name: "Visa Card", IsActive: true},
		{Name: "Master Card", IsActive: true},
		{Name: "Debit Card", IsActive: true},
	}
	for _, method := range methods {
		var existing entity.PaymentMethod
		if db.Where("name = ?", method.Name).First(&existing).RowsAffected == 0 {
			if err := db.Create(&method).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

func seedTables(db *gorm.DB) error {
	tables := []entity.Table{
		{TableNumber: "01", Floor: 1, Capacity: 4, Status: "available"},
		{TableNumber: "02", Floor: 1, Capacity: 4, Status: "available"},
		{TableNumber: "03", Floor: 1, Capacity: 4, Status: "available"},
		{TableNumber: "04", Floor: 1, Capacity: 4, Status: "available"},
		{TableNumber: "05", Floor: 1, Capacity: 6, Status: "available"},
		{TableNumber: "06", Floor: 1, Capacity: 6, Status: "available"},
		{TableNumber: "07", Floor: 1, Capacity: 8, Status: "available"},
		{TableNumber: "08", Floor: 2, Capacity: 4, Status: "available"},
		{TableNumber: "09", Floor: 2, Capacity: 4, Status: "available"},
		{TableNumber: "10", Floor: 2, Capacity: 4, Status: "available"},
		{TableNumber: "11", Floor: 2, Capacity: 4, Status: "available"},
		{TableNumber: "12", Floor: 2, Capacity: 6, Status: "available"},
		{TableNumber: "13", Floor: 2, Capacity: 6, Status: "available"},
		{TableNumber: "14", Floor: 2, Capacity: 8, Status: "available"},
		{TableNumber: "15", Floor: 3, Capacity: 4, Status: "available"},
		{TableNumber: "16", Floor: 3, Capacity: 4, Status: "available"},
		{TableNumber: "17", Floor: 3, Capacity: 4, Status: "available"},
		{TableNumber: "18", Floor: 3, Capacity: 4, Status: "available"},
		{TableNumber: "19", Floor: 3, Capacity: 6, Status: "available"},
		{TableNumber: "20", Floor: 3, Capacity: 6, Status: "available"},
		{TableNumber: "21", Floor: 3, Capacity: 10, Status: "available"},
	}
	for _, table := range tables {
		var existing entity.Table
		if db.Where("table_number = ?", table.TableNumber).First(&existing).RowsAffected == 0 {
			if err := db.Create(&table).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

func seedSuperAdmin(db *gorm.DB) error {
	var existingUser entity.User
	if db.Where("email = ?", "superadmin@posapp.com").First(&existingUser).RowsAffected > 0 {
		return nil
	}

	var role entity.Role
	if err := db.Where("name = ?", "superadmin").First(&role).Error; err != nil {
		return err
	}

	passwordHash, err := utils.HashPassword("Admin@123")
	if err != nil {
		return err
	}

	user := &entity.User{
		Email:        "superadmin@posapp.com",
		Username:     "superadmin",
		PasswordHash: passwordHash,
		FullName:     "Super Administrator",
		Phone:        "+62 812 3456 7890",
		RoleID:       role.ID,
		ShiftStart:   "09:00",
		ShiftEnd:     "18:00",
		IsActive:     true,
	}

	return db.Create(user).Error
}

// data dummy untuk permission
func seedPermission(db *gorm.DB) error {
	permissions := []entity.Permission{
		{ID: 1, Code: "user:read", Description: "Melihat daftar dan detail user"},
		{ID: 2, Code: "user:create", Description: "Menambahkan user baru"},
		{ID: 3, Code: "user:update", Description: "Mengubah data user"},
		{ID: 4, Code: "user:delete", Description: "Menghapus user (soft delete)"},
	}
	for _, permission := range permissions {
		var existing entity.Permission
		if db.Where("id = ?", permission.ID).First(&existing).RowsAffected == 0 {
			if err := db.Create(&permission).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

// data dummy untuk role permission
func seedRolePermissions(db *gorm.DB) error {
	RolePermissions := []entity.RolePermisson{
		{RoleID: 1, PermissionID: 1},
		{RoleID: 1, PermissionID: 2},
		{RoleID: 1, PermissionID: 3},
		{RoleID: 1, PermissionID: 4},
		{RoleID: 2, PermissionID: 1},
		{RoleID: 2, PermissionID: 2},
		{RoleID: 2, PermissionID: 3},
		{RoleID: 2, PermissionID: 4},
	}
	for _, RolePermission := range RolePermissions {
		var existing entity.RolePermisson
		if db.Where("role_id = ? AND permission_id =?", RolePermission.RoleID, RolePermission.PermissionID).First(&existing).RowsAffected == 0 {
			if err := db.Create(&RolePermission).Error; err != nil {
				return err
			}
		}
	}
	return nil
}
