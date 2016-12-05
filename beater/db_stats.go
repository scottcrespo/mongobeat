package beater

import (
	"fmt"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"

	"github.com/scottcrespo/mongobeat/models/dbstats"
)

// getDbStats calls db.stats() command and appends to a common.MapStr, with root key as
// "dbStats"
func (bt *Mongobeat) getDbStats(b *beat.Beat) {

	dbs, err := bt.getDatabases()
	if err != nil {
		logp.Err("DbStats: Failed to retrieve list of databases from mongo instance")
		return
	}

	// store common event info here
	eventTime := common.Time(time.Now())
	eventType := fmt.Sprintf("%s.db_stats", b.Name)

	for _, dbName := range dbs {
		db := bt.mongoConn.DB(dbName)

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
		logp.Info("dbStats Event sent")
	}
}

// getDatabases retrieves current list of databases in the Mongo instance
func (bt *Mongobeat) getDatabases() ([]string, error) {
	dbs, err := bt.mongoConn.DatabaseNames()
	if err != nil {
		logp.Err("Error retrieving mongo database names from Mongo instance")
		return []string{}, err
	}
	return dbs, nil
}
