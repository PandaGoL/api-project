package types

import (
	"fmt"
	"time"
)

// Options - настройки БД
type Options struct {
	QueryTimeout          time.Duration
	MaxOpenConnections    int
	MaxIdleConnections    int
	MaxConnectionLifetime time.Duration
	MigrationVersion      int64
	MigrationEnable       bool
	Host                  string
	Login                 string
	Password              string
	Database              string
}

func (o Options) DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", o.Login, o.Password, o.Host, "5432", o.Database)
}
