package beater

import (
	"fmt"
	"log"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"

	"github.com/fatih/structs"
)

// DbStats represents the fields returned from a call to db.stats() in Mongo
type DbStats struct {
	Db          string `bson:"db"            json:"db"`
	Collections int64  `bson:"collections"   json:"collections"`
	Objects     int64  `bson:"objects"       json:"objects"`
	AvgObjSize  int64  `bson:"avgObjectSize" json:"avgObjSize"`
	DataSize    int64  `bson:"dataSize"      json:"dataSize"`
	StorageSize int64  `bson:"storageSize"   json:"storageSize"`
	NumExtents  int64  `bson:"numExtents"    json:"numExtents"`
	Indexes     int64  `bson:"indexes"       json:"indexes"`
	IndexSize   int64  `bson:"indexSize"     json:"indexSize"`
	FileSize    int64  `bson:"fileSize"      json:"fileSize"`
	Ok          int    `bson:"ok"            json:"ok"`
}

// publishDbStats calls db.stats() command for each database in the mongo instance and
// publishes the results
func (bt *Mongobeat) publishDbStats(b *beat.Beat) {

	dbs, err := bt.getDatabases()
	if err != nil {
		return
	}

	for _, dbName := range dbs {
		db := bt.mongoConn.DB(dbName)

		results := DbStats{}

		err := db.Run("dbStats", &results)
		if err != nil {
			logp.Err("Failed to retrieve stats for db %s", dbName)
			continue
		}

		resultsMap := structs.Map(results)

		event := common.MapStr{
			"@timestamp": common.Time(time.Now()),
			"type":       fmt.Sprintf("%s.dbStats", b.Name),
		}

		event.Update(resultsMap)

		log.Fatal(event)
		bt.client.PublishEvent(event)
		logp.Info("Event sent")
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
