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
}

// Asserts is the 'asserts' key of serverStatus output
type Asserts struct {
	Regular   uint `bson:"regular"              json:"regular"`
	Warning   uint `bson:"warning"              json:"warning"`
	Msg       uint `bson:"msg"                  json:"msg"`
	User      uint `bson:"user"                 json:"user"`
	Rollovers uint `bson:"rollovers"            json:"rollovers"`
}

// BackgroundFlushing is the 'backgroundFlushing' key of serverStatus output
type BackgroundFlushing struct {
	Flushes      uint      `bson:"flushes"        json:"flushes"`
	TotalMs      uint      `bson:"total_ms"       json:"total_ms"`
	AverageMs    float64   `bson:"average_ms"     json:"avg_ms"`
	LastMs       uint      `bson:"last_ms"        json:"last_ms"`
	LastFinished time.Time `bson:"last_finished"  json:"last_finished"`
}

// Connections is the 'connections' key of serverStatus output
type Connections struct {
	Current      uint   `bson:"current"       json:"current"`
	Available    uint   `bson:"available"     json:"available"`
	TotalCreated uint64 `bson:"totalCreated"  json:"total_created"`
}

// Dur is the 'dur' key of serverStatus output
type Dur struct {
	Commmits           uint `bson:"commits"             json:"commits"`
	JournaledMB        uint `bson:"journaledMB"         json:"journaled_mb"`
	WriteToDataFilesMB uint `bson:"writeToDataFilesMB"  json:"write_to_data_files_mb"`
	Compression        uint `bson:"compression"         json:"compression:"`
	CommitsInWriteLock uint `bson:"commitsInWriteLock"  json:"commits_in_write_lock"`
	EarlyCommits       uint `bson:"earlyCommits"        json:"early_commits"`

	TimeMS struct {
		Dt                 uint `bson:"dt"                   json:"dt"`
		PrepLogBuffer      uint `bson:"prepLogBuffer"        json:"prep_log_buffer"`
		WriteToJournal     uint `bson:"writeToJournal"       json:"write_to_journal"`
		RemapPrivateView   uint `bson:"remapPrivateView"     json:"remap_private_view"`
		Commits            uint `bson:"commits"              json:"commits"`
		CommitsInWriteLock uint `bson:"commitsInWriteLock"   json:"commits_in_write_lock"`
	} `bson:"timeMs" json:"time_ms"`
}

// ExtraInfo is the 'extraInfo' key of serverStatus output
type ExtraInfo struct {
	Note       string `bson:"note"          json:"note"`
	PageFaults uint   `bson:"page_faults"   json:"page_faults"`
}

// GlobalLock is the 'globalLock' key of serverStatus output
type GlobalLock struct {
	TotalTime    uint64 `bson:"totalTime" json:"total_time"`
	CurrentQueue struct {
		Total   uint `bson:"total"   json:"total"`
		Readers uint `bson:"readers" json:"readers"`
		Writers uint `bson:"writers" json:"writers"`
	} `bson:"currentQueue" json:"current_queue"`
	ActiveClients struct {
		Total   uint `bson:"total"   json:"total"`
		Readers uint `bson:"readers" json:"readers"`
		Writers uint `bson:"writers" json:"writers"`
	} `bson:"activeClients" json:"active_clients"`
}

// Locks is the 'locks' key of serverStatus output
type Locks struct {
	Global struct {
		AcquireCount struct {
			R  uint64 `bson:"r" json:"r"`
			W  uint64 `bson:"w" json:"w"`
			WW uint64 `bson:"W" json:"W"`
		}
	}
	MMAPV1Journal struct {
		AcquireCount struct {
			R  uint64 `bson:"r" json:"r"`
			W  uint64 `bson:"w" json:"w"`
			RR uint64 `bson:"R" json:"R"`
		}
	}
	Database struct {
		AcquireCount struct {
			R  uint64 `bson:"r" json:"r"`
			RR uint64 `bson:"R" json:"R"`
			W  uint64 `bson:"W" json:"W"`
		}
	}
	Collection struct {
		AcquireCount struct {
			R uint64 `bson:"R" json:"R"`
		}
	}
	Metadata struct {
		AcquireCount struct {
			W uint64 `bson:"w" json:"w"`
			R uint64 `bson:"R" json:"R"`
		}
	}
}

// Network is the 'network' key of serverStatus output
type Network struct {
	BytesIn     uint64 `bson:"bytesIn"     json:"bytes_in"`
	BytesOut    uint64 `bson:"bytesOut"    json:"bytes_out"`
	NumRequests uint64 `bson:"numRequests" json:"num_requests"`
}

// OpsCounters is the 'opscounters' key of serverStatus output
type OpsCounters struct {
	Insert  uint `bson:"insert"  json:"insert"`
	Query   uint `bson:"query"   json:"query"`
	Update  uint `bson:"update"  json:"update"`
	Delete  uint `bson:"delete"  json:"delete"`
	Getmore uint `bson:"getmore" json:"get_more"`
	Command uint `bson:"command" json:"command"`
}

// OpsCountersRepl is the 'opscountersRepl' key of serverStatus output
type OpsCountersRepl struct {
	Insert  uint `bson:"insert"  json:"insert"`
	Query   uint `bson:"query"   json:"query"`
	Update  uint `bson:"update"  json:"update"`
	Delete  uint `bson:"delete"  json:"delete"`
	Getmore uint `bson:"getmore" json:"get_more"`
	Command uint `bson:"command" json:"command"`
}

// StorageEngine is the 'storageEngine' key of serverStatus output
type StorageEngine struct {
	Name                   string `bson:"name"                       json:"name"`
	SupportsCommittedReads bool   `bson:"supportsCommittedReads"     json:"supports_committed_reads"`
	Persistent             bool   `bson:"persistent"                 json:"persistent"`
}

// Mem is the 'mem' key of serverStatus output
type Mem struct {
	Bits              uint `bson:"bits"                  json:"bits"`
	Resident          uint `bson:"resident"              json:"resident"`
	Virtual           uint `bson:"virtual"               json:"virtual"`
	Supported         bool `bson:"supported"             json:"supported"`
	Mapped            uint `bson:"mapped"                json:"mapped"`
	MappedWithJournal uint `bson:"mappedWithJournal"     json:"mapped_with_journal"`
}

// Metrics is the 'metrics' key of serverStatus output
type Metrics struct {
	Cursor struct {
		TimedOut uint64 `bson:"timedOut"      json:"timed_out"`
		Open     struct {
			NoTimeout uint64 `bson:"noTimeout"   json:"no_timeout"`
			Pinned    uint64 `bson:"pinned"      json:"pinned"`
			Total     uint64 `bson:"total"       json:"total"`
		} `bson:"open" json:"open"`
	} `bson:"cursor" json:"cursor"`
	Document struct {
		Deleted  uint64 `bson:"deleted"  json:"deleted"`
		Inserted uint64 `bson:"inserted" json:"inserted"`
		Returned uint64 `bson:"returned" json:"returned"`
		Updated  uint64 `bson:"updated"  json:"updated"`
	} `bson:"document" json:"document"`
	GetLastError struct {
		WTime struct {
			Num         uint `bson:"num" json:"num"`
			TotalMillis uint `bson:"totalMillis" json:"total_ms"`
		} `bson:"wtime" json:"wtime"`
		WTimeouts uint64 `bson:"wtimeouts" json:"write_timeouts"`
	} `bson:"getLastError" json:"get_last_error"`
	Operation struct {
		FastMod        uint64 `bson:"fastmod" json:"fast_mod"`
		IDHack         uint64 `bson:"idhack" json:"id_hack"`
		ScanAndOrder   uint64 `bson:"scanAndOrder" json:"scan_and_order"`
		WriteConflicts uint64 `bson:"writeConflicts" json:"write_conflicts"`
	} `bson:"operation" json:"operation"`
	QueryExecutor struct {
		Scanned        uint64 `bson:"scanned" json:"scanned"`
		ScannedObjects uint64 `bson:"scannedObjects" json:"scanned_objects"`
	} `bson:"queryExecutor" json:"query_executor"`
	Record struct {
		Moves uint64 `bson:"moves" json:"moves"`
	} `bson:"record" json:"record"`
	Repl struct {
		Executor struct {
			Counters struct {
				EventCreated       uint `bson:"eventCreated" json:"event_created"`
				EventWait          uint `bson:"eventWait" json:"event_wait"`
				Cancels            uint `bson:"waits" json:"waits"`
				ScheduledNetCmd    uint `bson:"scheduledNetCmd" json:"scheduled_net_command"`
				ScheduledDBWork    uint `bson:"scheduledDBWork" json:"scheduled_db_work"`
				ScheduledXclWork   uint `bson:"scheduledXclWork" json:"scheduled_xcl_work"`
				ScheduledWorkAt    uint `bson:"scheduledWorkAt" json:"scheduled_work_at"`
				ScheduledWork      uint `bson:"scheduledWork" json:"scheduled_work"`
				SchedulingFailures uint `bson:"schedulingFailures" json:"scheduling_failures"`
			} `bson:"counters" json:"counters"`
			Queues struct {
				NetworkInProgress   uint `bson:"networkInProgress" json:"network_in_progress"`
				DBWorkInProgress    uint `bson:"dbWorkInProgress" json:"db_work_in_progress"`
				ExclusiveInProgress uint `bson:"exclusiveInProgress" json:"exclusive_in_progress"`
				Sleepers            uint `bson:"sleepers" json:"sleepers"`
				Ready               uint `bson:"ready" json:"ready"`
				Free                uint `bson:"free" json:"free"`
			} `bson:"queues" json:"queues"`
			UnsignaledEvents uint   `bson:"unsignaledEvents" json:"unsignaled_events"`
			EventWaiters     uint   `bson:"eventWaiters" json:"event_waiters"`
			ShuttingDown     bool   `bson:"shuttingDown" json:"shutting_down"`
			NetworkInterface string `bson:"networkInterface" json:"network_interface"`
		} `bson:"executor" json:"executor"`
		Apply struct {
			Batches struct {
				Num         uint `bson:"num" json:"num"`
				TotalMillis uint `bson:"totalMillis" json:"total_ms"`
			} `bson:"batches" json:"batches"`
			Ops uint64 `bson:"ops" json:"ops"`
		} `bson:"apply" json:"apply"`
		Buffer struct {
			Count        uint64 `bson:"count" json:"count"`
			MaxSizeBytes uint   `bson:"maxSizeBytes" json:"max_size_bytes"`
			SizeBytes    uint64 `bson:"sizeBytes" json:"sizeBytes"`
		} `bson:"buffer" json:"buffer"`
		Network struct {
			Bytes    uint64 `bson:"bytes" json:"bytes"`
			GetMores struct {
				Num         uint `bson:"num" json:"num"`
				TotalMillis uint `bson:"totalMillis" json:"total_ms"`
			} `bson:"getmores" json:"get_mores"`
			Ops            uint64 `bson:"ops" json:"ops"`
			ReadersCreated uint64 `bson:"readersCreated" json:"readers_created"`
		} `bson:"network" json:"network"`
		Preload struct {
			Docs struct {
				Num         uint `bson:"num" json:"num"`
				TotalMillis uint `bson:"totalMillis" json:"total_ms"`
			} `bson:"docs" json:"docs"`
			Indexes struct {
				Num         uint `bson:"num" json:"num"`
				TotalMillis uint `bson:"totalMillis" json:"total_ms"`
			} `bson:"indexes" json:"indexes"`
		} `bson:"preload" json:"preload"`
	} `bson:"repl" json:"repl"`
	Storage struct {
		FreeList struct {
			Search struct {
				BucketExhausted uint64 `bson:"bucketExhausted" json:"bucket_exhausted"`
				Requests        uint64 `bson:"requests" json:"requests"`
				Scanned         uint64 `bson:"scanned" json:"scanned"`
			} `bson:"search" json:"search"`
		} `bson:"freelist" json:"free_list"`
	} `bson:"storage" json:"storage"`

	TTL struct {
		DeletedDocuments uint64 `bson:"deletedDocuments" json:"deleted_documents"`
		Passes           uint64 `bson:"passes" json:"passes"`
	} `bson:"ttl" json:"ttl"`
}
