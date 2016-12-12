package serverstatus

// Mem is the 'mem' key of serverStatus output
type Mem struct {
	Bits              uint `bson:"bits,omitempty"                  json:"bits,omitempty"`
	Resident          uint `bson:"resident,omitempty"              json:"resident,omitempty"`
	Virtual           uint `bson:"virtual,omitempty"               json:"virtual,omitempty"`
	Supported         bool `bson:"supported,omitempty"             json:"supported,omitempty"`
	Mapped            uint `bson:"mapped,omitempty"                json:"mapped,omitempty"`
	MappedWithJournal uint `bson:"mappedWithJournal,omitempty"     json:"mapped_with_journal,omitempty"`
}
