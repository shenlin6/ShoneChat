package friend

import (
	"ShoneChat/apps/social/rpc/socialclient"
	"ShoneChat/pkg/ctxdata"
	"context"
	"github.com/jinzhu/copier"

	"ShoneChat/apps/social/api/internal/svc"
	"ShoneChat/apps/social/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendPutInListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFriendPutInListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendPutInListLogic {
	return &FriendPutInListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendPutInListLogic) FriendPutInList() (resp *types.FriendPutInListResp, err error) {
	list, err := l.svcCtx.Social.FriendPutInList(l.ctx, &socialclient.FriendPutInListReq{
		UserId: ctxdata.GetUId(l.ctx),
	})
	if err != nil {
		return nil, err
	}

	var respList []*types.FriendRequests
	copier.Copy(&respList, list.List)

	return &types.FriendPutInListResp{List: respList}, nil
}
