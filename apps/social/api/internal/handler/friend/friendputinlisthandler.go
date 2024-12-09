package friend

import (
	"net/http"

	"ShoneChat/apps/social/api/internal/logic/friend"
	"ShoneChat/apps/social/api/internal/svc"
	"ShoneChat/apps/social/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FriendPutInListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FriendPutInListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := friend.NewFriendPutInListLogic(r.Context(), svcCtx)
		resp, err := l.FriendPutInList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
