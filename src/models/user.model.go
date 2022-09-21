package models

type User struct {
	Username string `json:"username,omitempty"`
	ID int64 `json:"id,omitempty"`
	Passwd string `json:"passwd,omitempty"`
}