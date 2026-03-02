package system

import (
	"context"

	"github.com/xh-polaris/inno_agent/biz/infra/cache"
)

func InitService(ctx context.Context, cache cache.Cmdable) *SystemService {
	SystemSVC.cache = cache
	return SystemSVC
}
