package env

import (
	"os"
)

func IsDevelopment() bool {
	return os.Getenv("APP_ENV") == "development"
}

func IsStaging() bool {
	return os.Getenv("APP_ENV") == "staging"
}

func IsProduction() bool {
	return os.Getenv("APP_ENV") == "production"
}

func GoogleCloudProject() string {
	return os.Getenv("GOOGLE_CLOUD_PROJECT")
}

func GoogleCloudSqlSource() string {
	return os.Getenv("GOOGLE_CLOUD_SQL_SOURCE")
}
