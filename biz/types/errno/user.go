package errno

import "github.com/xh-polaris/inno_agent/biz/pkg/errorx/code"

// User 200 000 000	~ 200 999 999
const (
	UnSupportAuthType = 200_000_000
)

func init() {
	code.Register(
		UnSupportAuthType,
		"the auth type {type} is not supported",
		code.WithAffectStability(false),
	)
}
