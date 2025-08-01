package common_config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var env *Config

func Load() *Config {
	if env != nil {
		return env
	}

	if os.Getenv("IS_DOCKER") == "" {
		err := godotenv.Load()

		if err != nil {
			log.Fatalf("Unable to load envs: %v", err)
		}
	}

	env = &Config{
		Environment: getEnv("ENVIRONMENT", "development"),
		Port:        getEnv("PORT", "8080"),
		DatabaseURL: getEnv("DATABASE_URL", "postgres://yugabyte:yugabyte@localhost:5433/yugabyte?sslmode=disable"),
		LogLevel:    getEnv("LOG_LEVEL", "debug"),
		JWTSecret:   getEnv("JWT_SECRET", "your-secret-key"),

		// OpenTelemetry configuration
		OtelEnabled:      getEnvBool("OTEL_ENABLED", false),
		OtelServiceName:  getEnv("OTEL_SERVICE_NAME", "golang-boilerplate"),
		OtelExporterType: getEnv("OTEL_EXPORTER_TYPE", "jaeger"),
		OtelEndpoint:     getEnv("OTEL_ENDPOINT", "http://localhost:14268/api/traces"),

		// Temporal configuration
		TemporalHost:      getEnv("TEMPORAL_HOST", ""),
		TemporalPort:      getEnv("TEMPORAL_PORT", "7233"),
		TemporalNamespace: getEnv("TEMPORAL_NAMESPACE", "default"),
		TaskQueue:         getEnv("TASK_QUEUE", "user-onboarding"),

		// OPA configuration
		OPAEnabled: getEnvBool("OPA_ENABLED", false),
		OPAURL:     getEnv("OPA_URL", "http://localhost:8181"),
	}

	return env
}
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		return value == "true" || value == "1"
	}
	return defaultValue
}
