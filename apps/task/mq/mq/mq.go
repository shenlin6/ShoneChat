package mq

import "ShoneChat/pkg/constant"

type MsgChatTransfer struct {
	ConversationId    string `json:"conversationId"`
	constant.ChatType `json:"chatType"`
	SendId            string `json:"sendId"`
	RecvId            string `json:"recvId"`
	SendTime          int64  `json:"sendTime"`

	constant.MType `json:"mType"`
	Content        string `json:"content"`
}
