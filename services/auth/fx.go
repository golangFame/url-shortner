package auth

import (
	"github.com/goferHiro/url-shortner/internal/cache"
	"github.com/goferHiro/url-shortner/internal/genesis"
	"github.com/hiroBzinga/bun"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		newServices,
	),
)

type in struct {
	fx.In
	*genesis.Service
	DB            *bun.DB `name:"db"`
	CacheServices cache.Services
}

type out struct {
	fx.Out

	AuthServices Services
}

func newServices(i in) (o out) {
	o = out{
		AuthServices: newApp(i.Service, i.DB, i.CacheServices),
	}
	return
}

func newApp(genesis *genesis.Service, DB *bun.DB, cacheServices cache.Services) Services {
	return &service{
		genesis,
		DB,
	}
}
