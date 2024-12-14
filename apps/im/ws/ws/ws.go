package ws

import "ShoneChat/pkg/constant"

type (
	Msg struct {
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
		ConversationId    string `mapstructure:"conversationId"`
		constant.ChatType `mapstructure:"chatType"`
		SendId            string `mapstructure:"sendId"`
		RecvId            string `mapstructure:"recvId"`
		SendTime          int64  `mapstructure:"sendTime"`

		constant.MType `mapstructure:"mType"`
		Content        string `mapstructure:"content"`
	}
)
