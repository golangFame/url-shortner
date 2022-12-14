package cron

import (
	"github.com/gocraft/work"
	"github.com/goferHiro/url-shortner/internal/genesis"
	"github.com/gomodule/redigo/redis"
)

type service struct {
	*genesis.Service
	redisPool *redis.Pool
	enqueuer  *work.Enqueuer

	pool *work.WorkerPool
}
