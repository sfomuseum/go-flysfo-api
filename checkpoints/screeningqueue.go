package checkpoints

type ScreeningQueue struct {
	Name                  string `json:"screening_queue_name"`
	IsOpen                string `json:"is_open"`
	CurrentProcessingRate int    `json:"current_processing_rate"`
	TotalExitCount        int    `json:"total_exit_count"`
	MaxProcessingRate     int    `json:"max_processing_rate"`
	AverageTime           int    `json:"average_time"`
	DayPaxCount           int    `json:"day_pax_count"`
	DayTimeOpen           int    `json:"day_time_open"`
}
