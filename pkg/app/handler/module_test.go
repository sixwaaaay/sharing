package handler

import (
	"github.com/sixwaaaay/sharing/pkg/app/service"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
	"testing"
)

type Option struct {
	fx.In
	Pubs []*Handler `group:"public"`
	Opt  []*Handler `group:"option"`
	Pri  []*Handler `group:"private"`
}

func TestGroupedValues(t *testing.T) {
	fxtest.New(t, HandlerMoudle, fx.Provide(func() *service.AppContext { return nil }),
		fx.Invoke(func(option Option) {
			assert.NotEmpty(t, option.Pubs)
			assert.NotEmpty(t, option.Opt)
			assert.NotEmpty(t, option.Pri)
		}),
	)
}
