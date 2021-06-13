package metadata

// AudioInfo describe common audio properties of the track file.
type AudioInfo struct {
	Samplerate int `json:"samplerate,omitempty"`
	AvgBitrate int `json:"avg_bitrate,omitempty"`
	Channels   int `json:"channels,omitempty"`
	SampleSize int `json:"sample_size,omitempty"`
}

// IsEmpty проверяет коллекцию как не инициализированную.
func (ai *AudioInfo) IsEmpty() bool {
	return *ai == AudioInfo{}
}

// Clean сбрасывает всю коллекцию в nil, если поля структуры не отличаются от нулевых значений.
func (ai *AudioInfo) Clean() {}
