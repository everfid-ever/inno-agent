package application

import (
	"context"

	"github.com/xh-polaris/inno_agent/biz/application/base/appinfra"
	"github.com/xh-polaris/inno_agent/biz/application/system"
)

type BasicService struct {
	infra     *appinfra.AppDependencies
	systemSVC *system.SystemService
}

func InitApplication(ctx context.Context) error {
	infra, err := appinfra.Init(ctx)
	if err != nil {
		return err
	}
	initBasicServices(ctx, infra)
	return nil
}

func initBasicServices(ctx context.Context, infra *appinfra.AppDependencies) *BasicService {
	systemSVC := system.InitService(ctx, infra.CacheCli)
	return &BasicService{infra: infra, systemSVC: systemSVC}
}
