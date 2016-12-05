package serverstatus

// Asserts is the 'asserts' key of serverStatus output
type Asserts struct {
	Regular   uint `bson:"regular"              json:"regular"`
	Warning   uint `bson:"warning"              json:"warning"`
	Msg       uint `bson:"msg"                  json:"msg"`
	User      uint `bson:"user"                 json:"user"`
	Rollovers uint `bson:"rollovers"            json:"rollovers"`
}
