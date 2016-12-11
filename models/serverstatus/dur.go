package serverstatus

// Dur is the 'dur' key of serverStatus output
type Dur struct {
	Commmits           uint `bson:"commits,omitempty"             json:"commits,omitempty"`
	JournaledMB        uint `bson:"journaledMB,omitempty"         json:"journaled_mb,omitempty"`
	WriteToDataFilesMB uint `bson:"writeToDataFilesMB,omitempty"  json:"write_to_data_files_mb,omitempty"`
	Compression        uint `bson:"compression,omitempty"         json:"compression,omitempty"`
	CommitsInWriteLock uint `bson:"commitsInWriteLock,omitempty"  json:"commits_in_write_lock,omitempty"`
	EarlyCommits       uint `bson:"earlyCommits,omitempty"        json:"early_commits,omitempty"`

	TimeMS struct {
		Dt                 uint `bson:"dt,omitempty"                   json:"dt,omitempty"`
		PrepLogBuffer      uint `bson:"prepLogBuffer,omitempty"        json:"prep_log_buffer,omitempty"`
		WriteToJournal     uint `bson:"writeToJournal,omitempty"       json:"write_to_journal,omitempty"`
		RemapPrivateView   uint `bson:"remapPrivateView,omitempty"     json:"remap_private_view,omitempty"`
		Commits            uint `bson:"commits,omitempty"              json:"commits,omitempty"`
		CommitsInWriteLock uint `bson:"commitsInWriteLock,omitempty"   json:"commits_in_write_lock,omitempty"`
	} `bson:"timeMs,omitempty" json:"time_ms,omitempty"`
}
