package mongodb

import (
	"os"

	mgo "gopkg.in/mgo.v2"
)

// Connection struct
type Connection struct {
	Db *mgo.Database
}

// Init method
func (c *Connection) Init() {
	c.handler()
}

// handler method
func (c *Connection) handler() {
	// for {
	session, _ := mgo.Dial(os.Getenv("MONGO_HOST"))

	c.Db = session.DB("auth")

	// }
}
