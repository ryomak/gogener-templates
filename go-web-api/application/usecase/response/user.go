package response

import "time"

type User struct {
	ID        int64      `json:"id"`
	RandID    string     `json:"randID"`
	NickName  string     `json:"nickName"`
	Gender    string     `json:"gender"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}
