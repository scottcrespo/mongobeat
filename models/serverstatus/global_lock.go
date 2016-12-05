package serverstatus

// GlobalLock is the 'globalLock' key of serverStatus output
type GlobalLock struct {
	TotalTime    uint64 `bson:"totalTime" json:"total_time"`
	CurrentQueue struct {
		Total   uint `bson:"total"   json:"total"`
		Readers uint `bson:"readers" json:"readers"`
		Writers uint `bson:"writers" json:"writers"`
	} `bson:"currentQueue" json:"current_queue"`
	ActiveClients struct {
		Total   uint `bson:"total"   json:"total"`
		Readers uint `bson:"readers" json:"readers"`
		Writers uint `bson:"writers" json:"writers"`
	} `bson:"activeClients" json:"active_clients"`
}
