package svc

import (
	"ShoneChat/apps/im/immodels"
	"ShoneChat/apps/im/ws/websocket"
	"ShoneChat/apps/social/rpc/socialclient"
	"ShoneChat/apps/task/mq/internal/config"
	"ShoneChat/pkg/constant"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"net/http"
)

type ServiceContext struct {
	config.Config

	WsClient websocket.Client
	*redis.Redis

	socialclient.Social
	immodels.ChatLogModel
	immodels.ConversationModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	svc := &ServiceContext{
		Config:            c,
		Redis:             redis.MustNewRedis(c.Redisx),
		ChatLogModel:      immodels.MustChatLogModel(c.Mongo.Url, c.Mongo.Db),
		ConversationModel: immodels.MustConversationModel(c.Mongo.Url, c.Mongo.Db),

		Social: socialclient.NewSocial(zrpc.MustNewClient(c.SocialRpc)),
	}
	token, err := svc.GetSystemToken()
	if err != nil {
		panic(err)
	}

	header := http.Header{}
	header.Set("Authorization", token)
	svc.WsClient = websocket.NewClient(c.Ws.Host, websocket.WithClientHeader(header))
	return svc
}

func (svc *ServiceContext) GetSystemToken() (string, error) {
	return svc.Redis.Get(constant.REDIS_SYSTEM_ROOT_TOKEN)
}
