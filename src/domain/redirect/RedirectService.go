package redirect

import (
	"time"

	"github.com/teris-io/shortid"
)

type RedirectService struct {
	repository RedirectRepository
}

func NewRedirectService(redirectRepository RedirectRepository) *RedirectService {
	return &RedirectService{
		repository: redirectRepository,
	}
}

func (r *RedirectService) Find(code string) (*Redirect, error) {
	redirect, err := r.repository.Find(code)
	if err != nil {
		return nil, err
	} else if redirect == nil {
		return nil, ErrRedirectNotFound
	}

	return redirect, err
}

func (r *RedirectService) Store(url string) (*Redirect, error) {
	if url == "" {
		return nil, ErrRedirectInvalidUrl
	}

	redirect := Redirect{
		Url: url,
		Code: shortid.MustGenerate(),
		CreateAt:  time.Now().UTC().Unix(),
	}

	return r.repository.Store(&redirect)
}
