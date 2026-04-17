package container

import "green-api-web-ui/internal/config"

type Container struct {
	Config *config.AppConfig
}

func NewContainer(cfg *config.AppConfig) *Container {
	return &Container{}
}
