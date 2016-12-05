package serverstatus

// Connections is the 'connections' key of serverStatus output
type Connections struct {
	Current      uint   `bson:"current"       json:"current"`
	Available    uint   `bson:"available"     json:"available"`
	TotalCreated uint64 `bson:"totalCreated"  json:"total_created"`
}
