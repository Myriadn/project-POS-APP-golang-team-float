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
	if err := seedProducts(db); err != nil {
		return err
	}
	if err := seedSampleOrders(db); err != nil {
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

func seedSampleOrders(db *gorm.DB) error {
	// Check if orders already exist
	var count int64
	db.Model(&entity.Order{}).Count(&count)
	if count > 0 {
		return nil
	}

	// Get superadmin user
	var user entity.User
	if err := db.Where("email = ?", "superadmin@posapp.com").First(&user).Error; err != nil {
		return err
	}

	// Get payment method
	var paymentMethod entity.PaymentMethod
	if err := db.Where("name = ?", "Cash").First(&paymentMethod).Error; err != nil {
		return err
	}

	// Get some tables
	var tables []entity.Table
	if err := db.Limit(5).Find(&tables).Error; err != nil {
		return err
	}

	// Get products
	var products []entity.Product
	if err := db.Find(&products).Error; err != nil {
		return err
	}

	if len(products) == 0 || len(tables) == 0 {
		return nil
	}

	// Create sample orders for the last 30 days
	taxRate := 5.00
	for i := range 20 {
		orderDate := time.Now().AddDate(0, 0, -i)
		tableID := tables[i%len(tables)].ID
		paymentMethodID := paymentMethod.ID

		// Calculate order totals
		subtotal := 0.0
		orderItems := []entity.OrderItem{}

		// Add 2-4 random products to each order
		numItems := 2 + (i % 3)
		for j := range numItems {
			product := products[(i+j)%len(products)]
			quantity := 1 + (j % 3)
			unitPrice := product.Price
			totalPrice := unitPrice * float64(quantity)
			subtotal += totalPrice

			orderItems = append(orderItems, entity.OrderItem{
				ProductID:  product.ID,
				Quantity:   quantity,
				UnitPrice:  unitPrice,
				TotalPrice: totalPrice,
			})
		}

		taxAmount := subtotal * (taxRate / 100)
		total := subtotal + taxAmount

		// Determine order status
		status := "completed"
		if i%7 == 0 {
			status = "cancelled"
		} else if i%5 == 0 {
			status = "in_process"
		}

		order := &entity.Order{
			OrderNumber:     fmt.Sprintf("ORD%s%03d", orderDate.Format("20060102"), i+1),
			TableID:         &tableID,
			UserID:          user.ID,
			PaymentMethodID: &paymentMethodID,
			CustomerName:    fmt.Sprintf("Customer %d", i+1),
			Status:          status,
			Subtotal:        subtotal,
			TaxRate:         taxRate,
			TaxAmount:       taxAmount,
			Total:           total,
			OrderDate:       orderDate,
		}

		if err := db.Create(order).Error; err != nil {
			return err
		}

		// Create order items
		for _, item := range orderItems {
			item.OrderID = order.ID
			if err := db.Create(&item).Error; err != nil {
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
