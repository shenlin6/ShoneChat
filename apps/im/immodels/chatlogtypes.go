package immodels

import (
	"ShoneChat/pkg/constant"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var DefaultChatLogLimit int64 = 100

type ChatLog struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`

	ConversationId string            `bson:"conversationId"`
	SendId         string            `bson:"sendId"`
	RecvId         string            `bson:"recvId"`
	MsgFrom        int               `bson:"msgFrom"`
	ChatType       constant.ChatType `bson:"chatType"`
	MsgType        constant.MType    `bson:"msgType"`
	MsgContent     string            `bson:"msgContent"`
	SendTime       int64             `bson:"sendTime"`
	Status         int               `bson:"status"`
	ReadRecords    []byte            `bson:"readRecords"`

	UpdateAt time.Time `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
