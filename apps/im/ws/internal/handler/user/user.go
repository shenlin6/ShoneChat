package user

import (
	"ShoneChat/apps/im/ws/internal/svc"
	"ShoneChat/apps/im/ws/websocket"
	websocketx "github.com/gorilla/websocket"
)

// OnLine 获取所有在线用户
func OnLine(svc *svc.ServiceContext) websocket.HandlerFunc {
	return func(srv *websocket.Server, conn *websocketx.Conn, msg *websocket.Message) {
		uids := srv.GetUsers()
		u := srv.GetUsers(conn)
		err := srv.Send(websocket.NewMessage(u[0], uids), conn)
		srv.Info("err ", err)
	}
}
