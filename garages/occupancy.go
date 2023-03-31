package garages

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/sfomuseum/go-flysfo-api"
)

type Occupancy struct {
	GarageUUID      string   `json:"garage_uuid"`
	FreeSpaces      int      `json:"free_spaces"`
	OccupiedSpaces  int      `json:"occupied_spaces"`
	PercentOccupied float64  `json:"percent_occupied"`
	Status          string   `json:"status"`
	TotalSpaces     int      `json:"total_spaces"`
	Updated         string   `json:"updated_at"`
	Levels          []*Level `json:"levels"`
}

func (o *Occupancy) String() string {
	return fmt.Sprintf("%s %0.2f occupied", o.GarageUUID, o.PercentOccupied)
}

func GetOccupanciesForGarage(ctx context.Context, cl *api.Client, uuid string) (*Occupancy, error) {

	method := fmt.Sprintf("garages/%s/occupancy", uuid)

	rsp, err := cl.ExecuteMethod(ctx, method, nil)

	if err != nil {
		return nil, err
	}

	defer rsp.Close()

	var occupancy *Occupancy

	dec := json.NewDecoder(rsp)
	err = dec.Decode(&occupancy)

	if err != nil {
		return nil, err
	}

	return occupancy, nil
}
