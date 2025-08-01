package common_config

type Config struct {
	Environment string
	Port        string
	DatabaseURL string
	LogLevel    string
	JWTSecret   string

	// OpenTelemetry configuration
	OtelEnabled      bool
	OtelServiceName  string
	OtelExporterType string
	OtelEndpoint     string

	// Temporal configuration
	TemporalHost      string
	TemporalPort      string
	TemporalNamespace string
	TaskQueue         string

	// OPA configuration
	OPAEnabled bool
	OPAURL     string
}
