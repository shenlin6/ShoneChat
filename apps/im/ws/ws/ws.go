package ws

import "ShoneChat/pkg/constant"

type (
	Msg struct {
		MsgId          string            `mapstructure:"msgId"`
		ReadRecords    map[string]string `mapstructure:"readRecords"`
		constant.MType `mapstructure:"mType"`
		Content        string `mapstructure:"content"`
	}

	Chat struct {
		ConversationId    string `mapstructure:"conversationId"`
		constant.ChatType `mapstructure:"chatType"`
		SendId            string `mapstructure:"sendId"`
		RecvId            string `mapstructure:"recvId"`
		SendTime          int64  `mapstructure:"sendTime"`
		Msg               `mapstructure:"msg"`
	}

	Push struct {
		constant.MType    `mapstructure:"mType"`
		constant.ChatType `mapstructure:"chatType"`

		ConversationId string `mapstructure:"conversationId"`

		SendId   string   `mapstructure:"sendId"`
		RecvId   string   `mapstructure:"recvId"`
		RecvIds  []string `mapstructure:"recvIds"`
		SendTime int64    `mapstructure:"sendTime"`

		MsgId       string               `mapstructure:"msgId"`
		ReadRecords map[string]string    `mapstructure:"readRecords"`
		ContentType constant.ContentType `mapstructure:"contentType"`

		Content string `mapstructure:"content"`
	}

	MarkRead struct {
		constant.ChatType `mapstructure:"chatType"`
		RecvId            string   `mapstructure:"recvId"`
		ConversationId    string   `mapstructure:"conversationId"`
		MsgIds            []string `mapstructure:"msgIds"`
	}
)
