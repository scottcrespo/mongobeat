package serverstatus

// ExtraInfo is the 'extraInfo' key of serverStatus output
type ExtraInfo struct {
	Note       string `bson:"note"          json:"note"`
	PageFaults uint   `bson:"page_faults"   json:"page_faults"`
	// unix systems only
	HeapUsageBytes uint64 `bson:"heap_usage_bytes" json:"heap_usage_bytes"`
}
