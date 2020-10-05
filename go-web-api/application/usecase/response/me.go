package response

import (
	"[[.ModName]]/domain/entityx"
)

type Me struct {
	*User
}

func NewMe(ent *entityx.Me) *Me {
	return &Me{
		User: &User{
			ID:        ent.ID,
			RandID:    ent.RandID,
			NickName:  ent.Nickname,
			Gender:    ent.Gender,
			CreatedAt: ent.CreatedAt,
			UpdatedAt: ent.UpdatedAt,
			DeletedAt: ent.DeletedAt.Ptr(),
		},
	}
}

type MeGet struct {
	Me *Me `json:"me"`
}

func NewMeGet(ent *entityx.Me) *MeGet {
	return &MeGet{
		Me: NewMe(ent),
	}
}
