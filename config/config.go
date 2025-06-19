package config

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
)

// Config represents the application configuration
type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	JWT      JWTConfig      `mapstructure:"jwt"`
	Redis    RedisConfig    `mapstructure:"redis"`
}

// AppConfig represents application-level configuration
type AppConfig struct {
	Env string `mapstructure:"env"`
}

// ServerConfig represents server configuration
type ServerConfig struct {
	Port string `mapstructure:"port"`
}

// DatabaseConfig represents database configuration
type DatabaseConfig struct {
	Type     string         `mapstructure:"type"`
	Postgres PostgresConfig `mapstructure:"postgres"`
	MongoDB  MongoDBConfig  `mapstructure:"mongodb"`
}

// PostgresConfig represents PostgreSQL configuration
type PostgresConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
	SSLMode  string `mapstructure:"sslmode"`
	Timezone string `mapstructure:"timezone"`
}

// MongoDBConfig represents MongoDB configuration
type MongoDBConfig struct {
	URI      string `mapstructure:"uri"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
}

// JWTConfig represents JWT configuration
type JWTConfig struct {
	Secret          string        `mapstructure:"secret"`
	AccessDuration  time.Duration `mapstructure:"access_duration"`
	RefreshDuration time.Duration `mapstructure:"refresh_duration"`
}

// RedisConfig represents Redis configuration
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
}

// LoadConfig loads configuration from file using viper
func LoadConfig(configPath string) (*Config, error) {
	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	// Set default values
	viper.SetDefault("server.port", "8080")
	viper.SetDefault("database.type", "postgresql")
	viper.SetDefault("database.postgres.sslmode", "disable")
	viper.SetDefault("database.postgres.timezone", "UTC")
	viper.SetDefault("jwt.access_duration", "15m")
	viper.SetDefault("jwt.refresh_duration", "10080m")
	viper.SetDefault("redis.host", "localhost")
	viper.SetDefault("redis.port", "6379")

	// Read configuration file
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	// Create config instance
	var config Config

	// Unmarshal into struct using mapstructure
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	// Validate required fields
	if err := validateConfig(&config); err != nil {
		return nil, fmt.Errorf("config validation failed: %w", err)
	}

	return &config, nil
}

// LoadConfigFromEnv loads configuration based on environment
func LoadConfigFromEnv(env string) (*Config, error) {
	var configPath string

	switch env {
	case "development", "dev":
		configPath = "config/dev/config.dev.yaml"
	case "production", "prod":
		configPath = "config/prod/config.prod.yaml"
	default:
		configPath = "config/example/config.example.yaml"
	}

	return LoadConfig(configPath)
}

// validateConfig validates required configuration fields
func validateConfig(config *Config) error {
	if config.JWT.Secret == "" {
		return fmt.Errorf("JWT secret is required")
	}

	if config.Database.Type == "" {
		return fmt.Errorf("database type is required")
	}

	// Validate database configuration based on type
	switch config.Database.Type {
	case "postgresql", "postgres":
		if config.Database.Postgres.Host == "" {
			return fmt.Errorf("postgres host is required")
		}
		if config.Database.Postgres.DBName == "" {
			return fmt.Errorf("postgres database name is required")
		}
	case "mongodb", "mongo":
		if config.Database.MongoDB.URI == "" && config.Database.MongoDB.Host == "" {
			return fmt.Errorf("mongodb URI or host is required")
		}
	}

	return nil
}

// GetDatabaseDSN returns database connection string based on configuration
func (c *Config) GetDatabaseDSN() string {
	switch c.Database.Type {
	case "postgresql", "postgres":
		return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
			c.Database.Postgres.Host,
			c.Database.Postgres.Port,
			c.Database.Postgres.Username,
			c.Database.Postgres.Password,
			c.Database.Postgres.DBName,
			c.Database.Postgres.SSLMode,
			c.Database.Postgres.Timezone,
		)
	case "mongodb", "mongo":
		if c.Database.MongoDB.URI != "" {
			return c.Database.MongoDB.URI
		}
		return fmt.Sprintf("mongodb://%s:%s@%s:%s/%s",
			c.Database.MongoDB.Username,
			c.Database.MongoDB.Password,
			c.Database.MongoDB.Host,
			c.Database.MongoDB.Port,
			c.Database.MongoDB.DBName,
		)
	}
	return ""
}

// GetRedisAddr returns Redis address
func (c *Config) GetRedisAddr() string {
	return fmt.Sprintf("%s:%s", c.Redis.Host, c.Redis.Port)
}

// IsProduction checks if the application is running in production mode
func (c *Config) IsProduction() bool {
	return c.App.Env == "production" || c.App.Env == "prod"
}

// IsDevelopment checks if the application is running in development mode
func (c *Config) IsDevelopment() bool {
	return c.App.Env == "development" || c.App.Env == "dev"
}

// PrintConfig prints the configuration (excluding sensitive data)
func (c *Config) PrintConfig() {
	log.Printf("Configuration loaded:")
	log.Printf("  Environment: %s", c.App.Env)
	log.Printf("  Server Port: %s", c.Server.Port)
	log.Printf("  Database Type: %s", c.Database.Type)
	log.Printf("  Redis: %s", c.GetRedisAddr())
	log.Printf("  JWT Access Duration: %s", c.JWT.AccessDuration)
	log.Printf("  JWT Refresh Duration: %s", c.JWT.RefreshDuration)
}
