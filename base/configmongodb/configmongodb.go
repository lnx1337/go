package configmongodb

import (
	// "fmt"
	"gopkg.in/mgo.v2"
	//  "gopkg.in/mgo.v2/bson"
)

var (
	session *mgo.Session
)

func init() {
	initDb()
}

func initDb() {
	var err error
	session, err = mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	// defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	return
}

func Sess() *mgo.Session {
	return session
}

func Collection(name string) (col *mgo.Collection, err error) {
	col = session.DB("inspec").C(name)
	return col, err
}
