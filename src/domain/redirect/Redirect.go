package redirect

import (
	"errors"
)

var (
	ErrRedirectNotFound   = errors.New("redirect not found")
	ErrRedirectInvalidUrl = errors.New("invalid redirect url")
)

type Redirect struct {
	Code     string
	Url      string
	CreateAt int64
}
