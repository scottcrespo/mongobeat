package serverstatus

import "time"

// BackgroundFlushing is the 'backgroundFlushing' key of serverStatus output
type BackgroundFlushing struct {
	Flushes      uint      `bson:"flushes"        json:"flushes"`
	TotalMs      uint      `bson:"total_ms"       json:"total_ms"`
	AverageMs    float64   `bson:"average_ms"     json:"avg_ms"`
	LastMs       uint      `bson:"last_ms"        json:"last_ms"`
	LastFinished time.Time `bson:"last_finished"  json:"last_finished"`
}
