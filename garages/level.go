package garages

type Level struct {
	Level           string  `json:"level"`
	Status          string  `json:"status"`
	TotalSpaces     int     `json:"total_spaces"`
	FreeSpaces      int     `json:"free_spaces"`
	PercentOccupied float64 `json:"percent_occupied"`
	OccupiedSpaces  int     `json:"occupied_spaces"`
}
