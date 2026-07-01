package database

import (
	"fmt"

	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) error {
	fmt.Println("Running migrations...")

	if err := db.AutoMigrate(
		&User{},
		&Todo{},
	); err != nil {
		return fmt.Errorf("failed to migrate: %w", err)
	}

	fmt.Println("Migrations completed")

	fmt.Println("Seeding database...")
	if err := Seed(db); err != nil {
		return fmt.Errorf("failed to seed: %w", err)
	}

	return nil
}
