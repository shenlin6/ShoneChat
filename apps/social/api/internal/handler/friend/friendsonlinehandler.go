package friend

import (
	"net/http"

	"ShoneChat/apps/social/api/internal/logic/friend"
	"ShoneChat/apps/social/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 好友在线情况
func FriendsOnlineHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := friend.NewFriendsOnlineLogic(r.Context(), svcCtx)
		resp, err := l.FriendsOnline()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
