package serverstatus

// Mem is the 'mem' key of serverStatus output
type Mem struct {
	Bits              uint `bson:"bits"                  json:"bits"`
	Resident          uint `bson:"resident"              json:"resident"`
	Virtual           uint `bson:"virtual"               json:"virtual"`
	Supported         bool `bson:"supported"             json:"supported"`
	Mapped            uint `bson:"mapped"                json:"mapped"`
	MappedWithJournal uint `bson:"mappedWithJournal"     json:"mapped_with_journal"`
}
