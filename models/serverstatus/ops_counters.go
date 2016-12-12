package serverstatus

// OpsCounters is the 'opscounters' key of serverStatus output
type OpsCounters struct {
	Insert  uint `bson:"insert,omitempty"  json:"insert,omitempty"`
	Query   uint `bson:"query,omitempty"   json:"query,omitempty"`
	Update  uint `bson:"update,omitempty"  json:"update,omitempty"`
	Delete  uint `bson:"delete,omitempty"  json:"delete,omitempty"`
	Getmore uint `bson:"getmore,omitempty" json:"get_more,omitempty"`
	Command uint `bson:"command,omitempty" json:"command,omitempty"`
}
