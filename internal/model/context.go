package model

type Identity struct {
	Id       int    `json:"id" dc:"id"`
	Nickname string `json:"nickname" dc:"昵称"`
	Username string `json:"username" dc:"英文名"`
	Name     string `json:"name" dc:"中文名"`
	UserId   string `json:"userId" dc:"userId"`
}
