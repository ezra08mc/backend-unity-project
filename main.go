package main

// @title Mini Project Backend Unity API
// @version 1.0
// @description Documentation for Backend Unity Mini Project
// @host localhost:8080
// @BasePath /

import (
	"fmt"
	"os"

	"github.com/ezra08mc/backend-unity-project/config"
	dbConfig "github.com/ezra08mc/backend-unity-project/config/database"
	"github.com/ezra08mc/backend-unity-project/config/pkg/token"
	"github.com/ezra08mc/backend-unity-project/config/server"
	dbMigration "github.com/ezra08mc/backend-unity-project/database"
)

func main() {
	config.Load()
	token.Load()

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "migrate":
			runMigrations()
			return
		case "reset":
			runReset()
			return
		case "seed":
			runSeedOnly()
			return
		default:
			fmt.Println("Unknown command. Use: migrate | reset | seed")
			return
		}
	}

	db, _, err := dbConfig.ConnectDB()
	if err != nil {
		panic(fmt.Errorf("failed to connect database: %w", err))
	}

	if !db.Migrator().HasTable(&dbMigration.User{}) {
		fmt.Println("🔄 First run detected, running auto-migration...")
		if err := dbMigration.RunMigration(db); err != nil {
			panic(fmt.Errorf("auto-migration failed: %w", err))
		}
	} else {
		fmt.Println("✅ Database already migrated, ensuring seed...")
		if err := dbMigration.Seed(db); err != nil {
			panic(fmt.Errorf("seeding failed: %w", err))
		}
	}

	server.Run()
}

func runMigrations() {
	db, _, err := dbConfig.ConnectDB()
	if err != nil {
		panic(err)
	}

	if err := dbMigration.RunMigration(db); err != nil {
		panic(err)
	}
}

func runReset() {
	db, _, err := dbConfig.ConnectDB()
	if err != nil {
		panic(err)
	}

	fmt.Println("🗑️   Dropping all tables...")
	err = db.Migrator().DropTable(
		&dbMigration.User{},
		&dbMigration.Todo{},
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("🔄 Recreating tables with AutoMigrate...")
	if err := dbMigration.RunMigration(db); err != nil {
		panic(err)
	}

	fmt.Println("✅ Database reset completed")
}

func runSeedOnly() {
	db, _, err := dbConfig.ConnectDB()
	if err != nil {
		panic(err)
	}

	fmt.Println("🌱 Running seed only...")
	if err := dbMigration.Seed(db); err != nil {
		panic(err)
	}
	fmt.Println("✅ Seeding completed")
}
