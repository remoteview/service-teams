package common

import (
	"fmt"

	"github.com/gobuffalo/pop"
	"github.com/pkg/errors"
)

// Database -
type Database struct {
	*pop.Connection
}

// DB -
var DB *pop.Connection

// Init - Opening a database and save the reference to `Database` struct.
func Init(goEnv string) *pop.Connection {
	db, err := pop.Connect(goEnv)
	if err != nil {
		fmt.Println("db err: ", err)
	}
	DB = db
	return DB
}

// TestDBInit creates a temporarily database for running testing cases
func TestDBInit() *pop.Connection {
	testDb, err := pop.Connect("test")
	if err != nil {
		fmt.Println("db err: ", err)
	}
	DB = testDb
	return DB
}

// GetDB - Using this function to get a connection, you can create your connection pool here.
func GetDB() *pop.Connection {
	return DB
}

// DbMigrate - Run migrations
func DbMigrate(db *pop.Connection) error {
	mig, err := pop.NewFileMigrator("./migrations", db)
	if err != nil {
		return errors.WithStack(err)
	}
	return mig.Up()
}
