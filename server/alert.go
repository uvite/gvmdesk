package surge

type Alert struct {
	Name          string
	Symbol        string
	Interval      string
	Path          string //only for local
	NumChunks     int
	IsDownloading bool
	IsUploading   bool
	IsPaused      bool
	IsMissing     bool
	IsHashing     bool //only for local
	IsTracked     bool //only for local
	IsAvailable   bool //only for local
	ChunkMap      []byte
	ChunksShared  int
	Progress      float32 //only for remote
	Topic         string
	DateTimeAdded int64
}
