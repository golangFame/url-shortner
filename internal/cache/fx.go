package cache

import (
	"github.com/goferHiro/url-shortner/enums"
	"github.com/goferHiro/url-shortner/internal/genesis"
	"github.com/gomodule/redigo/redis"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		newServices,
	))

type in struct {
	fx.In

	*genesis.Service
}

type out struct {
	fx.Out
	Services
}

func newServices(i in) (o out) {
	cache := &cache{
		i.Service,
	}
	o = out{
		Services: newRedis(cache),
	}
	return
}

func newRedis(cache *cache) Services {
	conf := cache.Conf
	return &redisStore{
		pool: &redis.Pool{
			MaxActive: 50,
			MaxIdle:   5,
			Wait:      true,
			Dial: func() (redis.Conn, error) {
				return redis.Dial("tcp",
					conf.GetString(enums.REDIS_SERVER),
					redis.DialPassword(conf.GetString(enums.REDIS_MASTER_PASSWORD)))
			},
		},
		cache: cache,
	}
}
