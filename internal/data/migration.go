package data

import (
	"log"
	"project-POS-APP-golang-team-float/internal/data/entity"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {


	err := db.AutoMigrate(
		&entity.Role{},
		&entity.User{},
		&entity.Session{},
		&entity.Permission{},
		&entity.RolePermisson{},
		&entity.Category{},
		&entity.PaymentMethod{},
		&entity.Product{},
		&entity.Table{},
		&entity.Order{},
		&entity.OrderItem{},
		&entity.OTPCode{},
		&entity.UserPermission{},
		&entity.Notification{},
		&entity.Customer{},
		&entity.Reservation{},
	)

	if err != nil {
		return err
	}

	if err := createTriggers(db); err != nil {
		return err
	}

	log.Println("Database Migration & Triggers setup completed successfully.")
	return nil
}

func createTriggers(db *gorm.DB) error {
	queryFunc := `
	CREATE OR REPLACE FUNCTION update_updated_at_column()
	RETURNS TRIGGER AS $$
	BEGIN
		NEW.updated_at = NOW();
		RETURN NEW;
	END;
	$$ language 'plpgsql';
	`
	if err := db.Exec(queryFunc).Error; err != nil {
		return err
	}

	tables := []string{
		"roles", "users", "categories", "products",
		"tables", "orders", "sessions",
	}

	for _, table := range tables {
		if !db.Migrator().HasTable(table) {
			continue
		}

		db.Exec("DROP TRIGGER IF EXISTS update_" + table + "_updated_at ON " + table)

		createTrigger := `
		CREATE TRIGGER update_` + table + `_updated_at
		BEFORE UPDATE ON ` + table + `
		FOR EACH ROW
		EXECUTE FUNCTION update_updated_at_column();
		`
		if err := db.Exec(createTrigger).Error; err != nil {
			return err
		}
	}

	return nil
}
