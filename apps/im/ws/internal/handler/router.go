package handler

import (
	"ShoneChat/apps/im/ws/internal/handler/conversation"
	"ShoneChat/apps/im/ws/internal/handler/push"
	"ShoneChat/apps/im/ws/internal/handler/user"
	"ShoneChat/apps/im/ws/internal/svc"
	"ShoneChat/apps/im/ws/websocket"
)

func RegisterHandlers(srv *websocket.Server, svc *svc.ServiceContext) {
	srv.AddRoutes([]websocket.Route{
		{
			Method:  "user.online",
			Handler: user.OnLine(svc),
		},
		{
			Method:  "conversation.chat",
			Handler: conversation.Chat(svc),
		},
		{
			Method:  "conversation.markChat",
			Handler: conversation.MarkRead(svc),
		},
		{
			Method:  "push",
			Handler: push.Push(svc),
		},
	})
}
