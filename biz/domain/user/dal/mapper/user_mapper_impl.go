package mapper

import (
	"context"
	"strings"
	"time"

	"github.com/xh-polaris/inno_agent/biz/conf"
	"github.com/xh-polaris/inno_agent/biz/domain/user/dal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoMapperImpl struct {
	coll *mongo.Collection
}

func NewMongoMapper(cli *mongo.Client) MongoMapper {
	cfg := conf.GetConfig()
	coll := cli.Database(cfg.MongoDB.Database).Collection(model.CollectionName)
	return &mongoMapperImpl{coll: coll}
}

func (m *mongoMapperImpl) FindById(ctx context.Context, id string) (*model.User, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var u model.User
	err = m.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (m *mongoMapperImpl) FindByBasicUserId(ctx context.Context, basicUserId string) (*model.User, error) {
	var u model.User
	err := m.coll.FindOne(ctx, bson.M{"basicUserId": basicUserId}).Decode(&u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (m *mongoMapperImpl) FindOrCreate(ctx context.Context, basicUserId, authType, authId string) (*model.User, bool, error) {
	now := time.Now()

	namePrefix := basicUserId
	if len(namePrefix) > 6 {
		namePrefix = namePrefix[:6]
	}

	setOnInsert := bson.M{
		"basicUserId": basicUserId,
		"name":        "用户" + namePrefix,
		"avatar":      "",
		"createdAt":   now,
	}

	setFields := bson.M{"updatedAt": now}
	if strings.HasPrefix(authType, "phone") {
		setFields["phone"] = authId
	} else if strings.HasPrefix(authType, "email") {
		setFields["email"] = authId
	} else if strings.HasPrefix(authType, "studentId") {
		setFields["studentId"] = authId
	}

	filter := bson.M{"basicUserId": basicUserId}
	update := bson.M{
		"$setOnInsert": setOnInsert,
		"$set":         setFields,
	}
	opts := options.FindOneAndUpdate().
		SetUpsert(true).
		SetReturnDocument(options.After)

	var u model.User
	err := m.coll.FindOneAndUpdate(ctx, filter, update, opts).Decode(&u)
	if err != nil {
		return nil, false, err
	}
	isNew := u.CreatedAt.Equal(u.UpdatedAt) || u.UpdatedAt.Sub(u.CreatedAt) < time.Second
	return &u, isNew, nil
}

func (m *mongoMapperImpl) UpdateField(ctx context.Context, id primitive.ObjectID, fields bson.M) error {
	fields["updatedAt"] = time.Now()
	_, err := m.coll.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": fields})
	return err
}