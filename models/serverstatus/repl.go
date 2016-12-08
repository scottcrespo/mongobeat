package serverstatus

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Repl ...
type Repl struct {
	SetName    string        `bson:"setName" json:"set_name"`
	SetVersion uint          `bson:"setVersion" json:"set_version"`
	IsMaster   bool          `bson:"ismaster" json:"is_master"`
	Secondary  bool          `bson:"secondary" json:"secondary"`
	Hosts      []string      `bson:"hosts" json:"hosts"`
	Primary    string        `bson:"primary" json:"primary"`
	Me         string        `bson:"me" json:"me"`
	ElectionID bson.ObjectId `bson:"electionId" json:"election_id"`
	RBid       uint          `bson:"rbid" json:"rbid"`
	// 3.0 uses key 'slaves'
	Slaves []ReplProgress `bson:"slaves,omitempty" json:"replication_progress"`
	// Changed in 3.2 to 'replication_progress'
	ReplicationProgress []ReplProgress `bson:"replicationProgress,omitempty" json:"replication_progress"`
}

// ReplProgress ...
type ReplProgress struct {
	Rid      bson.ObjectId `bson:"rid" json:"rid"`
	OpTime   time.Time     `bson:"optime" json:"op_time"`
	Host     string        `bson:"host" json:"host"`
	MemberID uint          `bson:"memberID" json:"member_id"`
}
