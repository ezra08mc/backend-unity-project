package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ezra08mc/backend-unity-project/config"
	"github.com/ezra08mc/backend-unity-project/config/database"
	"github.com/ezra08mc/backend-unity-project/config/middleware"
	"github.com/ezra08mc/backend-unity-project/controller"
	"github.com/gin-gonic/gin"

	"github.com/ezra08mc/backend-unity-project/repository"
	"github.com/ezra08mc/backend-unity-project/service"
	"gorm.io/gorm"
)

func Run() {
	log.Println("Starting application...")

	cfg := config.Get()
	if cfg == nil {
		log.Fatal("Failed to load configuration")
		return
	}

	db, _, err := database.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
		return
	}

	startServer(cfg, db)
}

func startServer(cfg *config.AppConfigurationMap, db *gorm.DB) {
	repo := repository.New(db)
	serv := service.New(repo)

	if cfg.IsProduction {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.GlobalRateLimiter(cfg.RateLimitRPS, cfg.RateLimitBurst, map[string]struct{}{}))
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Static("/static", "./static")

	controller.New(r, serv)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      r,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Printf("Server is running on port %d", cfg.Port)
	log.Fatal(srv.ListenAndServe())
}
