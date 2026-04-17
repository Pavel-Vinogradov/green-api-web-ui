package container

import (
	"green-api-web-ui/internal/config"
	"net/http"
)

type Container struct {
	Config *config.AppConfig

	HTTPClient *http.Client
}

func NewContainer(cfg *config.AppConfig) *Container {
	httpClient := &http.Client{
		Timeout: cfg.HTTPClient.Timeout,
	}

	return &Container{
		Config:     cfg,
		HTTPClient: httpClient,
	}
}
