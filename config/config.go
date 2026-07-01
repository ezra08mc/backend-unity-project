package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/ezra08mc/backend-unity-project/config/pkg/utils"
	"github.com/joho/godotenv"
)

type AppConfigurationMap struct {
	Port                 int     // Port is the port number that the server will listen to.
	IsProduction         bool    // IsProduction is a flag that indicates whether the application is running in production mode.
	DbURI                string  // Database connection.
	AccessTokenLifeTime  uint    // AccessTokenLifeTime is the lifetime of the access token in seconds.
	RefreshTokenLifeTime uint    // RefreshTokenLifeTime is the lifetime of the refresh token in seconds.
	PrivateKeyPath       string  // Path to the private key file.
	PublicKeyPath        string  // Path to the public key file.
	BaseURL              string  // BaseURL is the base URL of the application, used for generating absolute URLs.
	RateLimitRPS         float64 // Global request-per-second limit (if <=0 disabled)
	RateLimitBurst       int     // Burst size for rate limiter token bucket
}

var config *AppConfigurationMap

func Get() *AppConfigurationMap {
	return config
}

func Load() {
	log.Println("Loading config from environment...")

	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading environment variables, try to get from environtment OS")
	}

	port, err := strconv.Atoi(os.Getenv("PORT"))

	if err != nil {
		port = 8080
	}

	isProduction := utils.SafeCompareString(os.Getenv("IS_PRODUCTION"), "true")

	AccessTokenLifeTime, err := strconv.Atoi(os.Getenv("ACCESS_TOKEN_LIFE_TIME"))
	if err != nil {
		AccessTokenLifeTime = 3600 
	}

	RefreshTokenLifeTime, err := strconv.Atoi(os.Getenv("REFRESH_TOKEN_LIFE_TIME"))
	if err != nil {
		RefreshTokenLifeTime = 86400 
	}

	PrivateKeyPath := os.Getenv("PRIVATE_KEY")
	if PrivateKeyPath == "" {
		log.Fatalf("PRIVATE_KEY_PATH environment variable is not set, check your .env file")
	}

	PublicKeyPath := os.Getenv("PUBLIC_KEY")
	if PublicKeyPath == "" {
		log.Fatalf("PUBLIC_KEY_PATH environment variable is not set, check your .env file")
	}

	BaseURL := os.Getenv("BASE_URL")
	if BaseURL == "" {
		BaseURL = fmt.Sprintf("http://localhost:%d", port)
	}

	// Global rate limiter configuration
	rpsEnv := os.Getenv("RATE_LIMIT_RPS")
	burstEnv := os.Getenv("RATE_LIMIT_BURST")
	rps := 10.0 // sensible default
	burst := 20 // sensible default
	if v, err := strconv.ParseFloat(rpsEnv, 64); err == nil && v > 0 {
		rps = v
	}
	if v, err := strconv.Atoi(burstEnv); err == nil && v > 0 {
		burst = v
	}

	config = &AppConfigurationMap{
		Port:                 port,
		IsProduction:         isProduction,
		DbURI:                loadDatabaseConfig(),
		AccessTokenLifeTime:  uint(AccessTokenLifeTime),
		RefreshTokenLifeTime: uint(RefreshTokenLifeTime),
		PrivateKeyPath:       PrivateKeyPath,
		PublicKeyPath:        PublicKeyPath,
		BaseURL:              BaseURL,
		RateLimitRPS:         rps,
		RateLimitBurst:       burst,
	}
}

func loadDatabaseConfig() string {
	user := getFromEnv("DB_USER")
	pass := getFromEnv("DB_PASS")
	name := getFromEnv("DB_NAME")
	host := getFromEnv("DB_HOST")
	port := getFromEnv("DB_PORT")
	timeZone := getFromEnv("DB_TIME_ZONE")

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=%s", host, user, pass, name, port, timeZone)
}

func getFromEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("%s Environment variable is not set", value)
	}

	return value
}
