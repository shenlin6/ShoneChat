package user

import (
	"ShoneChat/apps/user/rpc/user"
	"context"
	"fmt"
	"github.com/jinzhu/copier"

	"ShoneChat/apps/user/api/internal/svc"
	"ShoneChat/apps/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户登入
func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	loginResp, err := l.svcCtx.User.Login(l.ctx, &user.LoginReq{
		Phone:    req.Phone,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	fmt.Println("login")
	var res types.LoginResp
	copier.Copy(&res, loginResp)

	return &res, nil
}
