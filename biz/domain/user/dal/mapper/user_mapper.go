package mapper

import (
	"context"

	"github.com/xh-polaris/inno_agent/biz/domain/user/dal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MongoMapper 定义 User 的数据操作接口
type MongoMapper interface {
	// FindById 根据本地 ID 查询用户
	FindById(ctx context.Context, id string) (*model.User, error)
	// FindByBasicUserId 根据 Synapse basicUserId 查询用户
	FindByBasicUserId(ctx context.Context, basicUserId string) (*model.User, error)
	// FindOrCreate 根据 basicUserId 查询或创建用户，写入对应 authType 的认证字段
	FindOrCreate(ctx context.Context, basicUserId, authType, authId string) (*model.User, bool, error)
	// UpdateField 按字段名批量更新
	UpdateField(ctx context.Context, id primitive.ObjectID, update bson.M) error
}