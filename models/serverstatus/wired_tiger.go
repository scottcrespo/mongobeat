package serverstatus

// WiredTiger ...
type WiredTiger struct {
	URI                    string `bson:"uri" json:"uri"`
	LSM                    `bson:"LSM" json:"lsm"`
	Async                  `bson:"async" json:"async"`
	BlockManager           `bson:"block-manager" json:"block_manager"`
	Cache                  `bson:"cache" json:"cache"`
	Connection             `bson:"connection" json:"connection"`
	WiredTigerCursor       `bson:"cursor" json:"cursor"`
	DataHandle             `bson:"data-handle" json:"data_handle"`
	Log                    `bson:"log" json:"log"`
	Reconciliation         `bson:"reconciliation" json:"reconciliation"`
	Session                `bson:"session" json:"session"`
	ThreadYield            `bson:"thread-yield" json:"thread_yield"`
	Transaction            `bson:"transaction" json:"transaction"`
	ConcurrentTransactions `bson:"concurrentTransactions" json:"concurrent_transactions"`
}

// LSM ...
type LSM struct {
	SleepForLSMCheckpointThrottle       uint `bson:"sleep for LSM checkpoint throttle" json:"sleep_for_lsm_checkpoint_throttle"`
	SleepForLSMMergeThrottle            uint `bson:"sleep for LSM merge throttle" json:"sleep_for_lsm_merge_throttle"`
	RowsMergedInAnLSMTree               uint `bson:"rows merged in an LSM tree" json:"rows_merged_in_an_lsm_tree"`
	ApplicationWorkUnitsCurrentlyQueued uint `bson:"application work units currently queued" json:"application_work_units_currently_queued"`
	MergeWorkUnitsCurrentlyQueued       uint `bson:"merge work units currently queued" json:"merge_work_units_currently_queued"`
	TreeQueueHitMaximum                 uint `bson:"tree queue hit maximum" json:"tree_queue_hit_maximum"`
	SwitchWorkUnitsCurrentlyQueued      uint `bson:"switch work units currently queued" json:"switch_work_units_currently_queued"`
	TreeMaintenanceOperationsScheduled  uint `bson:"tree maintenance operations scheduled" json:"tree_maintenance_operations_scheduled"`
	TreeMaintenanceOperationsDiscarded  uint `bson:"tree maintenance operations discarded" json:"tree_maintenance_operations_discarded"`
	TreeMaintenanceOperationsExecuted   uint `bson:"tree maintenance operations executed" json:"tree_operations_executed"`
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
	MappedBytesRead  uint `bson:"mapped bytes read" json:"mapped_bytes_read"`
	BytesRead        uint `bson:"bytes read" json:"bytes_read"`
	BytesWritten     uint `bson:"bytes written" json:"bytes_written"`
	MappedBlocksRead uint `bson:"mapped blocks read" json:"mapped_blocked_read"`
	BlocksPreLoaded  uint `bson:"blocks pre-loaded" json:"blocks_pre_loaded"`
	BlocksRead       uint `bson:"blocks read" json:"blocks_read"`
	BlocksWritten    uint `bson:"blocks written" json:"blocks_written"`
}

// Cache ...
type Cache struct {
	TrackedDirtyBytesInTheCache                         uint `bson:"tracked dirty bytes in the cache" json:"tracked_dirty_bytes_in_the_cache"`
	BytesCurrentlyInTheCache                            uint `bson:"bytes currently in the cache" json:"bytes_currently_in_the_cache"`
	MaximumBytesConfigured                              uint `bson:"maximum bytes configured" json:"maximum_bytes_configured"`
	BytesReadIntoCache                                  uint `bson:"bytes read into cache" json:"bytes_read_into_cache"`
	BytesWrittenFromCache                               uint `bson:"bytes written from cache" json:"bytes_written_from_cache"`
	PagesEvictedByApplicationThreads                    uint `bson:"pages evicted by application threads" json:"pages_evicted_by_application_threads"`
	CheckpointBlockedPageEviction                       uint `bson:"checkpoint blocked page eviction" json:"checkpoint_blocked_page_eviction"`
	UnmodifiedPagesEvicted                              uint `bson:"unmodified pages evicted" json:"unmodified_pages_evicted"`
	PageSplitDuringEvictionDeepenedTheTree              uint `bson:"page split during eviction deepened the tree" json:"page_split_during_eviction_deepened_the_tree"`
	ModifiedPagesEvicted                                uint `bson:"modified pages evicted" json:"modified_pages_evicted"`
	PagesSelectedForEvictionUnableToBeEvicted           uint `bson:"pages selected for eviction unable to be evicted" json:"pages_selected_for_eviction_unable_to_be_evicted"`
	PagesEvictedBecauseTheyExceededTheInMemoryMaximum   uint `bson:"pages evicted because they exceeded the in-memory maximum" json:"pages_evicted_because_they_exceeded_the_in_memory_maximum"`
	PagesEvictedBecauseTheyHadChainsOfDeletedItems      uint `bson:"pages evicted because they had chains of deleted items" json:"pages_evicted_because_they_had_chains_of_deleted_items"`
	FailedEvictionOfPagesThatExceededTheInMemoryMaximum uint `bson:"failed eviction of pages that exceeded the in-memory maximum" json:"failed_eviction_of_pages_that_exceeded_the_in_memory_maximum"`
	HazardPointerBlockedPageEviction                    uint `bson:"hazard pointer blocked page eviction" json:"hazard_pointer_blocked_page_eviction"`
	InternalPagesEvicted                                uint `bson:"internal pages evicted" json:"internal_pages_evicted"`
	MaximumPageSizeAtEviction                           uint `bson:"maximum page size at eviction" json:"max_page_size_at_eviction"`
	EvictionServerCandidateQueueEmptyWhenToppingUp      uint `bson:"eviction server candidate queue empty when topping up" json:"eviction_server_candidate_queue_empty_when_topping_up"`
	EvictionServerCandidateQueueNotEmptyWhenToppingUp   uint `bson:"eviction server candidate queue not empty when topping up" json:"eviction_server_candidate_queue_not_empty_when_topping_up"`
	EvictionServerEvictingPages                         uint `bson:"eviction server evicting pages" json:"eviction_server_evicting_pages"`
	EvictionServerPopulatingQueueButNotEvictingPages    uint `bson:"eviction server populating queue, but not evicting pages" json:"eviction_server_populating_queue_but_not_evicting_pages"`
	EvictionServerUnableToReachEvictionGoal             uint `bson:"eviction server unable to reach eviction goal" json:"eviction_server_unable_to_reach_eviction_goal"`
	PagesSplitDuringEviction                            uint `bson:"pages split during eviction" json:"pages_split_during_eviction"`
	PagesWalkedForEviction                              uint `bson:"pages walked for eviction" json:"pages_walked_for_eviction"`
	EvictionWorkerThreadEvictingPages                   uint `bson:"eviction worker thread evicting pages" json:"eviction_worker_thread_evicting_pages"`
	InMemoryPageSplits                                  uint `bson:"in-memory page splits" json:"in_memory_page_splits"`
	PercentageOverhead                                  uint `bson:"percentage overhead" json:"percentage_overhead"`
	TrackedDirtyPagesInTheCache                         uint `bson:"tracked dirty pages in the cache" json:"tracked_dirty_pages_in_the_cache"`
	PagesCurrentlyHeldInTheCache                        uint `bson:"pages currently held in the cache" json:"pages_currently_held_in_the_cache"`
	PagesReadIntoCache                                  uint `bson:"pages read into cache" json:"pages_read_into_cache"`
	PagesWrittenFromCache                               uint `bson:"pages written from cache" json:"pages_written_from_cache"`
}

// Connection ...
type Connection struct {
	PThreadMutexConditionWaitCalls       uint `bson:"pthread mutex condition wait calls" json:"pthread_mutex_condition_wait_calls"`
	FilesCurrentlyOpen                   uint `bson:"files currently open" json:"files_currently_open"`
	MemoryAllocations                    uint `bson:"memory allocations" json:"memory_allocations"`
	MemoryFrees                          uint `bson:"memory frees" json:"memory_frees"`
	MemoryReAllocations                  uint `bson:"memory re-allocations" json:"memory_re_allocations"`
	TotalReadIOs                         uint `bson:"total read I/Os" json:"total_read_IOs"`
	PThreadMutexSharedLockReadLockCalls  uint `bson:"pthread mutex shared lock read-lock calls" json:"pthread_mutex_shared_lock_read_lock_calls"`
	PThreadMutexSharedLockWriteLockCalls uint `bson:"pthread mutex shared lock write-lock calls" json:"pthread_mutex_shared_lock_write_lock_calls"`
	TotalWriteIOs                        uint `bson:"total write I/Os" json:"total_write_IOs"`
}

// WiredTigerCursor ...
type WiredTigerCursor struct {
	CursorCreateCalls     uint `bson:"cursor create calls" json:"cursor_create_calls"`
	CursorInsertCalls     uint `bson:"cursor insert calls" json:"cursor_insert_calls"`
	CursorNextCalls       uint `bson:"cursor next calls" json:"cursor_next_calls"`
	CursorPrevCalls       uint `bson:"cursor prev calls" json:"cursor_prev_calls"`
	CursorRemoveCalls     uint `bson:"cursor remove calls" json:"cursor_remove_calls"`
	CursorResetCalls      uint `bson:"cursor reset calls" json:"cursor_reset_calls"`
	CursorSearchCalls     uint `bson:"cursor search calls" json:"cursor_search_calls"`
	CursorSearchNearCalls uint `bson:"cursor search near calls" json:"cursor_search_near_calls"`
	CursorUpdateCalls     uint `bson:"cursor update calls" json:"cursor_update_calls"`
}

// DataHandle ...
type DataHandle struct {
	ConnectionDHandlesSwept       uint `bson:"connection dhandles swept" json:"connection_d_handles_swept"`
	ConnectionCandidateReferences uint `bson:"connection candidate references" json:"connection_candidate_referenced"`
	ConnectionSweeps              uint `bson:"connection sweeps" json:"connection_sweeps"`
	ConnectionsTimeOfDeathSets    uint `bson:"connection time-of-death sets" json:"connection_time_of_death_sets"`
	SessionDHandlesSwept          uint `bson:"connection dhandles swept" json:"connection_d_handles_swept"`
	SessionSweepAttempts          uint `bson:"session sweep attempts" json:"session_sweep_attempts"`
}

// Log ...
type Log struct {
	LogBufferSizeIncreases                       uint `bson:" log buffer size increases" json:"log_buff_size_increases"`
	TotalLogBufferSize                           uint `bson:"total log buffer size" json:"total_log_buffer_size"`
	LogBytesOfPayLoadData                        uint `bson:"log bytes of payload data" json:"load_bytes_of_payload_data"`
	LogBytesWritten                              uint `bson:"log bytes written" json:"log_bytes_written"`
	YieldsWaitingForPreviousLogFileClose         uint `bson:"yields waiting for previous log file close" json:"yields_waiting_for_previous_log_file_close"`
	TotalSizeOfCompressedRecords                 uint `bson:"total size of compressed records" json:"total_size_of_compressed_records"`
	TotalInMemorySizeOfCompressedRecords         uint `bson:"total in-memory size of compressed records" json:"total_in_memory_size_of_compressed_records"`
	LogRecordsTooSmallToCompress                 uint `bson:"log records too small to compress" json:"log_records_too_small_to_compress"`
	LogRecordsNotCompressed                      uint `bson:"log records not compressed" json:"log_records_not_compressed"`
	LogRecordsCompressed                         uint `bson:"log records compressed" json:"log_records_compressed"`
	MaximumLogFileSize                           uint `bson:"maximum log file size" json:"max_log_file_size"`
	PreAllocatedLogFilesPrepared                 uint `bson:"pre-allocated log files prepared" json:"pre_allocated_log_files_prepared"`
	NumberOfPreAllocatedLogFilesToCreate         uint `bson:"number of pre-allocated log files to create" json:"number_of_pre_allocated_log_files_to_create"`
	PreAllocatedLogFilesUsed                     uint `bson:"pre-allocated log files used" json:"pre_allocated_log_files_used"`
	LogReadOperations                            uint `bson:"log read operations" json:"log_read_operations"`
	LogReleaseAdvancesWriteLSN                   uint `bson:"log release advances write LSN" json:"log_release_advances_write_lsn"`
	RecordsProcessedByLogScan                    uint `bson:"records processed by log scan" json:"records_processed_by_log_scan"`
	LogScanRecordsRequiringTwoReads              uint `bson:"log scan records requiring two reads" json:"log_scan_records_requiring_two_reads"`
	LogScanOperations                            uint `bson:"log scan operations" json:"log_scan_operations"`
	ConsolidatedSlotClosures                     uint `bson:"consolidated slot closures" json:"consolidated slot closures"`
	LoggingBytesConsolidated                     uint `bson:"logging bytes consolidated" json:"logging_bytes_consolidated"`
	ConsolidatedSlotJoins                        uint `bson:"consolidated slot joins" json:"consolidated_slot_joins"`
	ConsolidatedSlotJoinsRaces                   uint `bson:"consolidated slot joins races" json:"consolidated_slot_joins_races"`
	SlotsSelectedForSwitchingThatWereUnavailable uint `bson:"slots selected for switching that were unavailable" json:"slots_selected_for_switching_that_were_unavailable"`
	RecordSizeExceededMaximum                    uint `bson:"record size exceeded maximum" json:"record_size_exceeded_maximum"`
	FailedToFindASlotLargeEnoughForRecord        uint `bson:"failed to find a slot large enough for record" json:"failed_to_find_a_slot_large_enough_for_record"`
	ConsolidatedSlotJoinTransitions              uint `bson:"consolidated slot join transitions" json:"consolidated slot join transitions"`
	LogSyncOperations                            uint `bson:"log sync operations" json:"log_sync_operations"`
	LogSyncDirOperations                         uint `bson:"log sync_dir operations" json:"log_sync_dir_operations"`
	LogWriteOperations                           uint `bson:"log write operations" json:"log_write_operations"`
}

// Reconciliation ...
type Reconciliation struct {
	PageReconciliationCalls            uint `bson:"page reconciliation calls" json:"page_reconciliation_calls"`
	PageReconciliationCallsForEviction uint `bson:"page reconciliation calls for eviction" json:"page_reconciliation_calls_for_eviction"`
	SplitBytesCurrentlyAwaitingFree    uint `bson:"split bytes currently awaiting free" json:"split_bytes_currently_awaiting_free"`
	SplitObjectsCurrentlyAwaitingFree  uint `bson:"split objects currently awaiting free" json:"split_objects_currently_awaiting_free"`
}

// Session ...
type Session struct {
	OpenCursorCount  uint `bson:"open cursor count" json:"open_cursor_count"`
	OpenSessionCount uint `bson:"open session count" json:"open_session_count"`
}

// ThreadYield ...
type ThreadYield struct {
	PageAcquireBusyBlocked     uint `bson:"page acquire busy blocked" json:"page_acquire_busy_blocked"`
	PageAcquireEvictionBlocked uint `bson:"page acquire eviction blocked" json:"page_acquire_eviction_blocked"`
	PageAcquireLockedBlocked   uint `bson:"page acquire locked blocked" json:"page_acquire_locked_blocked"`
	PageAcquireReadBlocked     uint `bson:"page acquire read blocked" json:"page_acquire_read_blocked"`
	PageAcquireTimeSleeping    uint `bson:"page acquire time sleeping" json:"page_acquire_time_sleeping"`
}

// Transaction ...
type Transaction struct {
	TransactionBegins                     uint `bson:"transaction begins" json:"transaction_begins"`
	TransactionCheckPoints                uint `bson:"transaction checkpoints" json:"tansaction_checkpoints"`
	TransactionCheckPointCurrentlyRunning uint `bson:"transaction checkpoint curently running" json:"transaction_checkpoint_currently_running"`
	TransactionCheckPointMaxTime          uint `bson:"transaction checkpoint max time (msecs)" json:"transaction_checkpoint_max_time_ms"`
	TransactionCheckPointMinTime          uint `bson:"transaction checkpoint min time (msecs)" json:"transaction_checkpoint_min_time_ms"`
	TransactionCheckPointMostRecentTime   uint `bson:"transaction checkpoint most recent time (msecs)" json:"transaction_checkpoint_most_recent_time_ms"`
	TransactionCheckPointTotalTime        uint `bson:"transaction checkpoint total time (msecs)" json:"transaction_checkpoint_total_time_ms"`
	TransactionsCommitted                 uint `bson:"transactions committed" json:"transactions_committed"`
	TransactionFailuresDueToCacheOverflow uint `bson:"transaction failures due to chache overflow" json:"transaction_failures_due_to_cache_overflow"`
	TransactionRangeOfIDsCurrentlyPinned  uint `bson:"transaction range of IDs currently pinned" json:"transaction_range_of_ids_currently_pinned"`
	TransactionsRolledBack                uint `bson:"transactions rolled back" json:"transactions_rolled_back"`
}

// ConcurrentTransactions ...
type ConcurrentTransactions struct {
	Write struct {
		Out          uint `bson:"out" json:"out"`
		Available    uint `bson:"available" json:"available"`
		TotalTickets uint `bson:"totalTickets" json:"total_tickets"`
	} `bson:"write" json:"write"`

	Read struct {
		Out          uint `bson:"out" json:"out"`
		Available    uint `bson:"available" json:"available"`
		TotalTickets uint `bson:"totalTickets" json:"total_tickets"`
	} `bson:"read" json:"read"`
}
