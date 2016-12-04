package beater

import (
	"fmt"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
)

// DbStats represents the fields returned from a call to db.stats() in Mongo
type DbStats struct {
	Db          string `bson:"db"            json:"db"`
	Collections uint   `bson:"collections"   json:"collections"`
	Objects     uint64 `bson:"objects"       json:"objects"`
	AvgObjSize  uint64 `bson:"avgObjectSize" json:"avg_obj_size"`
	DataSize    uint64 `bson:"dataSize"      json:"data_size"`
	StorageSize uint64 `bson:"storageSize"   json:"storage_size"`
	NumExtents  uint64 `bson:"numExtents"    json:"num_extents"`
	Indexes     uint64 `bson:"indexes"       json:"indexes"`
	IndexSize   uint64 `bson:"indexSize"     json:"index_size"`
	FileSize    uint64 `bson:"fileSize"      json:"file_size"`
	Ok          uint   `bson:"ok"            json:"ok"`
}

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

		results := DbStats{}

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
