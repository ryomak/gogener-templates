package model

import "time"

type User struct {
	Name       string
	RoomName   string
	Created    time.Time
	Updated    time.Time
	Coordinate *Coordinate
}
