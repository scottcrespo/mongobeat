package serverstatus

// Asserts is the 'asserts' key of serverStatus output
type Asserts struct {
	Regular   uint `bson:"regular,omitempty"              json:"regular,omitempty"`
	Warning   uint `bson:"warning,omitempty"              json:"warning,omitempty"`
	Msg       uint `bson:"msg,omitempty"                  json:"msg,omitempty"`
	User      uint `bson:"user,omitempty"                 json:"user,omitempty"`
	Rollovers uint `bson:"rollovers,omitempty"            json:"rollovers,omitempty"`
}
