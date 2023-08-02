package vo

import (
	"github.com/TJxiaobao/go-ddd-template/pkg/errno"
	"regexp"
	"strings"
)

type Email struct {
	email string
}

func NewEmail(email string) Email {
	return Email{email: strings.TrimSpace(email)}
}

func (e Email) GetEmail() string {
	return e.email
}

func (e Email) CheckFormat() (bool, error) {
	if e.email == "" {
		return true, nil
	}
	if !isEmailValid(e.email) {
		return false, errno.NewSimpleError(errno.ErrInvalidEmailFormat, nil)
	}
	return true, nil
}

func isEmailValid(email string) bool {
	pattern := `^[0-9a-zA-Z][_.0-9a-zA-Z-]{0,31}@([0-9a-zA-Z][0-9a-zA-Z-]{0,30}[0-9a-zA-Z].){1,4}[a-z]{2,4}$`
	return regexp.MustCompile(pattern).MatchString(email)
}
