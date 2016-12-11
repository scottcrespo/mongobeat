package serverstatus

// GlobalLock is the 'globalLock' key of serverStatus output
type GlobalLock struct {
	TotalTime    uint64 `bson:"totalTime,omitempty" json:"total_time,omitempty"`
	CurrentQueue struct {
		Total   uint `bson:"total,omitempty"   json:"total,omitempty"`
		Readers uint `bson:"readers,omitempty" json:"readers,omitempty"`
		Writers uint `bson:"writers,omitempty" json:"writers,omitempty"`
	} `bson:"currentQueue,omitempty" json:"current_queue,omitempty"`
	ActiveClients struct {
		Total   uint `bson:"total,omitempty"   json:"total,omitempty"`
		Readers uint `bson:"readers,omitempty" json:"readers,omitempty"`
		Writers uint `bson:"writers,omitempty" json:"writers,omitempty"`
	} `bson:"activeClients,omitempty" json:"active_clients,omitempty"`
}
