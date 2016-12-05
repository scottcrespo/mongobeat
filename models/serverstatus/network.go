package serverstatus

// Network is the 'network' key of serverStatus output
type Network struct {
	BytesIn     uint64 `bson:"bytesIn"     json:"bytes_in"`
	BytesOut    uint64 `bson:"bytesOut"    json:"bytes_out"`
	NumRequests uint64 `bson:"numRequests" json:"num_requests"`
}
