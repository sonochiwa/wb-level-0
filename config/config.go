package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Server        Server
	Cors          Cors
	Postgres      Postgres
	NatsStreaming NatsStreaming
}

type Server struct {
	Host         string
	Port         string
	ReadTimeout  int
	WriteTimeout int
	IdleTimeout  int
}

type Cors struct {
	AllowedOrigins   []string
	MaxAge           int
	AllowCredentials bool
}

type Postgres struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

type NatsStreaming struct {
	ClusterID   string
	ClientID    string
	ChannelName string
}

func init() {
	loadConfig()
}

func loadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func GetConfig() *Config {
	return &Config{
		Server: Server{
			Host:         getEnv("HOST", "0.0.0.0"),
			Port:         getEnv("PORT", "8080"),
			ReadTimeout:  getEnvAsInt("READ_TIMEOUT", 60),
			WriteTimeout: getEnvAsInt("WRITE_TIMEOUT", 60),
			IdleTimeout:  getEnvAsInt("IDLE_TIMEOUT", 60),
		},
		Cors: Cors{
			AllowedOrigins:   getEnvAsSlice("ALLOWED_ORIGINS", []string{"*"}, ", "),
			AllowCredentials: getEnvAsBool("ALLOW_CREDENTIALS", true),
			MaxAge:           getEnvAsInt("MAX_AGE", 300),
		},
		Postgres: Postgres{
			User:     getEnv("POSTGRES_USER", "root"),
			Password: getEnv("POSTGRES_PASSWORD", "root"),
			Host:     getEnv("POSTGRES_HOST", "localhost"),
			Port:     getEnv("POSTGRES_PORT", "5432"),
			Database: getEnv("POSTGRES_DATABASE", "wb_db"),
		},
		NatsStreaming: NatsStreaming{
			ClusterID:   getEnv("NATS_CLUSTER_ID", "root"),
			ClientID:    getEnv("NATS_CLIENT_ID", "root"),
			ChannelName: getEnv("NATS_CHANNEL_NAME", "root"),
		},
	}
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if valueStr, exists := os.LookupEnv(key); exists {
		if value, err := strconv.Atoi(valueStr); err == nil {
			return value
		}
	}
	return defaultValue
}

func getEnvAsBool(key string, defaultValue bool) bool {
	if valueStr, exists := os.LookupEnv(key); exists {
		if value, err := strconv.ParseBool(valueStr); err == nil {
			return value
		}
	}
	return defaultValue
}

func getEnvAsSlice(key string, defaultValue []string, sep string) []string {
	if valueStr, exists := os.LookupEnv(key); exists {
		value := strings.Split(valueStr, sep)
		return value
	}
	return defaultValue
}
