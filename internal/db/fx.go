package db

import (
	"github.com/goferHiro/url-shortner/internal/genesis"
	"github.com/hiroBzinga/bun"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		newDBS,
	),
)

type in struct {
	fx.In
	Service *genesis.Service
}

type out struct {
	fx.Out

	DB1 *bun.DB `name:"db"`
}

func newDBS(i in) (o out) {
	database := &db{
		i.Service,
	}
	o = out{
		DB1: newPostgressDB(database),
	}
	return
}
