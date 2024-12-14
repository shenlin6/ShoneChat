package svc

import (
	"ShoneChat/apps/user/models"
	"ShoneChat/apps/user/rpc/internal/config"
	"ShoneChat/pkg/constant"
	"ShoneChat/pkg/ctxdata"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
)

type ServiceContext struct {
	Config config.Config
	*redis.Redis
	models.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config:     c,
		Redis:      redis.MustNewRedis(c.Redisx),
		UsersModel: models.NewUsersModel(sqlConn, c.Cache),
	}
}

func (svc *ServiceContext) SetRootToken() error {
	// 生成jwt
	systemToken, err := ctxdata.GetJwtToken(svc.Config.Jwt.AccessSecret, time.Now().Unix(), 999999999, constant.SYSTEM_ROOT_UID)
	if err != nil {
		return err
	}
	// 写入到redis
	return svc.Redis.Set(constant.REDIS_SYSTEM_ROOT_TOKEN, systemToken)
}
