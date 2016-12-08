package serverstatus

// StorageEngine is the 'storageEngine' key of serverStatus output
type StorageEngine struct {
	Name string `bson:"name"                       json:"name"`
	// New in 3.2
	SupportsCommittedReads bool `bson:"supportsCommittedReads"     json:"supports_committed_reads"`
	// New in 3.2.6
	Persistent bool `bson:"persistent"                 json:"persistent"`
}
