package json

import (
	"encoding/json"
	"github.com/NandanSatheesh/go-lang-playground/url-shortener/shortener"
	"github.com/pkg/errors"
)

type Redirect struct{}

func (r Redirect) Decode(input []byte) (*shortener.Redirect, error) {
	redirect := &shortener.Redirect{}
	if err := json.Unmarshal(input, redirect); err != nil {
		return nil, errors.Wrap(err, "Error in JSON serializer.Redirect.Decode")
	}
	return redirect, nil
}

func (r Redirect) Encode(input *shortener.Redirect) ([]byte, error) {
	rawMessage, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "Error in JSON serializer.Redirrect.Encode")
	}
	return rawMessage, nil
}
