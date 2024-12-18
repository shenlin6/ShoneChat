package logic

import (
	"ShoneChat/apps/im/rpc/imclient"
	"context"
	"github.com/jinzhu/copier"

	"ShoneChat/apps/im/api/internal/svc"
	"ShoneChat/apps/im/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetUpUserConversationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetUpUserConversationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetUpUserConversationLogic {
	return &SetUpUserConversationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetUpUserConversationLogic) SetUpUserConversation(req *types.SetUpUserConversationReq) (resp *types.SetUpUserConversationResp, err error) {
	data, err := l.svcCtx.SetUpUserConversation(l.ctx, &imclient.SetUpUserConversationReq{
		SendId:   req.SendId,
		RecvId:   req.RecvId,
		ChatType: req.ChatType,
	})
	if err != nil {
		return nil, err
	}

	var res types.SetUpUserConversationResp
	copier.Copy(&res, data)

	return &res, err
}
