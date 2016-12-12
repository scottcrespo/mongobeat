package beater

import (
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"

	"gopkg.in/mgo.v2"

	status "github.com/scottcrespo/mongobeat/models/serverstatus"
)

// monitorServerStatus creates an event loop to monitor a particlar mongo node by calling serverStatus()
func (bt *Mongobeat) monitorServerStatus(b *beat.Beat, node *mgo.Session, ticker *time.Ticker) {

	eventType := "server_status"

	db := node.DB("admin")

	for {
		select {

		case <-bt.done:
			return

		case <-ticker.C:

		}

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

		// fire
		bt.client.PublishEvent(event)
		logp.Info("server_status Event sent")

	}
}
