package beater

import (
	"fmt"
	"os"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"

	status "github.com/scottcrespo/mongobeat/models/serverstatus"
)

// getDbStats calls db.stats() command and appends to a common.MapStr, with root key as
// "dbStats"
func (bt *Mongobeat) getServerStatus(b *beat.Beat) {

	// store common event info here
	eventType := fmt.Sprintf("%s.server_status", b.Name)

	db := bt.mongoConn.DB("admin")

	results := status.ServerStatus{}

	err := db.Run("serverStatus", &results)
	if err != nil {
		logp.Err("Failed to retrieve server status")
		return
	}

	// instantiate event
	event := common.MapStr{
		"@timestamp":    common.Time(time.Now()),
		"type":          eventType,
		"server_status": results,
	}

	fmt.Printf("%v", event)
	os.Exit(0)

	// fire
	bt.client.PublishEvent(event)
	logp.Info("dbStats Event sent")
}
