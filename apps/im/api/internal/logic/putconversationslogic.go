package logic

import (
	"ShoneChat/apps/im/rpc/im"
	"ShoneChat/apps/im/rpc/imclient"
	"ShoneChat/pkg/ctxdata"
	"context"
	"github.com/jinzhu/copier"

	"ShoneChat/apps/im/api/internal/svc"
	"ShoneChat/apps/im/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutConversationsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPutConversationsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutConversationsLogic {
	return &PutConversationsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PutConversationsLogic) PutConversations(req *types.PutConversationsReq) (resp *types.PutConversationsResp, err error) {
	uid := ctxdata.GetUId(l.ctx)
	var list map[string]*im.Conversation
	copier.Copy(&list, req.ConversationList)

	data, err := l.svcCtx.PutConversations(l.ctx, &imclient.PutConversationsReq{
		UserId:           uid,
		ConversationList: list,
	})
	if err != nil {
		return nil, err
	}

	var res types.PutConversationsResp
	copier.Copy(&res, data)

	return &res, err
}
