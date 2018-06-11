package factory

import (
	"app/infrastructure/mongodb"
)

// getMongoDb return instance mongodb connection
func getMongoDb() *mongodb.Connection {
	conn := new(mongodb.Connection)
	conn.Init()

	return conn
}
