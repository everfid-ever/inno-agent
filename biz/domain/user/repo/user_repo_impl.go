package repo

import (
	"context"

	"github.com/xh-polaris/inno_agent/biz/domain/user/dal/mapper"
	"github.com/xh-polaris/inno_agent/biz/domain/user/dal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var _ UserRepo = (*userRepoImpl)(nil)

type userRepoImpl struct {
	mapper mapper.MongoMapper
}

// NewUserRepo 创建 UserRepo 实例
func NewUserRepo(m mapper.MongoMapper) UserRepo {
	return &userRepoImpl{mapper: m}
}

func (r *userRepoImpl) FindById(ctx context.Context, id string) (*model.User, error) {
	return r.mapper.FindById(ctx, id)
}

func (r *userRepoImpl) FindByBasicUserId(ctx context.Context, basicUserId string) (*model.User, error) {
	return r.mapper.FindByBasicUserId(ctx, basicUserId)
}

func (r *userRepoImpl) FindOrCreate(ctx context.Context, basicUserId, authType, authId string) (*model.User, bool, error) {
	return r.mapper.FindOrCreate(ctx, basicUserId, authType, authId)
}

func (r *userRepoImpl) UpdateField(ctx context.Context, id primitive.ObjectID, update bson.M) error {
	return r.mapper.UpdateField(ctx, id, update)
}
