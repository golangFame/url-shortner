package cache

import (
	"github.com/goferHiro/url-shortner/internal/genesis"
	"github.com/gomodule/redigo/redis"
	"time"
)

type cache struct {
	*genesis.Service
}

type redisStore struct {
	pool              *redis.Pool
	defaultExpiration time.Duration
	*cache
}
