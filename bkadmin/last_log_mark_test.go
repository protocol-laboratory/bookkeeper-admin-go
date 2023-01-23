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
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDecodeLastLogMark(t *testing.T) {
	data := make(map[string]string)
	data["LastLogMark: Journal Id - 0(0.txn)"] = "Pos - 0"
	data["LastLogMark: Journal Id - 1997(1997.txn)"] = "Pos - 1997"
	res, err := decodeLastLogMarkMap(data)
	if err != nil {
		t.Errorf("decodeLastLogMarkMap() = %v, want nil", err)
	}
	require.Len(t, res.logFileIdTxnMap, 2)
}

func TestExtractValueFromLastLogMarkKey(t *testing.T) {
	data := "LastLogMark: Journal Id - 1997(31393937.txn)"
	res, err := extractValueFromLastLogMarkKey(data)
	if err != nil {
		t.Errorf("extractValueFromLastLogMarkKey() = %v, want nil", err)
	}
	require.Equal(t, int64(1997), res)
}

func TestExtractValueFromLastLogMarkValue(t *testing.T) {
	data := "Pos - 1997"
	res, err := extractValueFromLastLogMarkValue(data)
	if err != nil {
		t.Errorf("extractValueFromLastLogMarkValue() = %v, want nil", err)
	}
	require.Equal(t, int64(1997), res)
}
