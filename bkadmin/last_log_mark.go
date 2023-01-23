// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package bkadmin

import (
	"fmt"
	"regexp"
	"strconv"
)

type LastLogMark struct {
	logFileIdTxnMap map[int64]int64
}

// decodeLastLogMarkMap
// key format is "LastLogMark: Journal Id - ${LogFieldId}(${hex-log-field}.txn)"
// example value: LastLogMark: Journal Id - 0(0.txn)
// value format is Pos - ${LogFileOffset}
// example value: Pos - 0
func decodeLastLogMarkMap(data map[string]string) (*LastLogMark, error) {
	l := &LastLogMark{}
	l.logFileIdTxnMap = make(map[int64]int64, len(data))
	for key, value := range data {
		logFieldId, err := extractValueFromLastLogMarkKey(key)
		if err != nil {
			return nil, err
		}
		logFileOffset, err := extractValueFromLastLogMarkValue(value)
		if err != nil {
			return nil, err
		}
		l.logFileIdTxnMap[logFieldId] = logFileOffset
	}
	return l, nil
}

func extractValueFromLastLogMarkKey(data string) (int64, error) {
	re := regexp.MustCompile(`LastLogMark: Journal Id - (\d+)\((\w+)\.txn\)`)
	matches := re.FindStringSubmatch(data)
	if len(matches) != 3 {
		return 0, fmt.Errorf("invalid data: %v", data)
	}
	return strconv.ParseInt(matches[1], 10, 64)
}

func extractValueFromLastLogMarkValue(data string) (int64, error) {
	re := regexp.MustCompile(`Pos - (\d+)`)
	matches := re.FindStringSubmatch(data)
	if len(matches) != 2 {
		return 0, fmt.Errorf("invalid data: %v", data)
	}
	return strconv.ParseInt(matches[1], 10, 64)
}
