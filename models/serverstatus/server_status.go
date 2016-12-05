package serverstatus

import "time"

// ServerStatus represents the comprehensive serverStatus output
type ServerStatus struct {
	Ok                    uint      `bson:"ok"                   json:"ok"`
	Host                  string    `bson:"host"                 json:"host"`
	AdvisoryHostFQDNs     []string  `bson:"advisoryHostFQDNs"    json:"advisory_host_fqdns"`
	Version               string    `bson:"version"              json:"version"`
	Process               string    `bson:"process"              json:"process"`
	PID                   uint64    `bson:"pid"                  json:"pid"`
	Uptime                uint      `bson:"uptime"               json:"uptime_sec"`
	UptimeMillis          uint64    `bson:"uptimeMillis"         json:"uptime_ms"`
	UptimeEstimate        uint      `bson:"uptimeEstimate"       json:"uptime_estimate"`
	LocalTime             time.Time `bson:"localTime"            json:"local_time"`
	Asserts               `bson:"asserts" json:"asserts"`
	BackgroundFlushing    `bson:"backgroundFlushing" json:"backgroung_flushing"`
	Connections           `bson:"connections" json:"connections"`
	Dur                   `bson:"dur" json:"dur"`
	ExtraInfo             `bson:"extra_info "json:"extra_info"`
	GlobalLock            `bson:"globalLock" json:"global_lock"`
	Locks                 `bson:"locks" json:"locks"`
	Network               `bson:"network" json:"network"`
	OpsCounters           `bson:"opscounters" json:"ops_counters"`
	OpsCountersRepl       `bson:"opscountersRepl" json:"ops_counters_repl"`
	StorageEngine         `bson:"storageEngine" json:"storage_engine"`
	WriteBacksBacksQueued bool `bson:"writeBacksQueued" json:"write_backs_queued"`
	Mem                   `bson:"mem" json:"mem"`
	Metrics               `bson:"metrics" json:"metrics"`
	WiredTiger            `bson:"wiredTiger" json:"wired_tiger"`
}
