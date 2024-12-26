package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

var (
	AppDebug           bool
	MysqlDns           string
	RuntimePath        string
	LogSavePath        string
	StaticPath         string
	TgBotToken         string
	TgProxy            string
	TgManage           int64
	UsdtRate           float64
	RedisHost          string
	RedisPort          string
	RedisDB            int
	RedisPassword      string
	RedisPoolSize      int
	RedisMaxRetries    int
	RedisIdleTimeout   time.Duration
	HttpListen         string
	MysqlMaxIdleConns  int
	MysqlMaxOpenConns  int
	MysqlMaxLifeTime   int
	MysqlTablePrefix   string
	QueueConcurrency   int
	QueueLevelCritical int
	QueueLevelDefault  int
	QueueLevelLow      int
	LogMaxSize         int
	LogMaxBackups      int
	LogMaxAge          int
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

func getEnvInt(key string, fallback int) int {
	if value := os.Getenv(key); value != "" {
		if i, err := strconv.Atoi(value); err == nil {
			return i
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

	// Add Redis configuration
	RedisHost = getEnv("REDIS_HOST", "localhost")
	RedisPort = getEnv("REDIS_PORT", "6379")
	RedisDB = getEnvInt("REDIS_DB", 0)
	RedisPassword = getEnv("REDIS_PASSWORD", "")
	RedisPoolSize = getEnvInt("REDIS_POOL_SIZE", 10)
	RedisMaxRetries = getEnvInt("REDIS_MAX_RETRIES", 3)
	RedisIdleTimeout = time.Second * time.Duration(getEnvInt("REDIS_IDLE_TIMEOUT", 300))

	// HTTP configs
	HttpListen = getEnv("HTTP_LISTEN", ":8080")

	// MySQL configs
	MysqlMaxIdleConns = getEnvInt("MYSQL_MAX_IDLE_CONNS", 10)
	MysqlMaxOpenConns = getEnvInt("MYSQL_MAX_OPEN_CONNS", 100)
	MysqlMaxLifeTime = getEnvInt("MYSQL_MAX_LIFE_TIME", 3600)
	MysqlTablePrefix = getEnv("MYSQL_TABLE_PREFIX", "epusdt_")

	// Queue configs
	QueueConcurrency = getEnvInt("QUEUE_CONCURRENCY", 10)
	QueueLevelCritical = getEnvInt("QUEUE_LEVEL_CRITICAL", 6)
	QueueLevelDefault = getEnvInt("QUEUE_LEVEL_DEFAULT", 3)
	QueueLevelLow = getEnvInt("QUEUE_LEVEL_LOW", 1)

	// Log configs
	LogMaxSize = getEnvInt("LOG_MAX_SIZE", 100)
	LogMaxBackups = getEnvInt("LOG_MAX_BACKUPS", 3)
	LogMaxAge = getEnvInt("LOG_MAX_AGE", 28)
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
		return 7.3
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
