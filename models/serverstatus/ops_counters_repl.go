package serverstatus

// OpsCountersRepl is the 'opscountersRepl' key of serverStatus output
type OpsCountersRepl struct {
	Insert  uint `bson:"insert"  json:"insert"`
	Query   uint `bson:"query"   json:"query"`
	Update  uint `bson:"update"  json:"update"`
	Delete  uint `bson:"delete"  json:"delete"`
	Getmore uint `bson:"getmore" json:"get_more"`
	Command uint `bson:"command" json:"command"`
}
