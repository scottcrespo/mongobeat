// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import (
	"time"

	"gopkg.in/mgo.v2"
)

// Config is the main configuration struct passed to Mongobeat.New() method at startup
type Config struct {
	Period          time.Duration `config:"period"`
	MongoConnection *mgo.DialInfo `config:"mongoConn"`
}

// DefaultConfig is the default configuration parameters used to run mongobeat
var DefaultConfig = Config{
	Period:          1 * time.Second,
	MongoConnection: &mgo.DialInfo{Addrs: []string{"localhost:27017"}},
}
