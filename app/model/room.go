package model

type Room struct {
	ID    uint64 `json:"id"`
	Users []int  `json:"users"`
}
