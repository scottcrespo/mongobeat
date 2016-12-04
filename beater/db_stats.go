package beater

import (
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"

	"github.com/fatih/structs"
)

// DbStats represents the fields returned from a call to db.stats() in Mongo
type DbStats struct {
	Db          string `bson:"db"            json:"db"`
	Collections int    `bson:"collections"   json:"collections"`
	Objects     int64  `bson:"objects"       json:"objects"`
	AvgObjSize  int64  `bson:"avgObjectSize" json:"avgObjSize"`
	DataSize    int64  `bson:"dataSize"      json:"dataSize"`
	StorageSize int64  `bson:"storageSize"   json:"storageSize"`
	NumExtents  int64  `bson:"numExtents"    json:"numExtents"`
	Indexes     int    `bson:"indexes"       json:"indexes"`
	IndexSize   int64  `bson:"indexSize"     json:"indexSize"`
	FileSize    int64  `bson:"fileSize"      json:"fileSize"`
	Ok          int    `bson:"ok"            json:"ok"`
}

// getDbStats calls db.stats() command and appends to a common.MapStr, with root key as
// "dbStats"
func (bt *Mongobeat) getDbStats(b *beat.Beat) {

	dbs, err := bt.getDatabases()
	if err != nil {
		logp.Err("DbStats: Failed to retrieve list of databases from mongo instance")
		return
	}

	// store event time here so all dbStats events have exactly the same time stamp
	eventTime := common.Time(time.Now())

	for _, dbName := range dbs {
		db := bt.mongoConn.DB(dbName)

		results := DbStats{}

		err := db.Run("dbStats", &results)
		if err != nil {
			logp.Err("Failed to retrieve stats for db %s", dbName)
			continue
		}

		// convert results to map[string]interface{}
		resultsMap := structs.Map(results)
		// instantiate event
		event := common.MapStr{
			"@timestamp": eventTime,
			"type":       b.Name,
		}
		// update event with db results
		event.Update(resultsMap)
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
