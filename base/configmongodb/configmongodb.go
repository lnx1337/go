package configmongodb

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"os"
)

var (
	session *mgo.Session
)

func init() {
	initDb()
}

func initDb() {
	var err error
	MONGO_USER := os.Getenv("MONGO_USER")
	MONGO_PASS := os.Getenv("MONGO_PASS")
	MONGO_HOST := os.Getenv("MONGO_HOST")
	MONGO_PORT := os.Getenv("MONGO_PORT")

	urlConection := fmt.Sprint("mongodb://", MONGO_USER, ":", MONGO_PASS, "@", MONGO_HOST, ":", MONGO_PORT)
	session, err = mgo.Dial(urlConection)

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
