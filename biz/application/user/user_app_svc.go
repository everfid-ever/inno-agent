package user

import (
	"context"

	"github.com/xh-polaris/inno_agent/biz/api/model/basicuser"
	"github.com/xh-polaris/inno_agent/biz/api/model/base"
	"github.com/xh-polaris/inno_agent/biz/application/base/token"
	domainSVC "github.com/xh-polaris/inno_agent/biz/domain/user/service"
	"go.mongodb.org/mongo-driver/bson"
)

var UserAppSVC *UserApplicationService

type UserApplicationService struct {
	DomainSVC domainSVC.UserDomainSVC
	TokenSVC  *token.TokenService
}

func (s *UserApplicationService) Login(ctx context.Context, req *basicuser.LoginReq) (*basicuser.LoginResp, error) {
	u, tok, isNew, err := s.DomainSVC.Login(ctx, req.AuthType, req.AuthId, req.Verify)
	if err != nil {
		return &basicuser.LoginResp{
			Resp: &base.Response{Code: -1, Msg: err.Error()},
		}, nil
	}
	name := u.Name
	avatar := u.Avatar
	return &basicuser.LoginResp{
		Resp:   &base.Response{Code: 0, Msg: "ok"},
		Token:  &tok,
		IsNew:  &isNew,
		Name:   &name,
		Avatar: &avatar,
	}, nil
}

func (s *UserApplicationService) Register(ctx context.Context, req *basicuser.RegisterReq) (*basicuser.RegisterResp, error) {
	password := req.GetPassword()
	tok, err := s.DomainSVC.Register(ctx, req.AuthType, req.AuthId, req.Verify, password)
	if err != nil {
		return &basicuser.RegisterResp{
			Resp: &base.Response{Code: -1, Msg: err.Error()},
		}, nil
	}
	return &basicuser.RegisterResp{
		Resp:  &base.Response{Code: 0, Msg: "ok"},
		Token: &tok,
	}, nil
}

func (s *UserApplicationService) ResetPassword(ctx context.Context, req *basicuser.ResetPasswordReq, authHeader string) (*basicuser.ResetPasswordResp, error) {
	if err := s.DomainSVC.ResetPassword(ctx, authHeader, req.NewPassword); err != nil {
		return &basicuser.ResetPasswordResp{
			Resp: &base.Response{Code: -1, Msg: err.Error()},
		}, nil
	}
	return &basicuser.ResetPasswordResp{
		Resp: &base.Response{Code: 0, Msg: "ok"},
	}, nil
}

func (s *UserApplicationService) GetProfile(ctx context.Context, req *basicuser.GetProfileReq, userId string) (*basicuser.GetProfileResp, error) {
	u, err := s.DomainSVC.GetUser(ctx, userId)
	if err != nil {
		return &basicuser.GetProfileResp{
			Resp: &base.Response{Code: -1, Msg: err.Error()},
		}, nil
	}
	name := u.Name
	avatar := u.Avatar
	return &basicuser.GetProfileResp{
		Resp:    &base.Response{Code: 0, Msg: "ok"},
		Name:    &name,
		Avatar:  &avatar,
		Profile: &basicuser.UserProfile{},
	}, nil
}

func (s *UserApplicationService) UpdateProfile(ctx context.Context, req *basicuser.UpdateProfileReq, userId string) (*basicuser.UpdateProfileResp, error) {
	u, err := s.DomainSVC.GetUser(ctx, userId)
	if err != nil {
		return &basicuser.UpdateProfileResp{
			Resp: &base.Response{Code: -1, Msg: err.Error()},
		}, nil
	}
	fields := bson.M{}
	if req.IsSetName() {
		fields["name"] = req.GetName()
	}
	if req.IsSetAvatar() {
		fields["avatar"] = req.GetAvatar()
	}
	if len(fields) > 0 {
		if err = s.DomainSVC.UpdateField(ctx, u.ObjectID(), fields); err != nil {
			return &basicuser.UpdateProfileResp{
				Resp: &base.Response{Code: -1, Msg: err.Error()},
			}, nil
		}
	}
	return &basicuser.UpdateProfileResp{
		Resp: &base.Response{Code: 0, Msg: "ok"},
	}, nil
}