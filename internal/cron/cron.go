package cron

import (
	"github.com/gocraft/work"
	"os"
	"os/signal"
	"time"
)

func (s *service) init() {

	s.Log.Infoln("starting cronjob")

	time.Sleep(time.Second * 5)
	s.pool.Middleware(func(c *service, job *work.Job, next work.NextMiddlewareFunc) error {
		c = &service{}
		return next()
	})

	// Map the name of jobs to handler functions
	//pool.Job("dummy", (*service).Dummy)

	s.pool.Start()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, os.Kill)
	<-signalChan

	s.Log.Infoln("ending cronjob")

	s.pool.Stop()

	s.Log.Infoln("ended cronjob")
}
