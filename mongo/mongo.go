package mongo

import (
	"github.com/elastic/beats/libbeat/logp"

	"gopkg.in/mgo.v2"
)

// NewMongoConnection returns a connection to MongoDB (mgo.Session) by dialing the mongo
// instance specified in settings. If a connection cannot be established, a Critical error is
// thrown and the program exits
func NewMongoConnection(dialInfo *mgo.DialInfo) *mgo.Session {

	mongo, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		logp.Critical("Failed to establish connection to MongDB at %s", dialInfo.Addrs)
	}
	return mongo
}
