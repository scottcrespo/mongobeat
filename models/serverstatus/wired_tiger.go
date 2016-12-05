package serverstatus

// WiredTiger ...
type WiredTiger struct {
	URI   string `bson:"uri" json:"uri"`
	LSM   `bson:"LSM" json:"lsm"`
	Async `bson:"async" json:"async"`
	BlockManager
	Cache
	Connection
	WiredTigerCursor
	DataHandle
	Log
	Reconciliation
	Session
	ThreadYield
	Transaction
	ConcurrentTransactions
}

// LSM ...
type LSM struct {
	SleepForLSMCheckpointThrottle       uint `bson:"sleep for LSM checkpoint throttle"`
	SleepForLSMMergeThrottle            uint `bson:"sleep for LSM merge throttle"`
	RowsMergedInAnLSMTree               uint `bson:"rows merged in an LSM tree"`
	ApplicationWorkUnitsCurrentlyQueued uint `bson:"application work units currently queued"`
	MergeWorkUnitsCurrentlyQueued       uint `bson:"merge work units currently queued"`
	TreeQueueHitMaximum                 uint `bson:"tree queue hit maximum"`
	SwitchWorkUnitsCurrentlyQueued      uint `bson:"switch work units currently queued"`
	TreeMaintenanceOperationsScheduled  uint `bson:"tree maintenance operations scheduled"`
	TreeMaintenanceOperationsDiscarded  uint `bson:"tree maintenance operations discarded"`
	TreeMaintenanceOperationsExecuted   uint `bson:"tree maintenance operations executed"`
}

// Async ...
type Async struct {
	NumberOfAllocationStateRaces              uint `bson:"number of allocation state races" json:"number_of_allocation_state_races,omitempty"`
	NumberOfOperationSlotsViewedForAllocation uint `bson:"number of operation slots viewed for allocation" json:"number_of_operation_slots_viewed_for_allocation,omitempty"`
	CurrentWorkQueueLength                    uint `bson:"current work queue length" json:"current_work_queue_length,omitempty"`
	NumberOfFlushCalls                        uint `bson:"number of flush calls" json:"number_of_flush_calls,omitempty"`
	NumberOfTimesOperationAllocationFailed    uint `bson:"number of times operation allocation failed" json:"number_of_times_operation_allocation_failed"`
	MaximumWorkQueueLength                    uint `bson:"maximum work queue length" json:"maximum_work_queue_length"`
	NumberOfTimesWorkerFoundNoWork            uint `bson:"number of times worker found no work" json:"number_of_times_worker_found_no_work"`
	TotalAllocations                          uint `bson:"total allocations" json:"total_allocations"`
	TotalCompactCalls                         uint `bson:"total compact calls" json:"total_compact_calls"`
	TotalInsertCalls                          uint `bson:"total insert calls" json:"total_insert_calls"`
	TotalRemoveCalls                          uint `bson:"total remove calls" json:"total_remove_calls"`
	TotalSearchCalls                          uint `bson:"total search calls" json:"total_search_calls"`
	TotalUpdateCalls                          uint `bson:"total update calls" json:"total_update_calls"`
}

// BlockManager ...
type BlockManager struct {
}

// Cache ...
type Cache struct{}

// Connection ...
type Connection struct{}

// WiredTigerCursor ...
type WiredTigerCursor struct{}

// DataHandle ...
type DataHandle struct{}

// Log ...
type Log struct{}

// Reconciliation ...
type Reconciliation struct{}

// Session ...
type Session struct{}

// ThreadYield ...
type ThreadYield struct{}

// Transaction ...
type Transaction struct{}

// ConcurrentTransactions ...
type ConcurrentTransactions struct{}
