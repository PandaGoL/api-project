package system

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

type SystemService struct {
	db *pgxpool.Pool
}

func Init(db *pgxpool.Pool) *SystemService {
	ss := &SystemService{
		db: db,
	}

	return ss
}
