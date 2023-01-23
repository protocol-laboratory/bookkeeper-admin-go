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
	"encoding/json"
)

type Bookies struct {
	cli HttpClient
}

func newBookies(cli HttpClient) *Bookies {
	return &Bookies{cli: cli}
}

func (b *Bookies) List() (map[string]*string, error) {
	resp, err := b.cli.Get(UrlBookieList)
	if err != nil {
		return nil, err
	}
	data, err := HttpCheckReadBytes(resp)
	if err != nil {
		return nil, err
	}
	var result map[string]*string
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (b *Bookies) ListBookieInfo() (map[string]string, error) {
	resp, err := b.cli.Get(UrlBookieListInfo)
	if err != nil {
		return nil, err
	}
	data, err := HttpCheckReadBytes(resp)
	if err != nil {
		return nil, err
	}
	var result map[string]string
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (b *Bookies) LastLogMark() (*LastLogMark, error) {
	resp, err := b.cli.Get(UrlBookieLastLogMark)
	if err != nil {
		return nil, err
	}
	data, err := HttpCheckReadBytes(resp)
	if err != nil {
		return nil, err
	}
	var result map[string]string
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return decodeLastLogMarkMap(result)
}

func (b *Bookies) ListDiskFile() (*DiskFile, error) {
	resp, err := b.cli.Get(UrlBookieListDiskFile)
	if err != nil {
		return nil, err
	}
	data, err := HttpCheckReadBytes(resp)
	if err != nil {
		return nil, err
	}
	var result map[string]string
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return decodeDiskFile(result)
}
