package serverstatus

// StorageEngine is the 'storageEngine' key of serverStatus output
type StorageEngine struct {
	Name                   string `bson:"name"                       json:"name"`
	SupportsCommittedReads bool   `bson:"supportsCommittedReads"     json:"supports_committed_reads"`
	Persistent             bool   `bson:"persistent"                 json:"persistent"`
}
