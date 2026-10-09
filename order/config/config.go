package config

import (
	"log"
	"os"
	"strconv"
)

// GetEnv retrieves the application runtime environment setting.
func GetEnv() string {
	return getEnvironmentValue("ENV")
}

// GetDataSourceURL retrieves the database connection string.
func GetDataSourceURL() string {
	return getEnvironmentValue("DATA_SOURCE_URL")
}

// GetApplicationPort retrieves and parses the application port into an integer.
func GetApplicationPort() int {
	portSrt := getEnvironmentValue("APPLICATION_PORT")

	// Parse string port to integer and fail on invalid format
	port, err := strconv.Atoi(portSrt)
	if err != nil {
		log.Fatalf("port: %s is invalid", portSrt)
	}

	return port
}

// GetPaymentServiceUrl retrieves the payment service network endpoint.
func GetPaymentServiceUrl() string {
	return getEnvironmentValue("PAYMENT_SERVICE_URL")
}

// getEnvironmentValue fetches an env variable or terminates the application if missing.
func getEnvironmentValue(key string) string {

	if os.Getenv(key) == "" {
		log.Fatalf("%s environment variable is missing.", key)
	}
	return os.Getenv(key)
}
