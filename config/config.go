package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	DbType   string
	Database DatabaseConfig
	JWT      JWTConfig
	Redis    RedisConfig
}

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	URI      *string
	Host     *string
	Port     *string
	Username *string
	Password *string
	DbName   *string
	SSLMode  *string
	TimeZone *string
}

type JWTConfig struct {
	Secret          string
	AccessDuration  string
	RefreshDuration string
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
}

func LoadConfig(env string) error {
	if env == "" {
		env = "dev"
	}

	// Set the config file
	viper.SetConfigName(fmt.Sprintf("config.%s", env))
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	// Allow viper to read environment variables
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Try to read config file
	if err := viper.ReadInConfig(); err != nil {
		// If config file is not found, that's okay - we'll rely on env vars
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return fmt.Errorf("error reading config file: %v", err)
		}
		fmt.Printf("Config file not found, using environment variables\n")
	} else {
		fmt.Printf("Using config file: %s\n", viper.ConfigFileUsed())
	}

	fmt.Printf("Loaded environment: %s\n", env)
	return nil
}

// stringPtr returns a pointer to the string if it's not empty, otherwise returns nil
func stringPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func Init() *Config {
	var dbConfig DatabaseConfig
	env := viper.GetString("app.env")
	if env == "" {
		env = viper.GetString("APP_ENV") // fallback to env var
	}
	if err := LoadConfig(env); err != nil {
		fmt.Printf("Warning: %v\n", err)
	}

	dbType := viper.GetString("database.type")
	if dbType == "" {
		dbType = viper.GetString("DB_TYPE") // fallback to env var
		if dbType == "" {
			dbType = "postgresql"
		}
	}

	serverConfig := ServerConfig{
		Port: getConfigString("server.port", "SERVER_PORT"),
	}

	switch dbType {
	case "postgresql":
		dbConfig = DatabaseConfig{
			Host:     stringPtr(getConfigString("database.postgres.host", "POSTGRES_HOST")),
			Port:     stringPtr(getConfigString("database.postgres.port", "POSTGRES_PORT")),
			Username: stringPtr(getConfigString("database.postgres.username", "POSTGRES_USER")),
			Password: stringPtr(getConfigString("database.postgres.password", "POSTGRES_PASSWORD")),
			DbName:   stringPtr(getConfigString("database.postgres.dbname", "POSTGRES_DB")),
			SSLMode:  stringPtr(getConfigString("database.postgres.sslmode", "POSTGRES_SSL_MODE")),
			TimeZone: stringPtr(getConfigString("database.postgres.timezone", "POSTGRES_TIMEZONE")),
		}
	case "mongodb":
		dbConfig = DatabaseConfig{
			URI:      stringPtr(getConfigString("database.mongodb.uri", "MONGODB_URI")),
			Host:     stringPtr(getConfigString("database.mongodb.host", "MONGO_HOST")),
			Port:     stringPtr(getConfigString("database.mongodb.port", "MONGO_PORT")),
			Username: stringPtr(getConfigString("database.mongodb.username", "MONGO_USER")),
			Password: stringPtr(getConfigString("database.mongodb.password", "MONGO_PASSWORD")),
			DbName:   stringPtr(getConfigString("database.mongodb.dbname", "MONGO_DB")),
		}
	default:
		dbConfig = DatabaseConfig{
			URI:      stringPtr(getConfigString("database.mongodb.uri", "MONGODB_URI")),
			Host:     stringPtr(getConfigString("database.mongodb.host", "MONGO_HOST")),
			Port:     stringPtr(getConfigString("database.mongodb.port", "MONGO_PORT")),
			Username: stringPtr(getConfigString("database.mongodb.username", "MONGO_USER")),
			Password: stringPtr(getConfigString("database.mongodb.password", "MONGO_PASSWORD")),
			DbName:   stringPtr(getConfigString("database.mongodb.dbname", "MONGO_DB")),
		}
	}

	jwtConfig := JWTConfig{
		Secret:          getConfigString("jwt.secret", "JWT_SECRET"),
		AccessDuration:  getConfigString("jwt.access_duration", "JWT_ACCESS_DURATION"),
		RefreshDuration: getConfigString("jwt.refresh_duration", "JWT_REFRESH_DURATION"),
	}

	redisConfig := RedisConfig{
		Host:     getConfigString("redis.host", "REDIS_HOST"),
		Port:     getConfigString("redis.port", "REDIS_PORT"),
		Password: getConfigString("redis.password", "REDIS_PASSWORD"),
	}

	return &Config{
		Server:   serverConfig,
		DbType:   dbType,
		Database: dbConfig,
		JWT:      jwtConfig,
		Redis:    redisConfig,
	}
}

// getConfigString tries to get value from YAML config first, then falls back to environment variable
func getConfigString(yamlKey, envKey string) string {
	value := viper.GetString(yamlKey)
	if value == "" {
		value = viper.GetString(envKey)
	}
	return value
}
