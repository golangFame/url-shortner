package config

import (
	"fmt"
	"github.com/goferHiro/url-shortner/enums"
	"time"
)

type argvMeta struct {
	desc       string
	defaultVal string
}

var confList = map[string]argvMeta{
	enums.APP_NAME: {
		"app name",
		"url-shortner",
	},
	enums.ENV: {
		"environment",
		enums.DEV,
	},
	enums.MODE: {
		"app run configuration",
		enums.SERVER,
	},
	enums.PORT: {
		"app listen port",
		"1422",
	},
	enums.POSTGRESQL_DB: {
		"postgresql db name",
		"url_shortner",
	},
	enums.POSTGRESQL_HOST: {
		"postgresql host",
		"localhost",
	},
	enums.POSTGRESQL_PORT: {
		"postgresql port",
		"5432",
	},
	enums.POSTGRESQL_USER: {
		"postgresql username",
		"postgres",
	},
	enums.POSTGRESQL_PASSWORD: {
		"postgresql password",
		"1",
	},
	enums.POSTGRESQL_REPLICAS: {
		"postgres replica urls separated by ||",
		"host=localhost port=5432 user=postgres password=1 dbname=url_shortner",
	},
	enums.REDIS_SERVER: {
		"redis server",
		"20.219.46.174:6379",
	},
	enums.REDIS_MASTER_PASSWORD: {
		"write auth for redis cluster",
		"Nashaath",
	},
	enums.REDIS_SLAVE_PASSWORD: {
		"read auth for redis cluster",
		"Nashaath",
	},
	enums.TIMEZONE: {
		"timezone to be used",
		"UTC",
	},
	enums.JWT_SECRET: {
		"jwt secret key",
		"Hiro",
	},
	enums.JWT_ISSUER: {
		"jwt issuer",
		"Bzinga",
	},
	enums.JWT_EXPIRY_INTERVAL: {
		"jwt token expiry",
		fmt.Sprint(time.Hour * 24 * 365),
	},
	enums.API_KEY: {
		"api key",
		"eb5cddc8-1978-4d7a-98ce-9798e183ea4e",
	},
}
