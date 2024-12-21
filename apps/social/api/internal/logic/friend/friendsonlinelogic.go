package friend

import (
	"ShoneChat/apps/social/rpc/socialclient"
	"ShoneChat/pkg/constant"
	"ShoneChat/pkg/ctxdata"
	"context"

	"ShoneChat/apps/social/api/internal/svc"
	"ShoneChat/apps/social/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendsOnlineLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 好友在线情况
func NewFriendsOnlineLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendsOnlineLogic {
	return &FriendsOnlineLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendsOnlineLogic) FriendsOnline() (resp *types.FriendsOnlineResp, err error) {
	uid := ctxdata.GetUId(l.ctx)

	friendList, err := l.svcCtx.Social.FriendList(l.ctx, &socialclient.FriendListReq{
		UserId: uid,
	})
	if err != nil || len(friendList.List) == 0 {
		return &types.FriendsOnlineResp{}, err
	}

	// 查询，缓存中查询在线的用户
	uids := make([]string, 0, len(friendList.List))
	for _, friend := range friendList.List {
		uids = append(uids, friend.FriendUid)
	}

	onlines, err := l.svcCtx.Redis.Hgetall(constant.REDIS_ONLINE_USER)
	if err != nil {
		return nil, err
	}

	resOnLineList := make(map[string]bool, len(uids))
	for _, s := range uids {
		if _, ok := onlines[s]; ok {
			resOnLineList[s] = true
		} else {
			resOnLineList[s] = false
		}
	}

	return &types.FriendsOnlineResp{
		OnlineList: resOnLineList,
	}, nil

}
