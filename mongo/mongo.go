package mongo

import (
	"errors"

	"github.com/elastic/beats/libbeat/logp"

	"gopkg.in/mgo.v2"
)

// NewMasterConnection returns a connection to MongoDB (mgo.Session) by dialing the mongo
// instance specified in settings. If a connection cannot be established, a Critical error is
// thrown and the program exits
func NewMasterConnection(dialInfo *mgo.DialInfo) *mgo.Session {

	mongo, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		logp.Critical("Failed to establish connection to MongDB at %s", dialInfo.Addrs)
	}
	return mongo
}

// NewNodeConnections estbalishes direct connections with a list of hosts
func NewNodeConnections(urls []string, dialInfo *mgo.DialInfo) ([]*mgo.Session, error) {

	nodes := make([]*mgo.Session, len(urls))

	for i, url := range urls {

		// make a copy
		nodeDialInfo := *dialInfo
		nodeDialInfo.Addrs = []string{
			url,
		}
		nodeDialInfo.Direct = true

		session, err := mgo.DialWithInfo(&nodeDialInfo)
		if err != nil {
			logp.Err("Error establishing direct connection to mongo node at %s", url)
			// set i back a value so we don't skip an index when adding successful connections
			i = i - 1
			continue
		}
		nodes[i] = session
	}

	if len(nodes) == 0 {
		msg := "Error establishing connection to any mongo nodes"
		logp.Err(msg)
		return []*mgo.Session{}, errors.New(msg)
	}
	return nodes, nil
}
