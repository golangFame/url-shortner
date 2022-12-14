package internal

import (
	"github.com/goferHiro/url-shortner/internal/cache"
	"github.com/goferHiro/url-shortner/internal/cron"
	"github.com/goferHiro/url-shortner/internal/db"
	"github.com/goferHiro/url-shortner/internal/genesis"
	"go.uber.org/fx"
)

var Module = fx.Options(
	cache.Module,
	db.Module,
	genesis.Module,
	cron.Module,
)
