package db

import (
	"github.com/rhinoman/couchdb-go"
	"time"
	"github.com/koin-bet/koin.backend/pkg/util"
	"os"
	"log"
	"github.com/koin-bet/koin.backend/pkg/supervisor"
	"github.com/kataras/iris/core/errors"
	"fmt"
)

var (
	l       = log.New(os.Stdout, "[DB] ", 0)
	conn    *couchdb.Connection
	timeout = time.Duration(1 * time.Second)
	auth    = &couchdb.BasicAuth{os.Getenv("db_username"), os.Getenv("db_password")}

	dbUser  *couchdb.Database = nil
	dbStats *couchdb.Database = nil
)

func init() {
	var err error
	conn, err = couchdb.NewConnection(
		util.GetEnvOrDefault("db_host", "127.0.0.1"),
		util.GetEnvOrDefaultInt("db_port", 5984),
		timeout)

	if err != nil {
		panic("Can't connect to database " + err.Error() + ".")
	}

	initializeDatabase("users")
	initializeDatabase("stats")

}

func InsertStats(stats interface{}, id string) {
	_, err := dbStats.Save(stats, id, "")
	if err != nil {
		l.Printf("Error on inserting stats %s: %s.\n", id, err.Error())
	}
}

func GetStats(id string, stats interface{}) (rev string, err error) {
	fmt.Println(dbStats)
	return dbStats.Read(id, stats, nil)
}

func UpdateStats(id string, stats interface{}, rev string) {
	_, err := dbStats.Save(stats, id, rev)
	if err != nil {
		l.Printf("Error on updating stats %s with revision %s: %s.\n", id, err.Error(), rev)
	}
}

func InsertUser(user interface{}, id string) {
	_, err := dbUser.Save(user, id, "")
	if err != nil {
		l.Printf("Error on inserting user %s: %s.\n", id, err.Error())
	}
}

func GetUser(id string, user interface{}) (rev string, err error) {
	return dbUser.Read(id, user, nil)
}

func UpdateUser(id string, user interface{}, rev string) {
	_, err := dbUser.Save(user, id, rev)
	if err != nil {
		l.Printf("Error on updating user %s with revision %s: %s.\n", id, err.Error(), rev)
	}
}

func initializeDatabase(name string) {
	supervisor := supervisor.New(20, 2*time.Second)
	success := supervisor.GoSync("Initializing satabase "+name, func() error {
		db := conn.SelectDB(name, auth)
		if err := db.DbExists(); err != nil {
			if err = conn.CreateDB(name, auth); err != nil {
				return errors.New("Cant't create database " + name + ": " + err.Error() + " (" + auth.DebugString() + ").")
			}
			setDatabase(name, db)
		}
		l.Printf("Database %s initialized.\n", name)
		setDatabase(name, db)
		return nil
	})
	if !success {
		l.Panic("Cant't connect to database...")
	}

}
func setDatabase(s string, database *couchdb.Database) {
	if s == "users" {
		dbUser = database
	} else if s == "stats" {
		dbStats = database
	}
}
