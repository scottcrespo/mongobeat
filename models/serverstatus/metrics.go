package serverstatus

// Metrics is the 'metrics' key of serverStatus output
type Metrics struct {
	Cursor        `bson:"cursor,omitempty" json:"cursor,omitempty"`
	Document      `bson:"document,omitempty" json:"document,omitempty"`
	GetLastError  `bson:"getLastError,omitempty" json:"get_last_error,omitempty"`
	Operation     `bson:"operation,omitempty" json:"operation,omitempty"`
	QueryExecutor `bson:"queryExecutor,omitempty" json:"query_executor,omitempty"`
	Record        `bson:"record,omitempty" json:"record,omitempty"`
	MetricsRepl   `bson:"repl,omitempty" json:"repl,omitempty"`
	Storage       `bson:"storage,omitempty" json:"storage,omitempty"`
	TTL           `bson:"ttl,omitempty" json:"ttl,omitempty"`
}

// Cursor ...
type Cursor struct {
	TimedOut uint64 `bson:"timedOut,omitempty"      json:"timed_out,omitempty"`
	Open     struct {
		NoTimeout uint64 `bson:"noTimeout,omitempty"   json:"no_timeout,omitempty"`
		Pinned    uint64 `bson:"pinned,omitempty"      json:"pinned,omitempty"`
		Total     uint64 `bson:"total,omitempty"       json:"total,omitempty"`
	} `bson:"open,omitempty" json:"open,omitempty"`
}

// Document ...
type Document struct {
	Deleted  uint64 `bson:"deleted,omitempty"  json:"deleted,omitempty"`
	Inserted uint64 `bson:"inserted,omitempty" json:"inserted,omitempty"`
	Returned uint64 `bson:"returned,omitempty" json:"returned,omitempty"`
	Updated  uint64 `bson:"updated,omitempty"  json:"updated,omitempty"`
}

// GetLastError ...
type GetLastError struct {
	WTime struct {
		Num         uint `bson:"num,omitempty" json:"num,omitempty"`
		TotalMillis uint `bson:"totalMillis,omitempty" json:"total_ms,omitempty"`
	} `bson:"wtime,omitempty" json:"wtime,omitempty"`
	WTimeouts uint64 `bson:"wtimeouts,omitempty" json:"write_timeouts,omitempty"`
}

// Operation ...
type Operation struct {
	FastMod        uint64 `bson:"fastmod,omitempty" json:"fast_mod,omitempty"`
	IDHack         uint64 `bson:"idhack,omitempty" json:"id_hack,omitempty"`
	ScanAndOrder   uint64 `bson:"scanAndOrder,omitempty" json:"scan_and_order,omitempty"`
	WriteConflicts uint64 `bson:"writeConflicts,omitempty" json:"write_conflicts,omitempty"`
}

// QueryExecutor ...
type QueryExecutor struct {
	Scanned        uint64 `bson:"scanned,omitempty" json:"scanned,omitempty"`
	ScannedObjects uint64 `bson:"scannedObjects,omitempty" json:"scanned_objects,omitempty"`
}

// Record ...
type Record struct {
	Moves uint64 `bson:"moves,omitempty" json:"moves,omitempty"`
}

// MetricsRepl ...
type MetricsRepl struct {
	Executor struct {
		Counters struct {
			EventCreated       uint `bson:"eventCreated,omitempty" json:"event_created,omitempty"`
			EventWait          uint `bson:"eventWait,omitempty" json:"event_wait,omitempty"`
			Cancels            uint `bson:"waits,omitempty" json:"waits,omitempty"`
			ScheduledNetCmd    uint `bson:"scheduledNetCmd,omitempty" json:"scheduled_net_command,omitempty"`
			ScheduledDBWork    uint `bson:"scheduledDBWork,omitempty" json:"scheduled_db_work,omitempty"`
			ScheduledXclWork   uint `bson:"scheduledXclWork,omitempty" json:"scheduled_xcl_work,omitempty"`
			ScheduledWorkAt    uint `bson:"scheduledWorkAt,omitempty" json:"scheduled_work_at,omitempty"`
			ScheduledWork      uint `bson:"scheduledWork,omitempty" json:"scheduled_work,omitempty"`
			SchedulingFailures uint `bson:"schedulingFailures,omitempty" json:"scheduling_failures,omitempty"`
		} `bson:"counters,omitempty" json:"counters,omitempty"`
		Queues struct {
			NetworkInProgress   uint `bson:"networkInProgress,omitempty" json:"network_in_progress,omitempty"`
			DBWorkInProgress    uint `bson:"dbWorkInProgress,omitempty" json:"db_work_in_progress,omitempty"`
			ExclusiveInProgress uint `bson:"exclusiveInProgress,omitempty" json:"exclusive_in_progress,omitempty"`
			Sleepers            uint `bson:"sleepers,omitempty" json:"sleepers,omitempty"`
			Ready               uint `bson:"ready,omitempty" json:"ready,omitempty"`
			Free                uint `bson:"free,omitempty" json:"free,omitempty"`
		} `bson:"queues,omitempty" json:"queues,omitempty"`
		UnsignaledEvents uint   `bson:"unsignaledEvents,omitempty" json:"unsignaled_events,omitempty"`
		EventWaiters     uint   `bson:"eventWaiters,omitempty" json:"event_waiters,omitempty"`
		ShuttingDown     bool   `bson:"shuttingDown,omitempty" json:"shutting_down,omitempty"`
		NetworkInterface string `bson:"networkInterface,omitempty" json:"network_interface,omitempty"`
	} `bson:"executor,omitempty" json:"executor,omitempty"`
	Apply struct {
		Batches struct {
			Num         uint `bson:"num,omitempty" json:"num,omitempty"`
			TotalMillis uint `bson:"totalMillis,omitempty" json:"total_ms,omitempty"`
		} `bson:"batches,omitempty" json:"batches,omitempty"`
		Ops uint64 `bson:"ops,omitempty" json:"ops,omitempty"`
	} `bson:"apply,omitempty" json:"apply,omitempty"`
	Buffer struct {
		Count        uint64 `bson:"count,omitempty" json:"count,omitempty"`
		MaxSizeBytes uint   `bson:"maxSizeBytes,omitempty" json:"max_size_bytes,omitempty"`
		SizeBytes    uint64 `bson:"sizeBytes,omitempty" json:"sizeBytes,omitempty"`
	} `bson:"buffer,omitempty" json:"buffer,omitempty"`
	Network struct {
		Bytes    uint64 `bson:"bytes,omitempty" json:"bytes,omitempty"`
		GetMores struct {
			Num         uint `bson:"num,omitempty" json:"num,omitempty"`
			TotalMillis uint `bson:"totalMillis,omitempty" json:"total_ms,omitempty"`
		} `bson:"getmores,omitempty" json:"get_mores,omitempty"`
		Ops            uint64 `bson:"ops,omitempty" json:"ops,omitempty"`
		ReadersCreated uint64 `bson:"readersCreated,omitempty" json:"readers_created,omitempty"`
	} `bson:"network,omitempty" json:"network,omitempty"`
	Preload struct {
		Docs struct {
			Num         uint `bson:"num,omitempty" json:"num,omitempty"`
			TotalMillis uint `bson:"totalMillis,omitempty" json:"total_ms,omitempty"`
		} `bson:"docs,omitempty" json:"docs,omitempty"`
		Indexes struct {
			Num         uint `bson:"num,omitempty" json:"num,omitempty"`
			TotalMillis uint `bson:"totalMillis,omitempty" json:"total_ms,omitempty"`
		} `bson:"indexes,omitempty" json:"indexes,omitempty"`
	} `bson:"preload,omitempty" json:"preload,omitempty"`
}

// Storage ...
type Storage struct {
	FreeList struct {
		Search struct {
			BucketExhausted uint64 `bson:"bucketExhausted,omitempty" json:"bucket_exhausted,omitempty"`
			Requests        uint64 `bson:"requests,omitempty" json:"requests,omitempty"`
			Scanned         uint64 `bson:"scanned,omitempty" json:"scanned,omitempty"`
		} `bson:"search,omitempty" json:"search,omitempty"`
	} `bson:"freelist,omitempty" json:"free_list,omitempty"`
}

// TTL ...
type TTL struct {
	DeletedDocuments uint64 `bson:"deletedDocuments,omitempty" json:"deleted_documents,omitempty"`
	Passes           uint64 `bson:"passes,omitempty" json:"passes,omitempty"`
}
