package request

import (
	"net/http"

	"[[.ModName]]/pkg/context"
)

type MeGet struct {
	AuthID string `json:"-"`
}

func NewMeGet(req *http.Request) (*MeGet, error) {
	r := &MeGet{}
	r.AuthID = context.AuthID(req.Context())
	return r, nil
}
