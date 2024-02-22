package domain

import "time"

type User struct {
	Id       int64     `json:"id"`
	Email    string    `json:"email"`
	Nickname string    `json:"nickname"`
	Password string    `json:"password"`
	Phone    string    `json:"phone"`
	AboutMe  string    `json:"about_me"`
	Ctime    time.Time `json:"ctime"`
	Birthday time.Time `json:"birthday"`
}
