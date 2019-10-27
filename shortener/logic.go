package shortener

import (
	"errors"
	errs "github.com/pkg/errors"
	"github.com/teris-io/shortid"
	"gopkg.in/dealancer/validate.v2"
	"time"
)

var (
	ErrorRedirectNotFound = errors.New("Redirect Not Found")
	ErrorRedirectInvalid  = errors.New("Redirect Invalid")
)

type redirectService struct {
	redirectRepository RedirectRepository
}

func NewRedirectService(repository RedirectRepository) RedirectService {
	return &redirectService{repository}
}

func (r redirectService) Find(Code string) (*Redirect, error) {
	return r.redirectRepository.Find(Code)
}

func (r redirectService) Store(redirect *Redirect) error {
	if err := validate.Validate(redirect); err != nil {
		return errs.Wrap(ErrorRedirectInvalid, "Error in service.Redirect.Store")
	}
	redirect.Code = shortid.MustGenerate()
	redirect.CreatedAt = time.Now().UTC().Unix()
	return r.redirectRepository.Store(redirect)
}
