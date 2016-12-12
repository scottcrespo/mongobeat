package serverstatus

// Locks is the 'locks' key of serverStatus output
type Locks struct {
	Global        LockType `bson:"Global,omitempty" json:"global,omitempty"`
	MMAPV1Journal LockType `bson:"MMAPV1Journal,omitempty" json:"mmap_v1_journal,omitempty"`
	Database      LockType `bson:"Database,omitempty" json:"database,omitempty"`
	Collection    LockType `bson:"Collection,omitempty" json:"collection,omitempty"`
	Metadata      LockType `bson:"Metadata,omitempty" json:"metadata,omitempty"`
	Oplog         LockType `bson:"oplog,omitempty" json:"oplog,omitempty"`
}

// LockType ..
type LockType struct {
	AcquireCount        LockMode `bson:"acquireCount,omitempty" json:"acquire_count,omitempty"`
	AcquireWaitCount    LockMode `bson:"acquireWaitCount,omitempty" json:"acquire_wait_count,omitempty"`
	TimeAcquiringMicros LockMode `bson:"timeAcquiringMicros,omitempty" json:"time_acquiring_micros,omitempty"`
	DeadLockCount       LockMode `bson:"deadLockCount,omitempty" json:"dead_lock_count,omitempty"`
}

// LockMode ...
type LockMode struct {
	R  uint64 `bson:"r,omitempty" json:"intent_shared_lock,omitempty"`
	RR uint64 `bson:"R,omitempty" json:"shared_lock,omitempty"`
	W  uint64 `bson:"w,omitempty" json:"intent_exclusive_lock,omitempty"`
	WW uint64 `bson:"W,omitempty" json:"exclusive_lock,omitempty"`
}
