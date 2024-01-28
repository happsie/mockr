package architecture

import (
	"github.com/jmoiron/sqlx"
	"log/slog"
)

type Container struct {
	DB     *sqlx.DB
	Config Config
	L      *slog.Logger
}
