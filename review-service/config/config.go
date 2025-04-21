package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AWSRegion     string
	S3Bucket      string
	S3Prefix      string
	DatabaseURL   string
	WorkerCount   int
}

func Load() Config {
	_ = godotenv.Load(".env")

	return Config{
		AWSRegion:     os.Getenv("AWS_REGION"),
		S3Bucket:      os.Getenv("S3_BUCKET"),
		S3Prefix:      os.Getenv("S3_PREFIX"),
		DatabaseURL:   os.Getenv("DATABASE_URL"),
		WorkerCount:   5,
	}
}

