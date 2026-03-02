package system

import (
	"github.com/xh-polaris/inno_agent/biz/infra/cache"
)

var SystemSVC = &SystemService{}

type SystemService struct {
	cache cache.Cmdable
}
