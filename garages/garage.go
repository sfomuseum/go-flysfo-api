package garages

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/sfomuseum/go-flysfo-api"
)

type GarageList struct {
	Garages []*Garage
}

type Garage struct {
	UUID        string  `json:"garage_uuid"`
	ID          int64   `json:"garage_id"`
	Name        string  `json:"garage_name"`
	Type        string  `json:"garage_type"`
	Description string  `json:"garage_type_desc"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Updated     string  `json:"updated_at"`
}

func (g *Garage) String() string {
	return g.Name
}

func GetGarage(ctx context.Context, cl *api.Client, uuid string) (*Garage, error) {

	method := fmt.Sprintf("garages/%s", uuid)

	rsp, err := cl.ExecuteMethod(ctx, method, nil)

	if err != nil {
		return nil, err
	}

	defer rsp.Close()

	var g *Garage

	dec := json.NewDecoder(rsp)
	err = dec.Decode(&g)

	if err != nil {
		return nil, err
	}

	return g, nil
}

func ListGarages(ctx context.Context, cl *api.Client, params *url.Values) (*GarageList, error) {

	rsp, err := cl.ExecuteMethod(ctx, "garages", params)

	if err != nil {
		return nil, err
	}

	defer rsp.Close()

	var garages *GarageList

	dec := json.NewDecoder(rsp)
	err = dec.Decode(&garages)

	if err != nil {
		return nil, err
	}

	return garages, nil
}
