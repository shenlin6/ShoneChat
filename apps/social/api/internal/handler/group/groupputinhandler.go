package group

import (
	"net/http"

	"ShoneChat/apps/social/api/internal/logic/group"
	"ShoneChat/apps/social/api/internal/svc"
	"ShoneChat/apps/social/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GroupPutInHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GroupPutInRep
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := group.NewGroupPutInLogic(r.Context(), svcCtx)
		resp, err := l.GroupPutIn(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
