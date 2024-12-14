package constant

type MType int

const (
	TextMType MType = iota
)

type ChatType int

const (
	//群聊和私聊
	GroupChatType ChatType = iota + 1
	SingleChatType
)
