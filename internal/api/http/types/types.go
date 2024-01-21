package types

import "github.com/PandaGoL/api-project/internal/database/postgres/models"

type GetUserResponse struct {
	User  []*models.User
	Count int
}
