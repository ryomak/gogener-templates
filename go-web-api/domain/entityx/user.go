package entityx

import (
	"[[.ModName]]/domain/entity"
)

type Me struct {
	*entity.User
	*FirebaseUser
}

type FirebaseUser struct {
	Mail  *string
	Phone *string
}
