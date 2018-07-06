package user

import (
	"github.com/koinkoin-io/koinkoin.backend/pkg/db"
	"github.com/dchest/uniuri"
	"regexp"
)

// UserService implement an user abstraction.
type UserService interface {
	Save()
	HasEnoughCoin(coins uint64) bool
	Update()
}

var NameRegexp = regexp.MustCompile("^([a-zA-Z0-9-_]{2,36})$")
var MailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// User contain strict minimum information about user.
// Hash is randomly generated at first connection.
type User struct {
	Hash  string `json:"hash"`
	Name  string `json:"name"`
	Mail  string `json:"mail"`
	Coins uint64 `json:"coins"`
}

// Save save to the database the structure 'u' without Hash field.
func (u User) Save() {
	db.SaveStructure("user."+u.Hash, u)
}

// Update update the fields of the structure (except Hash) from the database.
func (u *User) Update() {
	db.StructFromKey("user."+u.Hash, u)
}

// HasEnoughCoin update from the db the state of the user and then check if
// the user has enough coins that 'c'.
func (u *User) HasEnoughCoin(c uint64) bool {
	u.Update()
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
