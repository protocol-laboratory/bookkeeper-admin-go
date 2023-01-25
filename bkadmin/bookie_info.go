package bkadmin

import "encoding/json"

type BookieInfo struct {
	FreeSpace  int64
	TotalSpace int64
}

func decodeBookieInfo(data []byte) (*BookieInfo, error) {
	info := &BookieInfo{}
	var result map[string]interface{}
	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	info.FreeSpace = int64(result["freeSpace"].(float64))
	info.TotalSpace = int64(result["totalSpace"].(float64))
	return info, nil
}
