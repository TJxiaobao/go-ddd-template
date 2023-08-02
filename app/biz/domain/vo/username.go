package vo

import (
	"github.com/TJxiaobao/go-ddd-template/pkg/errno"
	"regexp"
	"strings"
)

type Username struct {
	username string
}

func NewUsername(username string) Username {
	return Username{username: strings.TrimSpace(username)}
}

func (u Username) GetUsername() string {
	return u.username
}

func (u Username) CheckFormat() (bool, error) {
	if u.username == "" {
		return false, errno.NewSimpleError(errno.ErrUsernameEmpty, nil)
	}
	if !isUsernameValid(u.username) {
		return false, errno.NewSimpleError(errno.ErrInvalidUsernameFormat, nil)
	}
	return true, nil
}

func isUsernameValid(username string) bool {
	pattern := `^[a-zA-Z0-9_-]{2,50}$`
	return regexp.MustCompile(pattern).MatchString(username)
}
