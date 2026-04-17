package config

type (
	AppConfig struct {
		HTTPClient HTTPClientConfig `mapstructure:"http_client"`
		Server     ServerConfig     `mapstructure:"server"`
	}
)
