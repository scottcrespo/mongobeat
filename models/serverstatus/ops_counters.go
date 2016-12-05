package serverstatus

// OpsCounters is the 'opscounters' key of serverStatus output
type OpsCounters struct {
	Insert  uint `bson:"insert"  json:"insert"`
	Query   uint `bson:"query"   json:"query"`
	Update  uint `bson:"update"  json:"update"`
	Delete  uint `bson:"delete"  json:"delete"`
	Getmore uint `bson:"getmore" json:"get_more"`
	Command uint `bson:"command" json:"command"`
}
