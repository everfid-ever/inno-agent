package errno

import "github.com/xh-polaris/inno_agent/biz/pkg/errorx/code"

// System 100 000 000	~ 100 999 999
const (
	ErrInvalidAuthType = 100_000_000
)

func init() {
	code.Register(
		ErrInvalidAuthType,
		"the auth type {type} is invalid",
		code.WithAffectStability(false),
	)
}
