package serverstatus

// Locks is the 'locks' key of serverStatus output
type Locks struct {
	Global        LockType `bson:"Global" json:"global"`
	MMAPV1Journal LockType `bson:"MMAPV1Journal" json:"mmap_v1_journal"`
	Database      LockType `bson:"Database" json:"database"`
	Collection    LockType `bson:"Collection" json:"collection"`
	Metadata      LockType `bson:"Metadata" json:"metadata"`
	Oplog         LockType `bson:"oplog" json:"oplog"`
}

// LockType ..
type LockType struct {
	AcquireCount        LockMode `bson:"acquireCount" json:"acquire_count"`
	AcquireWaitCount    LockMode `bson:"acquireWaitCount" json:"acquire_wait_count"`
	TimeAcquiringMicros LockMode `bson:"timeAcquiringMicros" json:"time_acquiring_micros"`
	DeadLockCount       LockMode `bson:"deadLockCount" json:"dead_lock_count"`
}

// LockMode ...
type LockMode struct {
	R  uint64 `bson:"r" json:"intent_shared_lock"`
	RR uint64 `bson:"R" json:"shared_lock"`
	W  uint64 `bson:"w" json:"intent_exclusive_lock"`
	WW uint64 `bson:"W json:exclusive_lock"`
}
