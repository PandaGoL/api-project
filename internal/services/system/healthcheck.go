package system

import (
	"context"
)

func (s *SystemService) BDCheck() error {
	return s.db.Ping(context.Background())
}
