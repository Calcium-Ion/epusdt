package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

var (
	AppDebug    bool
	MysqlDns    string
	RuntimePath string
	LogSavePath string
	StaticPath  string
	TgBotToken  string
	TgProxy     string
	TgManage    int64
	UsdtRate    float64
)

// getEnv retrieves env var with fallback
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

// getEnvBool retrieves bool env var with fallback
func getEnvBool(key string, fallback bool) bool {
	if value := os.Getenv(key); value != "" {
		b, err := strconv.ParseBool(value)
		if err == nil {
			return b
		}
	}
	return fallback
}

// getEnvFloat retrieves float64 env var with fallback
func getEnvFloat(key string, fallback float64) float64 {
	if value := os.Getenv(key); value != "" {
		f, err := strconv.ParseFloat(value, 64)
		if err == nil {
			return f
		}
	}
	return fallback
}

func Init() {
	// Load .env file if it exists
	_ = godotenv.Load()

	gwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	AppDebug = getEnvBool("APP_DEBUG", false)
	StaticPath = getEnv("STATIC_PATH", "static")
	RuntimePath = fmt.Sprintf("%s%s", gwd, getEnv("RUNTIME_ROOT_PATH", "/runtime"))
	LogSavePath = fmt.Sprintf("%s%s", RuntimePath, getEnv("LOG_SAVE_PATH", "/logs"))

	// Build MySQL DNS
	MysqlDns = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		getEnv("MYSQL_HOST", "localhost"),
		getEnv("MYSQL_PORT", "3306"),
		os.Getenv("MYSQL_DATABASE"))

	TgBotToken = os.Getenv("TG_BOT_TOKEN")
	TgProxy = os.Getenv("TG_PROXY")
	if manage := os.Getenv("TG_MANAGE"); manage != "" {
		TgManage, _ = strconv.ParseInt(manage, 10, 64)
	}
}

func GetAppVersion() string {
	return getEnv("APP_VERSION", "0.0.2")
}

func GetAppName() string {
	return getEnv("APP_NAME", "epusdt")
}

func GetAppUri() string {
	return os.Getenv("APP_URI")
}

func GetApiAuthToken() string {
	return os.Getenv("API_AUTH_TOKEN")
}

func GetUsdtRate() float64 {
	forcedRate := getEnvFloat("FORCED_USDT_RATE", 0)
	if forcedRate > 0 {
		return forcedRate
	}
	if UsdtRate <= 0 {
		return 6.4
	}
	return UsdtRate
}

func GetOrderExpirationTime() int {
	if value := os.Getenv("ORDER_EXPIRATION_TIME"); value != "" {
		if timer, err := strconv.Atoi(value); err == nil && timer > 0 {
			return timer
		}
	}
	return 10
}

func GetOrderExpirationTimeDuration() time.Duration {
	return time.Minute * time.Duration(GetOrderExpirationTime())
}
