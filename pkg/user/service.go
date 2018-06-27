package user

import (
	"../coin"
	"../db"
	"github.com/dchest/uniuri"
	"fmt"
	"strconv"
)

type UserService interface {
	Save() bool
	HasEnoughtCoin(coins coin.Coins) bool
	Update() bool
}

type User struct {
	Hash  string
	Name  string
	Coins coin.Coins
}

func (u User) Save() bool {
	if b, e := db.GetDb().HSet("user."+u.Hash, "Name", u.Name); !b || e != nil {
		return false
	}
	if b, e := db.GetDb().HSet("user."+u.Hash, "Coins", fmt.Sprint(u.Coins)); !b || e != nil {
		return false
	}
	return true
}

func (u *User) Update() bool {
	if v, success, err := db.GetDb().HGet("user."+u.Hash, "Name"); success && err == nil {
		u.Name = v
	} else {
		return false
	}
	if v, success, err := db.GetDb().HGet("user."+u.Hash, "Coins"); success && err == nil {
		if x, err := strconv.ParseUint(v, 10, 64); err == nil {
			u.Coins = coin.Coins(x)
		}
	} else {
		return false
	}
	return true
}

func (u *User) HasEnoughtCoin(c coin.Coins) bool {
	u.Update()
	return u.Coins >= c
}

func New() *User {
	return &User{uniuri.NewLen(16), "", 0}
}
