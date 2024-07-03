package config

import (
	"os"
)

var (
	MongoUrl = GetEnv("MONGODB_URL", "mongodb://localhost:27017")
	// prod
	// MongoUrl        = GetEnv("MONGODB_URL", "mongodb://dexter:Cetnbcm88Cetnbcm88$@localhost:27017/?authSource=admin&readPreference=primary&authMechanism=SCRAM-SHA-256&appname=MongoDB%20Compass&ssl=false")
	MongoDatabase = GetEnv("MONGODB_DATABASE", "BOTDB")
)

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}
