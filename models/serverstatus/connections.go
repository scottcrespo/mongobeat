package serverstatus

// Connections is the 'connections' key of serverStatus output
type Connections struct {
	Current      uint   `bson:"current,omitempty"       json:"current,omitempty"`
	Available    uint   `bson:"available,omitempty"     json:"available,omitempty"`
	TotalCreated uint64 `bson:"totalCreated,omitempty"  json:"total_created,omitempty"`
}
