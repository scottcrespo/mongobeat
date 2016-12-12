package serverstatus

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Repl ...
type Repl struct {
	SetName    string        `bson:"setName,omitempty" json:"set_name,omitempty"`
	SetVersion uint          `bson:"setVersion,omitempty" json:"set_version,omitempty"`
	IsMaster   bool          `bson:"ismaster,omitempty" json:"is_master,omitempty"`
	Secondary  bool          `bson:"secondary,omitempty" json:"secondary,omitempty"`
	Hosts      []string      `bson:"hosts,omitempty" json:"hosts,omitempty"`
	Primary    string        `bson:"primary,omitempty" json:"primary,omitempty"`
	Me         string        `bson:"me,omitempty" json:"me,omitempty"`
	ElectionID bson.ObjectId `bson:"electionId,omitempty" json:"election_id,omitempty"`
	RBid       uint          `bson:"rbid,omitempty" json:"rbid,omitempty"`
	// 3.0 uses key 'slaves'
	Slaves []ReplProgress `bson:"slaves,omitempty,omitempty" json:"replication_progress,omitempty"`
	// Changed in 3.2 to 'replication_progress'
	ReplicationProgress []ReplProgress `bson:"replicationProgress,omitempty,omitempty" json:"replication_progress,omitempty"`
}

// ReplProgress ...
type ReplProgress struct {
	Rid      bson.ObjectId `bson:"rid,omitempty" json:"rid,omitempty"`
	OpTime   time.Time     `bson:"optime,omitempty" json:"op_time,omitempty"`
	Host     string        `bson:"host,omitempty" json:"host,omitempty"`
	MemberID uint          `bson:"memberID,omitempty" json:"member_id,omitempty"`
}
