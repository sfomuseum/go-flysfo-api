package checkpoints

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/sfomuseum/go-flysfo-api"
)

type CheckpointList struct {
	Checkpoints []*Checkpoint
}

type Checkpoint struct {
	UUID     string `json:"checkpoint_uuid"`
	Name     string `json:"checkpoint_name"`
	Code     string `json:"checkpoint_code"`
	Hours    string `json:"checkpoint_hours"`
	Gates    string `json:"checkpoint_gates"`
	Terminal int64  `json:"terminal_id"`
	Updated  string `json:"updated_at"`
	Note     string `json:"note"`
}

func (c *Checkpoint) String() string {
	return c.Name
}

func ListCheckpoints(ctx context.Context, cl *api.Client, params *url.Values) (*CheckpointList, error) {

	rsp, err := cl.ExecuteMethod(ctx, "checkpoints", params)

	if err != nil {
		return nil, err
	}

	defer rsp.Close()

	var checkpoints *CheckpointList

	dec := json.NewDecoder(rsp)
	err = dec.Decode(&checkpoints)

	if err != nil {
		return nil, err
	}

	return checkpoints, nil
}
