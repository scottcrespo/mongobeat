package beater

import (
	"fmt"
	"log"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/publisher"

	"gopkg.in/mgo.v2"

	"github.com/scottcrespo/mongobeat/mongo"

	"github.com/scottcrespo/mongobeat/config"
)

// Mongobeat implements the Beater interface and adds additional methods to pull data from
// MongoDB's reporting utilities
type Mongobeat struct {
	done      chan struct{}
	config    config.Config
	client    publisher.Client
	mongoConn mgo.Session
}

// New creates a new Mongobeat Instance
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	config := config.DefaultConfig
	if err := cfg.Unpack(&config); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	mongoConn := mongo.NewMongoConnection(config.MongoConnection)

	bt := &Mongobeat{
		done:      make(chan struct{}),
		config:    config,
		mongoConn: *mongoConn,
	}

	return bt, nil
}

// Run runs the program in an event loop, publishing events periodically according to
// config.Period
func (bt *Mongobeat) Run(b *beat.Beat) error {
	logp.Info("mongobeat is running! Hit CTRL-C to stop it.")

	bt.client = b.Publisher.Connect()
	ticker := time.NewTicker(bt.config.Period)
	counter := 1

	for {

		select {

		case <-bt.done:
			return nil

		case <-ticker.C:

		}

		dbStats, err := bt.getDbStats(b)
		if err != nil {
			logp.Err("Error retrieving dbStats")
		}

		event := common.MapStr{
			"@timestamp": common.Time(time.Now()),
			"type":       b.Name,
			"dbStats":    dbStats,
		}

		event.Update(event)

		log.Fatal(event)
		bt.client.PublishEvent(event)
		logp.Info("Event sent")
		counter++
	}
}

// Stop safely exits the mongobeat program
func (bt *Mongobeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
