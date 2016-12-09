package beater

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"

	"github.com/scottcrespo/mongobeat/models/dbstats"
)

// monitorDbStats calls db.stats() command for each database on the instance. This function runs
// in an event loop and sends reports according to the configured period
func (bt *Mongobeat) monitorDbStats(b *beat.Beat, node *mgo.Session, ticker *time.Ticker) {

	for {
		select {

		case <-bt.done:
			return

		case <-ticker.C:
		}

		// Get a list of databases stored on the node
		dbs, err := bt.getDatabases(node)
		if err != nil {
			logp.Err("DbStats: Failed to retrieve list of databases from mongo instance")
			return
		}

		// store common event info here
		eventTime := common.Time(time.Now())
		eventType := fmt.Sprintf("%s.db_stats", b.Name)

		// for each database, call db.stats() and publish the event
		for _, dbName := range dbs {
			db := node.DB(dbName)

			results := dbstats.DbStats{}

			err := db.Run("dbStats", &results)
			if err != nil {
				logp.Err("Failed to retrieve stats for db %s", dbName)
				continue
			}

			// instantiate event
			event := common.MapStr{
				"@timestamp": eventTime,
				"type":       eventType,
				"db_stats":   results,
			}

			// fire
			bt.client.PublishEvent(event)
		}
		logp.Info("db_stats Event sent")
	}
}

// getDatabases retrieves current list of databases in the Mongo instance
func (bt *Mongobeat) getDatabases(node *mgo.Session) ([]string, error) {
	dbs, err := node.DatabaseNames()
	if err != nil {
		logp.Err("Error retrieving mongo database names from Mongo instance")
		return []string{}, err
	}
	return dbs, nil
}
