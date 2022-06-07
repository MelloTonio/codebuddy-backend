package owner

import (
	"time"

	"github.com/mellotonio/desafiogo/app/infra/utils"
)

// Owner Entity
type Owner struct {
	Id         string    `json:"id"`
	UserID     string    `json:"user_id"`
	Created_at time.Time `json:"created_at"`
}

func NewOwner(userID string) *Owner {
	return &Owner{
		Id:         utils.GenUUID(),
		UserID:     userID,
		Created_at: time.Now(),
	}
}
