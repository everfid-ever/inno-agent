package service

import (
	"context"

	"github.com/xh-polaris/inno_agent/biz/domain/user/entity"
	"github.com/xh-polaris/inno_agent/biz/domain/user/repo"
	"github.com/xh-polaris/inno_agent/biz/infra/synapse"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var _ UserDomainSVC = (*userDomainSVCImpl)(nil)

type userDomainSVCImpl struct {
	userRepo      repo.UserRepo
	synapseClient synapse.Client
}

func NewUserDomainSVC(userRepo repo.UserRepo, synapseClient synapse.Client) UserDomainSVC {
	return &userDomainSVCImpl{
		userRepo:      userRepo,
		synapseClient: synapseClient,
	}
}

func (s *userDomainSVCImpl) Login(ctx context.Context, authType, authId, verify string) (*entity.User, string, bool, error) {
	result, err := s.synapseClient.Login(ctx, authType, authId, verify)
	if err != nil {
		return nil, "", false, err
	}
	u, isNew, err := s.userRepo.FindOrCreate(ctx, result.BasicUserId, authType, authId)
	if err != nil {
		return nil, "", false, err
	}
	return entity.NewUser(u), result.Token, isNew, nil
}

func (s *userDomainSVCImpl) Register(ctx context.Context, authType, authId, verify, password string) (string, error) {
	result, err := s.synapseClient.Register(ctx, authType, authId, verify, password)
	if err != nil {
		return "", err
	}
	return result.Token, nil
}

func (s *userDomainSVCImpl) ResetPassword(ctx context.Context, authHeader, newPassword string) error {
	return s.synapseClient.ResetPassword(ctx, authHeader, newPassword)
}

func (s *userDomainSVCImpl) GetUser(ctx context.Context, userId string) (*entity.User, error) {
	u, err := s.userRepo.FindById(ctx, userId)
	if err != nil {
		return nil, err
	}
	return entity.NewUser(u), nil
}

func (s *userDomainSVCImpl) UpdateField(ctx context.Context, id primitive.ObjectID, fields bson.M) error {
	return s.userRepo.UpdateField(ctx, id, fields)
}