package database

import (
	"errors"
	"barqi.com/user/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Username string `bson:"username" json:"username"`
	Password string `bson:"password" json:"password"`
}

type UserInformation struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username string `bson:"username" json:"username"`
}

type Error struct {
	Code int `json:"code"`
	Message string `json:"message"`
}

type Token struct {
	Token string `json: "token"`
}

type AddUser struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required" format:"password"`
}

func (a AddUser) Validate() error {
	switch {
	case len(a.Username) == 0:
		return errors.New(common.ErrUsernameEmpty)
	case len(a.Password) == 0:
		return errors.New(common.ErrPasswordEmpty)
	default:
		return nil
	}
}

type Message struct {
	Message string `json: "message"`
}