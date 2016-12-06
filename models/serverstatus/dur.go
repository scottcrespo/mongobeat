package serverstatus

// Dur is the 'dur' key of serverStatus output
type Dur struct {
	Commmits           uint `bson:"commits"             json:"commits"`
	JournaledMB        uint `bson:"journaledMB"         json:"journaled_mb"`
	WriteToDataFilesMB uint `bson:"writeToDataFilesMB"  json:"write_to_data_files_mb"`
	Compression        uint `bson:"compression"         json:"compression"`
	CommitsInWriteLock uint `bson:"commitsInWriteLock"  json:"commits_in_write_lock"`
	EarlyCommits       uint `bson:"earlyCommits"        json:"early_commits"`

	TimeMS struct {
		Dt                 uint `bson:"dt"                   json:"dt"`
		PrepLogBuffer      uint `bson:"prepLogBuffer"        json:"prep_log_buffer"`
		WriteToJournal     uint `bson:"writeToJournal"       json:"write_to_journal"`
		RemapPrivateView   uint `bson:"remapPrivateView"     json:"remap_private_view"`
		Commits            uint `bson:"commits"              json:"commits"`
		CommitsInWriteLock uint `bson:"commitsInWriteLock"   json:"commits_in_write_lock"`
	} `bson:"timeMs" json:"time_ms"`
}
