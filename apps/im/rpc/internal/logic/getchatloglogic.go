package logic

import (
	"ShoneChat/apps/im/rpc/im"
	"ShoneChat/apps/im/rpc/internal/svc"
	"ShoneChat/pkg/xerr"
	"context"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetChatLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetChatLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetChatLogLogic {
	return &GetChatLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取会话记录
func (l *GetChatLogLogic) GetChatLog(in *im.GetChatLogReq) (*im.GetChatLogResp, error) {
	// 1. 根据 id 查询
	if in.MsgId != "" {
		chatlog, err := l.svcCtx.ChatLogModel.FindOne(l.ctx, in.MsgId)
		if err != nil {
			return nil, errors.Wrapf(xerr.NewDBErr(), "find chatLog by msgId err %v, req %v", err, in.MsgId)
		}

		return &im.GetChatLogResp{
			List: []*im.ChatLog{{
				Id:             chatlog.ID.Hex(),
				ConversationId: chatlog.ConversationId,
				SendId:         chatlog.SendId,
				RecvId:         chatlog.RecvId,
				MsgType:        int32(chatlog.MsgType),
				MsgContent:     chatlog.MsgContent,
				ChatType:       int32(chatlog.ChatType),
				SendTime:       chatlog.SendTime,
				ReadRecords:    chatlog.ReadRecords,
			}},
		}, nil
	}
	// 2. 根据时间段分段查询
	data, err := l.svcCtx.ChatLogModel.ListBySendTime(l.ctx, in.ConversationId, in.StartSendTime, in.EndSendTime, in.Count)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "find chatLog by msgId err %v, req %v", err, in.MsgId)
	}

	res := make([]*im.ChatLog, 0, len(data))
	
	for _, d := range data {
		res = append(res, &im.ChatLog{
			Id:             d.ID.Hex(),
			ConversationId: d.ConversationId,
			SendId:         d.SendId,
			RecvId:         d.RecvId,
			MsgType:        int32(d.MsgType),
			MsgContent:     d.MsgContent,
			ChatType:       int32(d.ChatType),
			SendTime:       d.SendTime,
			ReadRecords:    d.ReadRecords,
		})
	}

	return &im.GetChatLogResp{
		List: res,
	}, nil
}
