package repo

import (
	"context"

	"github.com/xh-polaris/inno_agent/biz/domain/user/dal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserRepo 定义 Domain 层使用的存储接口（对 mapper 的语义化封装）
type UserRepo interface {
	FindById(ctx context.Context, id string) (*model.User, error)
	FindByBasicUserId(ctx context.Context, basicUserId string) (*model.User, error)
	FindOrCreate(ctx context.Context, basicUserId, authType, authId string) (*model.User, bool, error)
	UpdateField(ctx context.Context, id primitive.ObjectID, update bson.M) error
}