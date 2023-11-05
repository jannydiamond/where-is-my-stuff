package config

type AppConfig struct {
	Server ServerConfig
}

type ServerConfig struct {
	Port int
}

func LoadConfig() *AppConfig {
	// Load configuration from environment variables or config files
	config := &AppConfig{
		Server: ServerConfig{
			Port: 4000,
		},
	}

	return config
}
