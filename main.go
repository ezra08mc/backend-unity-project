package main

// @title Mini Project Backend Unity API
// @version 1.0
// @description Documentation for Backend Unity Mini Project
// @host localhost:8080
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Enter your token in the following format: "Bearer <token>" (Example: Bearer eyJhbGci...)
// @scheme bearer
// @bearerFormat JWT

import (
	"fmt"
	"os"

	"github.com/ezra08mc/backend-unity-project/config"
	dbConfig "github.com/ezra08mc/backend-unity-project/config/database"
	"github.com/ezra08mc/backend-unity-project/config/pkg/token"
	"github.com/ezra08mc/backend-unity-project/config/server"
	"github.com/ezra08mc/backend-unity-project/database"
    
	_ "github.com/ezra08mc/backend-unity-project/docs"
)

func main() {
	config.Load()
	token.Load()

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "migrate":
			runMigrations()
			return
		case "seed":
			runSeedOnly()
			return
		case "reset":
			runReset()
			return
		default:
			fmt.Println("Unknown command. Use: migrate | seed | reset")
			return
		}
	}

	server.Run()
}

func runMigrations() {
	db, _, err := dbConfig.ConnectDB()
	if err != nil {
		panic(err)
	}
	fmt.Println("🔄 Running database migration...")
	if err := database.RunMigration(db); err != nil {
		panic(err)
	}
	fmt.Println("✅ Migration completed")
}

func runSeedOnly() {
	db, _, err := dbConfig.ConnectDB()
	if err != nil {
		panic(err)
	}
	fmt.Println("🌱 Running seed only...")
	if err := database.Seed(db); err != nil {
		panic(err)
	}
	fmt.Println("✅ Seeding completed")
}

func runReset() {
	db, _, err := dbConfig.ConnectDB()
	if err != nil {
		panic(err)
	}
	fmt.Println("🗑️ Dropping all tables...")
	err = db.Migrator().DropTable(&database.Todo{}, &database.User{})
	if err != nil {
		panic(err)
	}
	runMigrations()
	fmt.Println("✅ Database reset completed")
}