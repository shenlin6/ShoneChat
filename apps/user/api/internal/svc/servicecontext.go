package svc

import (
	"ShoneChat/apps/user/api/internal/config"
	"ShoneChat/apps/user/rpc/userclient"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	*redis.Redis
	userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		Redis: redis.MustNewRedis(c.Redisx),
		User:  userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
