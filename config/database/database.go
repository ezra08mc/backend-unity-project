package database

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/ezra08mc/backend-unity-project/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB() (*gorm.DB, *sql.DB, error) {
	cfg := config.Get() 
	sqlLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), 
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			Colorful:                  true,
		})

	log.Println("connecting to database...")

	db, err := gorm.Open(postgres.Open(cfg.DbURI), &gorm.Config{
		Logger:                 sqlLogger,
		SkipDefaultTransaction: true, 
		AllowGlobalUpdate: false, 
	})
	if err != nil {
		log.Fatalf("error connecting to SQL: %v", err)
	}

	log.Println("setting database connection configuration...")

	sqlDB, err := db.DB()

	if err != nil {
		log.Fatalf("error setting database connection configuration: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)

	sqlDB.SetConnMaxLifetime(time.Hour)
	return db, sqlDB, nil
}
