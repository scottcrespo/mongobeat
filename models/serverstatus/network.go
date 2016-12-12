package serverstatus

// Network is the 'network' key of serverStatus output
type Network struct {
	BytesIn     uint64 `bson:"bytesIn,omitempty"     json:"bytes_in,omitempty"`
	BytesOut    uint64 `bson:"bytesOut,omitempty"    json:"bytes_out,omitempty"`
	NumRequests uint64 `bson:"numRequests,omitempty" json:"num_requests,omitempty"`
}
