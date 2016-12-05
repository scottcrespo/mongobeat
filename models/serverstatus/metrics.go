package serverstatus

// Metrics is the 'metrics' key of serverStatus output
type Metrics struct {
	Cursor        `bson:"cursor" json:"cursor"`
	Document      `bson:"document" json:"document"`
	GetLastError  `bson:"getLastError" json:"get_last_error"`
	Operation     `bson:"operation" json:"operation"`
	QueryExecutor `bson:"queryExecutor" json:"query_executor"`
	Record        `bson:"record" json:"record"`
	MetricsRepl   `bson:"repl" json:"repl"`
	Storage       `bson:"storage" json:"storage"`
	TTL           `bson:"ttl" json:"ttl"`
}

type Cursor struct {
	TimedOut uint64 `bson:"timedOut"      json:"timed_out"`
	Open     struct {
		NoTimeout uint64 `bson:"noTimeout"   json:"no_timeout"`
		Pinned    uint64 `bson:"pinned"      json:"pinned"`
		Total     uint64 `bson:"total"       json:"total"`
	} `bson:"open" json:"open"`
}

type Document struct {
	Deleted  uint64 `bson:"deleted"  json:"deleted"`
	Inserted uint64 `bson:"inserted" json:"inserted"`
	Returned uint64 `bson:"returned" json:"returned"`
	Updated  uint64 `bson:"updated"  json:"updated"`
}

type GetLastError struct {
	WTime struct {
		Num         uint `bson:"num" json:"num"`
		TotalMillis uint `bson:"totalMillis" json:"total_ms"`
	} `bson:"wtime" json:"wtime"`
	WTimeouts uint64 `bson:"wtimeouts" json:"write_timeouts"`
}

type Operation struct {
	FastMod        uint64 `bson:"fastmod" json:"fast_mod"`
	IDHack         uint64 `bson:"idhack" json:"id_hack"`
	ScanAndOrder   uint64 `bson:"scanAndOrder" json:"scan_and_order"`
	WriteConflicts uint64 `bson:"writeConflicts" json:"write_conflicts"`
}

type QueryExecutor struct {
	Scanned        uint64 `bson:"scanned" json:"scanned"`
	ScannedObjects uint64 `bson:"scannedObjects" json:"scanned_objects"`
}

type Record struct {
	Moves uint64 `bson:"moves" json:"moves"`
}

type MetricsRepl struct {
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
}

type Storage struct {
	FreeList struct {
		Search struct {
			BucketExhausted uint64 `bson:"bucketExhausted" json:"bucket_exhausted"`
			Requests        uint64 `bson:"requests" json:"requests"`
			Scanned         uint64 `bson:"scanned" json:"scanned"`
		} `bson:"search" json:"search"`
	} `bson:"freelist" json:"free_list"`
}

type TTL struct {
	DeletedDocuments uint64 `bson:"deletedDocuments" json:"deleted_documents"`
	Passes           uint64 `bson:"passes" json:"passes"`
}
