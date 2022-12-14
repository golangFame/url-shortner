package cron

import (
	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
)

type Services interface {
	Enque(fName string, secondsInFuture int64, arg work.Q) (err error)

	Redis() (db *redis.Pool)
}

func (s *service) Enque(fName string, secondsInFuture int64, arg work.Q) (err error) {
	_, err = s.enqueuer.EnqueueIn(fName, secondsInFuture, arg)
	return
}

func (s *service) Redis() *redis.Pool {
	return s.redisPool
}
