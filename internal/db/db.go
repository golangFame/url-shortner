package db

import (
	"github.com/goferHiro/url-shortner/internal/genesis"
)

type db struct {
	*genesis.Service
}

type PostgresDB struct {
	db
}
