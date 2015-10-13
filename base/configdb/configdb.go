// Package configdb dinamically loads database settings
package configdb

import (
	"fmt"
	"appworkz/base/config"
	"upper.io/db"
	_ "upper.io/db/mysql"
	// _ "upper.io/db/postgresql"
)

type Settings struct {
	Database *db.Settings
}

var (
	settings *Settings
	sess     db.Database
	logFmt   = `ERROR: %s, make sure you load config on revel init`
	Load     = config.Load
)

func init() {
	initDb()
}

func initDb() {
	var err error

	settings = &Settings{}

	err = config.Load(settings)
	if err != nil {
		fmt.Printf(logFmt, err.Error())
		return
	}

	sess, err = db.Open(`mysql`, *settings.Database)
	if err != nil {
		fmt.Printf(logFmt, err.Error())
		return
	}
}

func LoadConf() {
	if settings.Database != nil {
		return
	}

	initDb()
}

func Sess() db.Database {
	return sess
}

func Collection(name string) (col db.Collection, err error) {
	col, err = sess.Collection(name)
	if err != nil {
		return nil, fmt.Errorf("Collection %s: %s", name, err.Error())
	}
	return col, nil
}
