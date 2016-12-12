package serverstatus

import "time"

// ServerStatus represents the comprehensive serverStatus output
type ServerStatus struct {
	Ok   uint   `bson:"ok,omitempty"                   json:"ok,omitempty"`
	Host string `bson:"host,omitempty"                 json:"host,omitempty"`
	// New in 3.2
	AdvisoryHostFQDNs     []string  `bson:"advisoryHostFQDNs,omitempty"    json:"advisory_host_fqdns,omitempty"`
	Version               string    `bson:"version,omitempty"              json:"version,omitempty"`
	Process               string    `bson:"process,omitempty"              json:"process,omitempty"`
	PID                   uint64    `bson:"pid,omitempty"                  json:"pid,omitempty"`
	Uptime                uint      `bson:"uptime,omitempty"               json:"uptime_sec,omitempty"`
	UptimeMillis          uint64    `bson:"uptimeMillis,omitempty"         json:"uptime_ms,omitempty"`
	UptimeEstimate        uint      `bson:"uptimeEstimate,omitempty"       json:"uptime_estimate,omitempty"`
	LocalTime             time.Time `bson:"localTime,omitempty"            json:"local_time,omitempty"`
	Asserts               `bson:"asserts,omitempty" json:"asserts,omitempty"`
	BackgroundFlushing    `bson:"backgroundFlushing,omitempty" json:"backgroung_flushing,omitempty"`
	Connections           `bson:"connections,omitempty" json:"connections,omitempty"`
	Dur                   `bson:"dur,omitempty" json:"dur,omitempty"`
	ExtraInfo             `bson:"extra_info "json:"extra_info,omitempty"`
	GlobalLock            `bson:"globalLock,omitempty" json:"global_lock,omitempty"`
	Locks                 `bson:"locks,omitempty" json:"locks,omitempty"`
	Network               `bson:"network,omitempty" json:"network,omitempty"`
	OpsCounters           `bson:"opscounters,omitempty" json:"ops_counters,omitempty"`
	OpsCountersRepl       `bson:"opscountersRepl,omitempty" json:"ops_counters_repl,omitempty"`
	Repl                  `bson:"repl,omitempty" json:"repl,omitempty"`
	StorageEngine         `bson:"storageEngine,omitempty" json:"storage_engine,omitempty"`
	WriteBacksBacksQueued bool `bson:"writeBacksQueued,omitempty" json:"write_backs_queued,omitempty"`
	Mem                   `bson:"mem,omitempty" json:"mem,omitempty"`
	Metrics               `bson:"metrics,omitempty" json:"metrics,omitempty"`
	WiredTiger            `bson:"wiredTiger,omitempty" json:"wired_tiger,omitempty"`
}
