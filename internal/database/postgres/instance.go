package postgres

import (
	"context"
	"errors"

	"github.com/PandaGoL/api-project/internal/database/postgres/types"
	"github.com/PandaGoL/api-project/internal/metrics"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

const ModuleName = "Postgres"

type PgStorage struct {
	conf    *pgxpool.Config
	pool    *pgxpool.Pool
	options *types.Options
	metrics metrics.IDatabaseMetrics
}

func New(opt *types.Options, m metrics.IDatabaseMetrics) (pgs *PgStorage, err error) {
	logrus.Info("Start connect to db")
	if opt.MaxOpenConnections < 1 {
		return nil, errors.New("MaxOpenConnections < 1")
	}

	pgs = &PgStorage{
		options: opt,
		metrics: m,
	}

	pgs.conf, err = pgxpool.ParseConfig(pgs.options.DSN())
	if err != nil {
		return nil, err
	}

	pgs.conf.MaxConns = int32(opt.MaxOpenConnections)
	pgs.conf.MaxConnIdleTime = opt.MaxConnectionLifetime
	pgs.conf.MaxConnLifetime = opt.MaxConnectionLifetime

	if pgs.pool, err = pgxpool.ConnectConfig(context.Background(), pgs.conf); err != nil {
		return
	}

	return
}

func (pgs *PgStorage) Close() {
	logrus.Warn("Shutting down database")
	pgs.pool.Close()
}
