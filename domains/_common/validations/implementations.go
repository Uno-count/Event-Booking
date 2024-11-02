package _common

import (
	"database/sql"
)

type AppConfig struct {
	DB_INSTANCE *sql.DB
}
