package mq

import "ShoneChat/pkg/constant"

type MsgChatTransfer struct {
	ConversationId    string `json:"conversationId"`
	constant.ChatType `json:"chatType"`
	SendId            string   `json:"sendId"`
	RecvId            string   `json:"recvId"`
	RecvIds           []string `json:"recvIds"`
	SendTime          int64    `json:"sendTime"`

	constant.MType `json:"mType"`
	Content        string `json:"content"`
}

type MsgMarkRead struct {
	constant.ChatType `json:"chatType"`
	ConversationId    string   `json:"conversationId"`
	SendId            string   `json:"sendId"`
	RecvId            string   `json:"recvId"`
	MsgIds            []string `json:"msgIds"`
}
