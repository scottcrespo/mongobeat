package dbstats

// DbStats represents the fields returned from a call to db.stats() in Mongo
type DbStats struct {
	Db          string `bson:"db,omitempty"            json:"db,omitempty"`
	Collections uint   `bson:"collections,omitempty"  json:"collections,omitempty"`
	Objects     uint64 `bson:"objects,omitempty"       json:"objects,omitempty"`
	AvgObjSize  uint64 `bson:"avgObjectSize,omitempty" json:"avg_obj_size,omitempty"`
	DataSize    uint64 `bson:"dataSize,omitempty"      json:"data_size,omitempty"`
	StorageSize uint64 `bson:"storageSize,omitempty"   json:"storage_size,omitempty"`
	NumExtents  uint64 `bson:"numExtents,omitempty"    json:"num_extents,omitempty"`
	Indexes     uint64 `bson:"indexes,omitempty"       json:"indexes,omitempty"`
	IndexSize   uint64 `bson:"indexSize,omitempty"     json:"index_size,omitempty"`
	FileSize    uint64 `bson:"fileSize,omitempty"      json:"file_size,omitempty"`
	Ok          uint   `bson:"ok,omitempty"            json:"ok,omitempty"`
}
