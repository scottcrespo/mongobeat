package beater

import (
	"fmt"
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
	done       chan struct{}
	config     config.Config
	client     publisher.Client
	masterConn *mgo.Session
}

// New creates a new Mongobeat Instance
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	config := config.DefaultConfig
	if err := cfg.Unpack(&config); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	masterConn := mongo.NewMasterConnection(config.ConnectionInfo)

	bt := &Mongobeat{
		done:       make(chan struct{}),
		config:     config,
		masterConn: masterConn,
	}

	return bt, nil
}

// Run runs the program in an event loop, publishing events periodically according to
// config.Period
func (bt *Mongobeat) Run(b *beat.Beat) error {
	logp.Info("mongobeat is running! Hit CTRL-C to stop it.")

	bt.client = b.Publisher.Connect()
	ticker := time.NewTicker(bt.config.Period)
	// urls is a list of urls of the services discovered in the cluster
	urls := bt.masterConn.LiveServers()
	// get a list of direct connections to each of the nodes in the cluster
	nodes, err := mongo.NewNodeConnections(urls, bt.config.ConnectionInfo)
	if err != nil {
		return nil
	}

	// for each node, spawn another thread for each desired metrics
	// Having a thread per node and per metric helps safeguard against a non-responsive server,
	// or server call blocking other healthy reports
	for _, node := range nodes {

		if bt.config.DbStats {
			go bt.monitorDbStats(b, node, ticker)
		}

		if bt.config.ServerStatus {
			go bt.monitorServerStatus(b, node, ticker)
		}
	}

	// block here until the done signal is received
	for {

		select {

		case <-bt.done:
			return nil
		}
	}
}

// Stop safely exits the mongobeat program
func (bt *Mongobeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
