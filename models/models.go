package models

import (
	"errors"
	"fmt"
	"github.com/upper/db/v4"
	"strings"
)

var (
	ErrNoMoreRows     = errors.New("no record found")
	ErrDuplicateEmail = errors.New("email already exists")
	ErrUserNotActive  = errors.New("your account is inactive")
	ErrInvalidLogin   = errors.New("invalid login")
)

type Models struct {
	Users UsersModel
	Posts PostModel
}

func New(db db.Session) Models {
	return Models{
		Users: UsersModel{db: db},
		Posts: PostModel{db: db},
	}
}

func convertUpperIdToInt(id db.ID) int {
	idType := fmt.Sprintf("%T", id)
	if idType == "int64" {
		return int(id.(int64))
	}
	return id.(int)
}

func errHasDuplicate(err error, key string) bool {
	str := fmt.Sprintf(`ERROR: duplicate key value violates unique constraint "%s"`, key)
	return strings.Contains(err.Error(), str)
}
