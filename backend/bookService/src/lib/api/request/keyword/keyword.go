package keyword

import (
	"errors"
	"net/http"
)

type Keyword struct {
	Keyword string `json:"keyword"`
	Limit   int    `json:"limit"`
}

func (k *Keyword) Bind(r *http.Request) error {
	if k.Keyword == "" {
		return errors.New("keyword required")
	}
	return nil
}
