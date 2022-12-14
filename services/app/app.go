package app

import (
	"fmt"
	"github.com/goferHiro/url-shortner/internal/cache"
	"github.com/goferHiro/url-shortner/internal/genesis"
	"github.com/hiroBzinga/bun"
	"time"
)

type service struct {
	*genesis.Service
	db            *bun.DB
	cacheServices cache.Services
}

const serviceKey = "app"

// val can be a ptr
func (s *service) setCache(key string, val interface{}, expire time.Duration) (err error) {
	key = fmt.Sprintf("%s:%s", serviceKey, key)

	err = s.cacheServices.Set(key, val, expire)
	if err != nil {
		s.Log.Errorln("cache failed to set -", key)
		s.Log.Debugln(err.Error())
		panic(err.Error()) //FIXME
	}
	return
}

func (s *service) getCache(key string, ptr interface{}) (miss bool) {
	key = fmt.Sprintf("%s:%s", serviceKey, key)

	err := s.cacheServices.Get(key, ptr)

	if err != nil {
		s.Log.Errorln("cache missed -", key)
		s.Log.Debugln(err.Error())
		miss = true
	} else {
		s.Log.Infoln("cache hit -", key)
	}

	return
}
