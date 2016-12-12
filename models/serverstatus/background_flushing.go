package serverstatus

import "time"

// BackgroundFlushing is the 'backgroundFlushing' key of serverStatus output
type BackgroundFlushing struct {
	Flushes      uint      `bson:"flushes,omitempty"        json:"flushes,omitempty"`
	TotalMs      uint      `bson:"total_ms,omitempty"       json:"total_ms,omitempty"`
	AverageMs    float64   `bson:"average_ms,omitempty"     json:"avg_ms,omitempty"`
	LastMs       uint      `bson:"last_ms,omitempty"        json:"last_ms,omitempty"`
	LastFinished time.Time `bson:"last_finished,omitempty"  json:"last_finished,omitempty"`
}
