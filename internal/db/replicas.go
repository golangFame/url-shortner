package db

import (
	"database/sql"
	"github.com/goferHiro/url-shortner/enums"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strings"
	"sync"
)

func getReplicas(conf *viper.Viper, driver string) (rdbs []*sql.DB) {

	replicaUrls := conf.GetString(enums.POSTGRESQL_REPLICAS)

	if replicaUrls == "" {
		return
	}

	_replicas := strings.Split(replicaUrls, "||")

	var wg sync.WaitGroup

	for i, replicaUrl := range _replicas {

		wg.Add(1)
		go func(i int, replicaUrl string) {
			defer wg.Done()

			if sqlDB, err := sql.Open(driver, replicaUrl); err == nil {

				log.Infoln("opened connection to the replica")

				err = sqlDB.Ping()

				if err != nil {
					log.Errorln("ignoring replica-%d", i)
				} else {
					rdbs = append(rdbs, sqlDB)
				}
				log.Infoln("pinged the replica")
			}

		}(i, replicaUrl)
	}
	wg.Wait()
	return
}
