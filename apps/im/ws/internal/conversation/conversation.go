package conversation

import (
	"ShoneChat/apps/im/ws/internal/logic"
	"ShoneChat/apps/im/ws/internal/svc"
	"ShoneChat/apps/im/ws/websocket"
	"ShoneChat/apps/im/ws/ws"
	"ShoneChat/pkg/constant"
	"context"
	"github.com/mitchellh/mapstructure"
	"time"
)

func Chat(svc *svc.ServiceContext) websocket.HandlerFunc {
	return func(srv *websocket.Server, conn *websocket.Conn, msg *websocket.Message) {
		var data ws.Chat
		if err := mapstructure.Decode(msg.Data, &data); err != nil {
			srv.Send(websocket.NewErrMessage(err), conn)
			return
		}

		switch data.ChatType {
		case constant.SingleChatType:
			err := logic.NewConversation(context.Background(), srv, svc).SingleChat(&data, conn.Uid)
			if err != nil {
				srv.Send(websocket.NewErrMessage(err), conn)
				return
			}

			srv.SendByUserId(websocket.NewMessage(conn.Uid, ws.Chat{
				ConversationId: data.ConversationId,
				ChatType:       data.ChatType,
				SendId:         conn.Uid,
				RecvId:         data.RecvId,
				SendTime:       time.Now().UnixMilli(),
				Msg:            data.Msg,
			}), data.RecvId)
		}
	}
}
