package serverstatus

// StorageEngine is the 'storageEngine' key of serverStatus output
type StorageEngine struct {
	Name string `bson:"name,omitempty" json:"name,omitempty"`
	// New in 3.2
	SupportsCommittedReads bool `bson:"supportsCommittedReads,omitempty" json:"supports_committed_reads,omitempty"`
	// New in 3.2.6
	Persistent bool `bson:"persistent,omitempty" json:"persistent,omitempty"`
}
