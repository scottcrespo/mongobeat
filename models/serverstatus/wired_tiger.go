package serverstatus

// WiredTiger ...
type WiredTiger struct {
	URI                    string `bson:"uri,omitempty" json:"uri,omitempty"`
	LSM                    `bson:"LSM,omitempty" json:"lsm,omitempty"`
	Async                  `bson:"async,omitempty" json:"async,omitempty"`
	BlockManager           `bson:"block-manager,omitempty" json:"block_manager,omitempty"`
	Cache                  `bson:"cache,omitempty" json:"cache,omitempty"`
	Connection             `bson:"connection,omitempty" json:"connection,omitempty"`
	WiredTigerCursor       `bson:"cursor,omitempty" json:"cursor,omitempty"`
	DataHandle             `bson:"data-handle,omitempty" json:"data_handle,omitempty"`
	Log                    `bson:"log,omitempty" json:"log,omitempty"`
	Reconciliation         `bson:"reconciliation,omitempty" json:"reconciliation,omitempty"`
	Session                `bson:"session,omitempty" json:"session,omitempty"`
	ThreadYield            `bson:"thread-yield,omitempty" json:"thread_yield,omitempty"`
	Transaction            `bson:"transaction,omitempty" json:"transaction,omitempty"`
	ConcurrentTransactions `bson:"concurrentTransactions,omitempty" json:"concurrent_transactions,omitempty"`
}

// LSM ...
type LSM struct {
	SleepForLSMCheckpointThrottle       uint `bson:"sleep for LSM checkpoint throttle" json:"sleep_for_lsm_checkpoint_throttle,omitempty"`
	SleepForLSMMergeThrottle            uint `bson:"sleep for LSM merge throttle" json:"sleep_for_lsm_merge_throttle,omitempty"`
	RowsMergedInAnLSMTree               uint `bson:"rows merged in an LSM tree" json:"rows_merged_in_an_lsm_tree,omitempty"`
	ApplicationWorkUnitsCurrentlyQueued uint `bson:"application work units currently queued" json:"application_work_units_currently_queued,omitempty"`
	MergeWorkUnitsCurrentlyQueued       uint `bson:"merge work units currently queued" json:"merge_work_units_currently_queued,omitempty"`
	TreeQueueHitMaximum                 uint `bson:"tree queue hit maximum" json:"tree_queue_hit_maximum,omitempty"`
	SwitchWorkUnitsCurrentlyQueued      uint `bson:"switch work units currently queued" json:"switch_work_units_currently_queued,omitempty"`
	TreeMaintenanceOperationsScheduled  uint `bson:"tree maintenance operations scheduled" json:"tree_maintenance_operations_scheduled,omitempty"`
	TreeMaintenanceOperationsDiscarded  uint `bson:"tree maintenance operations discarded" json:"tree_maintenance_operations_discarded,omitempty"`
	TreeMaintenanceOperationsExecuted   uint `bson:"tree maintenance operations executed" json:"tree_operations_executed,omitempty"`
}

// Async ...
type Async struct {
	NumberOfAllocationStateRaces              uint `bson:"number of allocation state races" json:"number_of_allocation_state_races,omitempty,omitempty"`
	NumberOfOperationSlotsViewedForAllocation uint `bson:"number of operation slots viewed for allocation" json:"number_of_operation_slots_viewed_for_allocation,omitempty,omitempty"`
	CurrentWorkQueueLength                    uint `bson:"current work queue length" json:"current_work_queue_length,omitempty,omitempty"`
	NumberOfFlushCalls                        uint `bson:"number of flush calls" json:"number_of_flush_calls,omitempty,omitempty"`
	NumberOfTimesOperationAllocationFailed    uint `bson:"number of times operation allocation failed" json:"number_of_times_operation_allocation_failed,omitempty"`
	MaximumWorkQueueLength                    uint `bson:"maximum work queue length" json:"maximum_work_queue_length,omitempty"`
	NumberOfTimesWorkerFoundNoWork            uint `bson:"number of times worker found no work" json:"number_of_times_worker_found_no_work,omitempty"`
	TotalAllocations                          uint `bson:"total allocations" json:"total_allocations,omitempty"`
	TotalCompactCalls                         uint `bson:"total compact calls" json:"total_compact_calls,omitempty"`
	TotalInsertCalls                          uint `bson:"total insert calls" json:"total_insert_calls,omitempty"`
	TotalRemoveCalls                          uint `bson:"total remove calls" json:"total_remove_calls,omitempty"`
	TotalSearchCalls                          uint `bson:"total search calls" json:"total_search_calls,omitempty"`
	TotalUpdateCalls                          uint `bson:"total update calls" json:"total_update_calls,omitempty"`
}

// BlockManager ...
type BlockManager struct {
	MappedBytesRead  uint `bson:"mapped bytes read" json:"mapped_bytes_read,omitempty"`
	BytesRead        uint `bson:"bytes read" json:"bytes_read,omitempty"`
	BytesWritten     uint `bson:"bytes written" json:"bytes_written,omitempty"`
	MappedBlocksRead uint `bson:"mapped blocks read" json:"mapped_blocked_read,omitempty"`
	BlocksPreLoaded  uint `bson:"blocks pre-loaded" json:"blocks_pre_loaded,omitempty"`
	BlocksRead       uint `bson:"blocks read" json:"blocks_read,omitempty"`
	BlocksWritten    uint `bson:"blocks written" json:"blocks_written,omitempty"`
}

// Cache ...
type Cache struct {
	TrackedDirtyBytesInTheCache                         uint `bson:"tracked dirty bytes in the cache" json:"tracked_dirty_bytes_in_the_cache,omitempty"`
	BytesCurrentlyInTheCache                            uint `bson:"bytes currently in the cache" json:"bytes_currently_in_the_cache,omitempty"`
	MaximumBytesConfigured                              uint `bson:"maximum bytes configured" json:"maximum_bytes_configured,omitempty"`
	BytesReadIntoCache                                  uint `bson:"bytes read into cache" json:"bytes_read_into_cache,omitempty"`
	BytesWrittenFromCache                               uint `bson:"bytes written from cache" json:"bytes_written_from_cache,omitempty"`
	PagesEvictedByApplicationThreads                    uint `bson:"pages evicted by application threads" json:"pages_evicted_by_application_threads,omitempty"`
	CheckpointBlockedPageEviction                       uint `bson:"checkpoint blocked page eviction" json:"checkpoint_blocked_page_eviction,omitempty"`
	UnmodifiedPagesEvicted                              uint `bson:"unmodified pages evicted" json:"unmodified_pages_evicted,omitempty"`
	PageSplitDuringEvictionDeepenedTheTree              uint `bson:"page split during eviction deepened the tree" json:"page_split_during_eviction_deepened_the_tree,omitempty"`
	ModifiedPagesEvicted                                uint `bson:"modified pages evicted" json:"modified_pages_evicted,omitempty"`
	PagesSelectedForEvictionUnableToBeEvicted           uint `bson:"pages selected for eviction unable to be evicted" json:"pages_selected_for_eviction_unable_to_be_evicted,omitempty"`
	PagesEvictedBecauseTheyExceededTheInMemoryMaximum   uint `bson:"pages evicted because they exceeded the in-memory maximum" json:"pages_evicted_because_they_exceeded_the_in_memory_maximum,omitempty"`
	PagesEvictedBecauseTheyHadChainsOfDeletedItems      uint `bson:"pages evicted because they had chains of deleted items" json:"pages_evicted_because_they_had_chains_of_deleted_items,omitempty"`
	FailedEvictionOfPagesThatExceededTheInMemoryMaximum uint `bson:"failed eviction of pages that exceeded the in-memory maximum" json:"failed_eviction_of_pages_that_exceeded_the_in_memory_maximum,omitempty"`
	HazardPointerBlockedPageEviction                    uint `bson:"hazard pointer blocked page eviction" json:"hazard_pointer_blocked_page_eviction,omitempty"`
	InternalPagesEvicted                                uint `bson:"internal pages evicted" json:"internal_pages_evicted,omitempty"`
	MaximumPageSizeAtEviction                           uint `bson:"maximum page size at eviction" json:"max_page_size_at_eviction,omitempty"`
	EvictionServerCandidateQueueEmptyWhenToppingUp      uint `bson:"eviction server candidate queue empty when topping up" json:"eviction_server_candidate_queue_empty_when_topping_up,omitempty"`
	EvictionServerCandidateQueueNotEmptyWhenToppingUp   uint `bson:"eviction server candidate queue not empty when topping up" json:"eviction_server_candidate_queue_not_empty_when_topping_up,omitempty"`
	EvictionServerEvictingPages                         uint `bson:"eviction server evicting pages" json:"eviction_server_evicting_pages,omitempty"`
	EvictionServerPopulatingQueueButNotEvictingPages    uint `bson:"eviction server populating queue, but not evicting pages" json:"eviction_server_populating_queue_but_not_evicting_pages,omitempty"`
	EvictionServerUnableToReachEvictionGoal             uint `bson:"eviction server unable to reach eviction goal" json:"eviction_server_unable_to_reach_eviction_goal,omitempty"`
	PagesSplitDuringEviction                            uint `bson:"pages split during eviction" json:"pages_split_during_eviction,omitempty"`
	PagesWalkedForEviction                              uint `bson:"pages walked for eviction" json:"pages_walked_for_eviction,omitempty"`
	EvictionWorkerThreadEvictingPages                   uint `bson:"eviction worker thread evicting pages" json:"eviction_worker_thread_evicting_pages,omitempty"`
	InMemoryPageSplits                                  uint `bson:"in-memory page splits" json:"in_memory_page_splits,omitempty"`
	PercentageOverhead                                  uint `bson:"percentage overhead" json:"percentage_overhead,omitempty"`
	TrackedDirtyPagesInTheCache                         uint `bson:"tracked dirty pages in the cache" json:"tracked_dirty_pages_in_the_cache,omitempty"`
	PagesCurrentlyHeldInTheCache                        uint `bson:"pages currently held in the cache" json:"pages_currently_held_in_the_cache,omitempty"`
	PagesReadIntoCache                                  uint `bson:"pages read into cache" json:"pages_read_into_cache,omitempty"`
	PagesWrittenFromCache                               uint `bson:"pages written from cache" json:"pages_written_from_cache,omitempty"`
}

// Connection ...
type Connection struct {
	PThreadMutexConditionWaitCalls       uint `bson:"pthread mutex condition wait calls" json:"pthread_mutex_condition_wait_calls,omitempty"`
	FilesCurrentlyOpen                   uint `bson:"files currently open" json:"files_currently_open,omitempty"`
	MemoryAllocations                    uint `bson:"memory allocations" json:"memory_allocations,omitempty"`
	MemoryFrees                          uint `bson:"memory frees" json:"memory_frees,omitempty"`
	MemoryReAllocations                  uint `bson:"memory re-allocations" json:"memory_re_allocations,omitempty"`
	TotalReadIOs                         uint `bson:"total read I/Os" json:"total_read_IOs,omitempty"`
	PThreadMutexSharedLockReadLockCalls  uint `bson:"pthread mutex shared lock read-lock calls" json:"pthread_mutex_shared_lock_read_lock_calls,omitempty"`
	PThreadMutexSharedLockWriteLockCalls uint `bson:"pthread mutex shared lock write-lock calls" json:"pthread_mutex_shared_lock_write_lock_calls,omitempty"`
	TotalWriteIOs                        uint `bson:"total write I/Os" json:"total_write_IOs,omitempty"`
}

// WiredTigerCursor ...
type WiredTigerCursor struct {
	CursorCreateCalls     uint `bson:"cursor create calls" json:"cursor_create_calls,omitempty"`
	CursorInsertCalls     uint `bson:"cursor insert calls" json:"cursor_insert_calls,omitempty"`
	CursorNextCalls       uint `bson:"cursor next calls" json:"cursor_next_calls,omitempty"`
	CursorPrevCalls       uint `bson:"cursor prev calls" json:"cursor_prev_calls,omitempty"`
	CursorRemoveCalls     uint `bson:"cursor remove calls" json:"cursor_remove_calls,omitempty"`
	CursorResetCalls      uint `bson:"cursor reset calls" json:"cursor_reset_calls,omitempty"`
	CursorSearchCalls     uint `bson:"cursor search calls" json:"cursor_search_calls,omitempty"`
	CursorSearchNearCalls uint `bson:"cursor search near calls" json:"cursor_search_near_calls,omitempty"`
	CursorUpdateCalls     uint `bson:"cursor update calls" json:"cursor_update_calls,omitempty"`
}

// DataHandle ...
type DataHandle struct {
	ConnectionDHandlesSwept       uint `bson:"connection dhandles swept" json:"connection_d_handles_swept,omitempty"`
	ConnectionCandidateReferences uint `bson:"connection candidate references" json:"connection_candidate_referenced,omitempty"`
	ConnectionSweeps              uint `bson:"connection sweeps" json:"connection_sweeps,omitempty"`
	ConnectionsTimeOfDeathSets    uint `bson:"connection time-of-death sets" json:"connection_time_of_death_sets,omitempty"`
	SessionDHandlesSwept          uint `bson:"connection dhandles swept" json:"connection_d_handles_swept,omitempty"`
	SessionSweepAttempts          uint `bson:"session sweep attempts" json:"session_sweep_attempts,omitempty"`
}

// Log ...
type Log struct {
	LogBufferSizeIncreases                       uint `bson:"log buffer size increases" json:"log_buff_size_increases,omitempty"`
	TotalLogBufferSize                           uint `bson:"total log buffer size" json:"total_log_buffer_size,omitempty"`
	LogBytesOfPayLoadData                        uint `bson:"log bytes of payload data" json:"load_bytes_of_payload_data,omitempty"`
	LogBytesWritten                              uint `bson:"log bytes written" json:"log_bytes_written,omitempty"`
	YieldsWaitingForPreviousLogFileClose         uint `bson:"yields waiting for previous log file close" json:"yields_waiting_for_previous_log_file_close,omitempty"`
	TotalSizeOfCompressedRecords                 uint `bson:"total size of compressed records" json:"total_size_of_compressed_records,omitempty"`
	TotalInMemorySizeOfCompressedRecords         uint `bson:"total in-memory size of compressed records" json:"total_in_memory_size_of_compressed_records,omitempty"`
	LogRecordsTooSmallToCompress                 uint `bson:"log records too small to compress" json:"log_records_too_small_to_compress,omitempty"`
	LogRecordsNotCompressed                      uint `bson:"log records not compressed" json:"log_records_not_compressed,omitempty"`
	LogRecordsCompressed                         uint `bson:"log records compressed" json:"log_records_compressed,omitempty"`
	MaximumLogFileSize                           uint `bson:"maximum log file size" json:"max_log_file_size,omitempty"`
	PreAllocatedLogFilesPrepared                 uint `bson:"pre-allocated log files prepared" json:"pre_allocated_log_files_prepared,omitempty"`
	NumberOfPreAllocatedLogFilesToCreate         uint `bson:"number of pre-allocated log files to create" json:"number_of_pre_allocated_log_files_to_create,omitempty"`
	PreAllocatedLogFilesUsed                     uint `bson:"pre-allocated log files used" json:"pre_allocated_log_files_used,omitempty"`
	LogReadOperations                            uint `bson:"log read operations" json:"log_read_operations,omitempty"`
	LogReleaseAdvancesWriteLSN                   uint `bson:"log release advances write LSN" json:"log_release_advances_write_lsn,omitempty"`
	RecordsProcessedByLogScan                    uint `bson:"records processed by log scan" json:"records_processed_by_log_scan,omitempty"`
	LogScanRecordsRequiringTwoReads              uint `bson:"log scan records requiring two reads" json:"log_scan_records_requiring_two_reads,omitempty"`
	LogScanOperations                            uint `bson:"log scan operations" json:"log_scan_operations,omitempty"`
	ConsolidatedSlotClosures                     uint `bson:"consolidated slot closures" json:"consolidated slot closures,omitempty"`
	LoggingBytesConsolidated                     uint `bson:"logging bytes consolidated" json:"logging_bytes_consolidated,omitempty"`
	ConsolidatedSlotJoins                        uint `bson:"consolidated slot joins" json:"consolidated_slot_joins,omitempty"`
	ConsolidatedSlotJoinsRaces                   uint `bson:"consolidated slot joins races" json:"consolidated_slot_joins_races,omitempty"`
	SlotsSelectedForSwitchingThatWereUnavailable uint `bson:"slots selected for switching that were unavailable" json:"slots_selected_for_switching_that_were_unavailable,omitempty"`
	RecordSizeExceededMaximum                    uint `bson:"record size exceeded maximum" json:"record_size_exceeded_maximum,omitempty"`
	FailedToFindASlotLargeEnoughForRecord        uint `bson:"failed to find a slot large enough for record" json:"failed_to_find_a_slot_large_enough_for_record,omitempty"`
	ConsolidatedSlotJoinTransitions              uint `bson:"consolidated slot join transitions" json:"consolidated slot join transitions,omitempty"`
	LogSyncOperations                            uint `bson:"log sync operations" json:"log_sync_operations,omitempty"`
	LogSyncDirOperations                         uint `bson:"log sync_dir operations" json:"log_sync_dir_operations,omitempty"`
	LogWriteOperations                           uint `bson:"log write operations" json:"log_write_operations,omitempty"`
}

// Reconciliation ...
type Reconciliation struct {
	PageReconciliationCalls            uint `bson:"page reconciliation calls" json:"page_reconciliation_calls,omitempty"`
	PageReconciliationCallsForEviction uint `bson:"page reconciliation calls for eviction" json:"page_reconciliation_calls_for_eviction,omitempty"`
	SplitBytesCurrentlyAwaitingFree    uint `bson:"split bytes currently awaiting free" json:"split_bytes_currently_awaiting_free,omitempty"`
	SplitObjectsCurrentlyAwaitingFree  uint `bson:"split objects currently awaiting free" json:"split_objects_currently_awaiting_free,omitempty"`
}

// Session ...
type Session struct {
	OpenCursorCount  uint `bson:"open cursor count" json:"open_cursor_count,omitempty"`
	OpenSessionCount uint `bson:"open session count" json:"open_session_count,omitempty"`
}

// ThreadYield ...
type ThreadYield struct {
	PageAcquireBusyBlocked     uint `bson:"page acquire busy blocked" json:"page_acquire_busy_blocked,omitempty"`
	PageAcquireEvictionBlocked uint `bson:"page acquire eviction blocked" json:"page_acquire_eviction_blocked,omitempty"`
	PageAcquireLockedBlocked   uint `bson:"page acquire locked blocked" json:"page_acquire_locked_blocked,omitempty"`
	PageAcquireReadBlocked     uint `bson:"page acquire read blocked" json:"page_acquire_read_blocked,omitempty"`
	PageAcquireTimeSleeping    uint `bson:"page acquire time sleeping" json:"page_acquire_time_sleeping,omitempty"`
}

// Transaction ...
type Transaction struct {
	TransactionBegins                     uint `bson:"transaction begins" json:"transaction_begins,omitempty"`
	TransactionCheckPoints                uint `bson:"transaction checkpoints" json:"tansaction_checkpoints,omitempty"`
	TransactionCheckPointCurrentlyRunning uint `bson:"transaction checkpoint curently running" json:"transaction_checkpoint_currently_running,omitempty"`
	TransactionCheckPointMaxTime          uint `bson:"transaction checkpoint max time (msecs)" json:"transaction_checkpoint_max_time_ms,omitempty"`
	TransactionCheckPointMinTime          uint `bson:"transaction checkpoint min time (msecs)" json:"transaction_checkpoint_min_time_ms,omitempty"`
	TransactionCheckPointMostRecentTime   uint `bson:"transaction checkpoint most recent time (msecs)" json:"transaction_checkpoint_most_recent_time_ms,omitempty"`
	TransactionCheckPointTotalTime        uint `bson:"transaction checkpoint total time (msecs)" json:"transaction_checkpoint_total_time_ms,omitempty"`
	TransactionsCommitted                 uint `bson:"transactions committed" json:"transactions_committed,omitempty"`
	TransactionFailuresDueToCacheOverflow uint `bson:"transaction failures due to chache overflow" json:"transaction_failures_due_to_cache_overflow,omitempty"`
	TransactionRangeOfIDsCurrentlyPinned  uint `bson:"transaction range of IDs currently pinned" json:"transaction_range_of_ids_currently_pinned,omitempty"`
	TransactionsRolledBack                uint `bson:"transactions rolled back" json:"transactions_rolled_back,omitempty"`
}

// ConcurrentTransactions ...
type ConcurrentTransactions struct {
	Write struct {
		Out          uint `bson:"out" json:"out,omitempty"`
		Available    uint `bson:"available" json:"available,omitempty"`
		TotalTickets uint `bson:"totalTickets" json:"total_tickets,omitempty"`
	} `bson:"write" json:"write,omitempty"`

	Read struct {
		Out          uint `bson:"out" json:"out,omitempty"`
		Available    uint `bson:"available" json:"available,omitempty"`
		TotalTickets uint `bson:"totalTickets" json:"total_tickets,omitempty"`
	} `bson:"read" json:"read,omitempty"`
}
