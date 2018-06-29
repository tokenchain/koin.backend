package user

import (
	"../db"
	"github.com/dchest/uniuri"
	"fmt"
	"strconv"
	"reflect"
)

// UserService implement an user abstraction.
type UserService interface {
	Save() bool
	HasEnoughCoin(coins uint64) bool
	Update() bool
}

// User contain strict minimum information about user.
// Hash is randomly generated at first connection.
type User struct {
	Hash  string `json:"hash"`
	Name  string `json:"name"`
	Mail  string `json:"mail"`
	Coins uint64 `json:"coins"`
}

// Save save to the database the structure 'u' without Hash field.
func (u User) Save() bool {
	r := reflect.ValueOf(u)
	for _, value := range []string{"Name", "Mail", "Coins"} {
		if _, e := db.GetDb().HSet("user."+u.Hash, value, fmt.Sprint(r.FieldByName(value)));  e != nil {
			fmt.Println(e.Error())
			return false
		}
	}
	return true
}

// Update update the fields of the structure (except Hash) from the database.
func (u *User) Update() bool {
	if v, success, err := db.GetDb().HGet("user."+u.Hash, "Name"); success && err == nil {
		u.Name = v
	} else {
		return false
	}
	if v, success, err := db.GetDb().HGet("user."+u.Hash, "Coins"); success && err == nil {
		if x, err := strconv.ParseUint(v, 10, 64); err == nil {
			u.Coins = x
		}
	} else {
		return false
	}
	return true
}

// HasEnoughCoin update from the db the state of the user and then check if
// the user has enough coins that 'c'.
func (u *User) HasEnoughCoin(c uint64) bool {
	if !u.Update() {
		return false
	}
	return u.Coins >= c
}

// New create a new user with a random hash and 100 coins.
func New() *User {
	return &User{uniuri.NewLen(128), "unknown", "", 100}
}

// Get retrieve an user from a hash in the database.
func Get(hash string) *User {
	u := New()
	u.Hash = hash
	u.Update()
	return u
}