package svc

import (
	"ShoneChat/apps/task/mq/internal/config"
)

type ServiceContext struct {
	config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
