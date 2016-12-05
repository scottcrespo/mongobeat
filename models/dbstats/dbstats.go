package dbstats

// DbStats represents the fields returned from a call to db.stats() in Mongo
type DbStats struct {
	Db          string `bson:"db"            json:"db"`
	Collections uint   `bson:"collections"   json:"collections"`
	Objects     uint64 `bson:"objects"       json:"objects"`
	AvgObjSize  uint64 `bson:"avgObjectSize" json:"avg_obj_size"`
	DataSize    uint64 `bson:"dataSize"      json:"data_size"`
	StorageSize uint64 `bson:"storageSize"   json:"storage_size"`
	NumExtents  uint64 `bson:"numExtents"    json:"num_extents"`
	Indexes     uint64 `bson:"indexes"       json:"indexes"`
	IndexSize   uint64 `bson:"indexSize"     json:"index_size"`
	FileSize    uint64 `bson:"fileSize"      json:"file_size"`
	Ok          uint   `bson:"ok"            json:"ok"`
}
