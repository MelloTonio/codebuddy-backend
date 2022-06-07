package user

import (
	"time"

	"github.com/mellotonio/desafiogo/app/infra/utils"
)

// User Entity
type User struct {
	Id         string    `json:"id"`
	Nickname   string    `json:"nickname"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Created_at time.Time `json:"created_at"`
}

func NewUser(nickname string, email string, password string) *User {
	return &User{
		Id:         utils.GenUUID(),
		Nickname:   nickname,
		Email:      email,
		Password:   password,
		Created_at: time.Now(),
	}
}
