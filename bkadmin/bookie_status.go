package bkadmin

import "encoding/json"

type BookieStatus struct {
	Running                        bool
	ReadOnly                       bool
	ShuttingDown                   bool
	AvailableForHighPriorityWrites bool
}

func decodeBookieStatus(data []byte) (*BookieStatus, error) {
	bookieStatus := &BookieStatus{}
	var result map[string]bool
	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	bookieStatus.Running = result["running"]
	bookieStatus.ReadOnly = result["readOnly"]
	bookieStatus.ShuttingDown = result["shuttingDown"]
	bookieStatus.AvailableForHighPriorityWrites = result["availableForHighPriorityWrites"]
	return bookieStatus, nil
}
