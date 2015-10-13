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
	session, err = mgo.Dial("mongodb://admin:cyS6Uq67Vghq@127.12.197.2:27017/")
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
