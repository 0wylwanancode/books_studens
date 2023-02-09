package models

import "testing"

func TestUpdateUser(t *testing.T) {
	user := User{
		ID:         0,
		Username:   "王云龙",
		Password:   "tt123456",
		Eamil:      "a2465381@qq.com",
		Name:       "小王",
		CoverPic:   "头像1",
		ProfilePic: "www.wyl.con",
		City:       "",
		WebSite:    "",
	}
	UpdateUser(user)
}
