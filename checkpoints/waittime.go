package checkpoints

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/sfomuseum/go-flysfo-api"
)

type WaitTime struct {
	CheckpointUUID        string            `json:"checkpoint_uuid"`
	CheckpointName        string            `json:"checkpoint_name"`
	OpenScreeningQueues   int               `json:"open_screening_queues"`
	CurrentProcessingRate int               `json:"current_processing_rate"`
	Updated               string            `json:"updated_at": "2021-02-26T18:01:01.728-0800"`
	PreScreeningQueues    []*ScreeningQueue `json:"prescreening_queues"`
	ScreeningQueuees      []*ScreeningQueue `json:"screening_queues"`
}

func (wt *WaitTime) String() string {
	return fmt.Sprintf("%s %d (processing rate)", wt.CheckpointName, wt.CurrentProcessingRate)
}

func GetWaitTimeForCheckpoint(ctx context.Context, cl *api.Client, uuid string) (*WaitTime, error) {

	method := fmt.Sprintf("checkpoints/%s/waittime", uuid)

	rsp, err := cl.ExecuteMethod(ctx, method, nil)

	if err != nil {
		return nil, err
	}

	defer rsp.Close()

	var wt *WaitTime

	dec := json.NewDecoder(rsp)
	err = dec.Decode(&wt)

	if err != nil {
		return nil, err
	}

	return wt, nil
}
