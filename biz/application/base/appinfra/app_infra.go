package appinfra

import (
	"context"

	"github.com/xh-polaris/inno_agent/biz/infra/contract/cache"
	"github.com/xh-polaris/inno_agent/biz/infra/contract/id"
	"github.com/xh-polaris/inno_agent/biz/infra/impl/cache/redis"
	"github.com/xh-polaris/inno_agent/biz/infra/impl/mongoid"
)

type AppDependencies struct {
	IDGen    id.IDGenerator
	CacheCli cache.Cmdable
}

// Init 初始化
func Init(ctx context.Context) (_ *AppDependencies, err error) {
	infra := &AppDependencies{}
	infra.IDGen = mongoid.New(ctx)
	infra.CacheCli = redis.New()
	return infra, nil
}
