package serverstatus

// ExtraInfo is the 'extraInfo' key of serverStatus output
type ExtraInfo struct {
	Note       string `bson:"note,omitempty"          json:"note,omitempty"`
	PageFaults uint   `bson:"page_faults,omitempty"   json:"page_faults,omitempty"`
	// unix systems only
	HeapUsageBytes uint64 `bson:"heap_usage_bytes,omitempty" json:"heap_usage_bytes,omitempty"`
}
