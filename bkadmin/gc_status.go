package bkadmin

import "encoding/json"

type GarbageCollectionStatus struct {
	ForceCompacting         bool
	MajorCompacting         bool
	MinorCompacting         bool
	LastMajorCompactionTime int64
	LastMinorCompactionTime int64
	MajorCompactionCounter  int64
	MinorCompactionCounter  int64
}

func decodeGcStatus(data []byte) ([]*GarbageCollectionStatus, error) {
	var result []*GarbageCollectionStatus
	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
