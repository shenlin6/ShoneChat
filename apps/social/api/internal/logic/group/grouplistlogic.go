package group

import (
	"ShoneChat/apps/social/rpc/socialclient"
	"ShoneChat/pkg/ctxdata"
	"context"
	"github.com/jinzhu/copier"

	"ShoneChat/apps/social/api/internal/svc"
	"ShoneChat/apps/social/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupListLogic {
	return &GroupListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupListLogic) GroupList() (resp *types.GroupListResp, err error) {
	uid := ctxdata.GetUId(l.ctx)
	list, err := l.svcCtx.Social.GroupList(l.ctx, &socialclient.GroupListReq{
		UserId: uid,
	})

	if err != nil {
		return nil, err
	}

	var respList []*types.Groups
	copier.Copy(&respList, list.List)

	return &types.GroupListResp{List: respList}, nil
}
