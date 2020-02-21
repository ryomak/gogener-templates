package model

import (
	"time"
)

type Room struct {
	Name    string
	Users   []*User
	Created time.Time
	Updated time.Time
}
