// Code generated by goctl. DO NOT EDIT.
// Source: social.proto

package socialclient

import (
	"context"

	"ShoneChat/apps/social/rpc/social"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	FriendListReq         = social.FriendListReq
	FriendListResp        = social.FriendListResp
	FriendPutInHandleReq  = social.FriendPutInHandleReq
	FriendPutInHandleResp = social.FriendPutInHandleResp
	FriendPutInListReq    = social.FriendPutInListReq
	FriendPutInListResp   = social.FriendPutInListResp
	FriendPutInReq        = social.FriendPutInReq
	FriendPutInResp       = social.FriendPutInResp
	FriendRequests        = social.FriendRequests
	Friends               = social.Friends
	GroupCreateReq        = social.GroupCreateReq
	GroupCreateResp       = social.GroupCreateResp
	GroupListReq          = social.GroupListReq
	GroupListResp         = social.GroupListResp
	GroupMembers          = social.GroupMembers
	GroupPutInHandleReq   = social.GroupPutInHandleReq
	GroupPutInHandleResp  = social.GroupPutInHandleResp
	GroupPutinListReq     = social.GroupPutinListReq
	GroupPutinListResp    = social.GroupPutinListResp
	GroupPutinReq         = social.GroupPutinReq
	GroupPutinResp        = social.GroupPutinResp
	GroupRequests         = social.GroupRequests
	GroupUsersReq         = social.GroupUsersReq
	GroupUsersResp        = social.GroupUsersResp
	Groups                = social.Groups

	Social interface {
		FriendPutIn(ctx context.Context, in *FriendPutInReq, opts ...grpc.CallOption) (*FriendPutInResp, error)
		FriendPutInHandle(ctx context.Context, in *FriendPutInHandleReq, opts ...grpc.CallOption) (*FriendPutInHandleResp, error)
		FriendPutInList(ctx context.Context, in *FriendPutInListReq, opts ...grpc.CallOption) (*FriendPutInListResp, error)
		FriendList(ctx context.Context, in *FriendListReq, opts ...grpc.CallOption) (*FriendListResp, error)
		GroupCreate(ctx context.Context, in *GroupCreateReq, opts ...grpc.CallOption) (*GroupCreateResp, error)
		GroupPutin(ctx context.Context, in *GroupPutinReq, opts ...grpc.CallOption) (*GroupPutinResp, error)
		GroupPutinList(ctx context.Context, in *GroupPutinListReq, opts ...grpc.CallOption) (*GroupPutinListResp, error)
		GroupPutInHandle(ctx context.Context, in *GroupPutInHandleReq, opts ...grpc.CallOption) (*GroupPutInHandleResp, error)
		GroupList(ctx context.Context, in *GroupListReq, opts ...grpc.CallOption) (*GroupListResp, error)
		GroupUsers(ctx context.Context, in *GroupUsersReq, opts ...grpc.CallOption) (*GroupUsersResp, error)
	}

	defaultSocial struct {
		cli zrpc.Client
	}
)

func NewSocial(cli zrpc.Client) Social {
	return &defaultSocial{
		cli: cli,
	}
}

func (m *defaultSocial) FriendPutIn(ctx context.Context, in *FriendPutInReq, opts ...grpc.CallOption) (*FriendPutInResp, error) {
	client := social.NewSocialClient(m.cli.Conn())
	return client.FriendPutIn(ctx, in, opts...)
}

func (m *defaultSocial) FriendPutInHandle(ctx context.Context, in *FriendPutInHandleReq, opts ...grpc.CallOption) (*FriendPutInHandleResp, error) {
	client := social.NewSocialClient(m.cli.Conn())
	return client.FriendPutInHandle(ctx, in, opts...)
}

func (m *defaultSocial) FriendPutInList(ctx context.Context, in *FriendPutInListReq, opts ...grpc.CallOption) (*FriendPutInListResp, error) {
	client := social.NewSocialClient(m.cli.Conn())
	return client.FriendPutInList(ctx, in, opts...)
}

func (m *defaultSocial) FriendList(ctx context.Context, in *FriendListReq, opts ...grpc.CallOption) (*FriendListResp, error) {
	client := social.NewSocialClient(m.cli.Conn())
	return client.FriendList(ctx, in, opts...)
}

func (m *defaultSocial) GroupCreate(ctx context.Context, in *GroupCreateReq, opts ...grpc.CallOption) (*GroupCreateResp, error) {
	client := social.NewSocialClient(m.cli.Conn())
	return client.GroupCreate(ctx, in, opts...)
}

func (m *defaultSocial) GroupPutin(ctx context.Context, in *GroupPutinReq, opts ...grpc.CallOption) (*GroupPutinResp, error) {
	client := social.NewSocialClient(m.cli.Conn())
	return client.GroupPutin(ctx, in, opts...)
}

func (m *defaultSocial) GroupPutinList(ctx context.Context, in *GroupPutinListReq, opts ...grpc.CallOption) (*GroupPutinListResp, error) {
	client := social.NewSocialClient(m.cli.Conn())
	return client.GroupPutinList(ctx, in, opts...)
}

func (m *defaultSocial) GroupPutInHandle(ctx context.Context, in *GroupPutInHandleReq, opts ...grpc.CallOption) (*GroupPutInHandleResp, error) {
	client := social.NewSocialClient(m.cli.Conn())
	return client.GroupPutInHandle(ctx, in, opts...)
}

func (m *defaultSocial) GroupList(ctx context.Context, in *GroupListReq, opts ...grpc.CallOption) (*GroupListResp, error) {
	client := social.NewSocialClient(m.cli.Conn())
	return client.GroupList(ctx, in, opts...)
}

func (m *defaultSocial) GroupUsers(ctx context.Context, in *GroupUsersReq, opts ...grpc.CallOption) (*GroupUsersResp, error) {
	client := social.NewSocialClient(m.cli.Conn())
	return client.GroupUsers(ctx, in, opts...)
}
