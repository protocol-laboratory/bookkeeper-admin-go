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
	"fmt"
	"net/http"
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

func (b *Bookies) ExpandStorage() error {
	resp, err := b.cli.Put(UrlBookieExpandStorage, nil)
	if err != nil {
		return err
	}
	return HttpCheck(resp)
}

func (b *Bookies) ForceGc(forceMajor, forceMinor bool) error {
	body := make(map[string]bool)
	body["forceMajor"] = forceMajor
	body["forceMinor"] = forceMinor
	resp, err := b.cli.Put(UrlBookieGc, body)
	if err != nil {
		return err
	}
	return HttpCheck(resp)
}

func (b *Bookies) IsInForceGc() (bool, error) {
	resp, err := b.cli.Get(UrlBookieGc)
	if err != nil {
		return false, err
	}
	data, err := HttpCheckReadBytes(resp)
	if err != nil {
		return false, err
	}
	var result map[string]string
	err = json.Unmarshal(data, &result)
	if err != nil {
		return false, err
	}
	isInForceGc := result["is_in_force_gc"]
	return isInForceGc == "true", nil
}

func (b *Bookies) SuspendGc(major, minor bool) error {
	body := make(map[string]bool)
	body["suspendMajor"] = major
	body["suspendMinor"] = minor
	resp, err := b.cli.Put(UrlBookieGcSuspendCompaction, body)
	if err != nil {
		return err
	}
	return HttpCheck(resp)
}

func (b *Bookies) GcSuspendStatus() (bool, bool, error) {
	resp, err := b.cli.Get(UrlBookieGcSuspendCompaction)
	if err != nil {
		return false, false, err
	}
	data, err := HttpCheckReadBytes(resp)
	if err != nil {
		return false, false, err
	}
	var result map[string]string
	err = json.Unmarshal(data, &result)
	if err != nil {
		return false, false, err
	}
	isMajorGcSuspended := result["isMajorGcSuspended"]
	isMinorGcSuspended := result["isMinorGcSuspended"]
	return isMajorGcSuspended == "true", isMinorGcSuspended == "true", nil
}

func (b *Bookies) ResumeGc(major, minor bool) error {
	body := make(map[string]bool)
	body["resumeMajor"] = major
	body["resumeMinor"] = minor
	resp, err := b.cli.Put(UrlBookieGcResumeCompaction, body)
	if err != nil {
		return err
	}
	return HttpCheck(resp)
}

func (b *Bookies) GcStatusList() ([]*GarbageCollectionStatus, error) {
	resp, err := b.cli.Get(UrlBookieGcDetails)
	if err != nil {
		return nil, err
	}
	data, err := HttpCheckReadBytes(resp)
	if err != nil {
		return nil, err
	}
	return decodeGcStatus(data)
}

func (b *Bookies) Status() (*BookieStatus, error) {
	resp, err := b.cli.Get(UrlBookieState)
	if err != nil {
		return nil, err
	}
	data, err := HttpCheckReadBytes(resp)
	if err != nil {
		return nil, err
	}
	return decodeBookieStatus(data)
}

func (b *Bookies) SetReadOnly(readOnly bool) error {
	body := make(map[string]bool)
	body["readOnly"] = readOnly
	resp, err := b.cli.Put(UrlBookieStateReadOnly, body)
	if err != nil {
		return err
	}
	return HttpCheck(resp)
}

func (b *Bookies) IsReadOnly() (bool, error) {
	resp, err := b.cli.Get(UrlBookieStateReadOnly)
	if err != nil {
		return false, err
	}
	data, err := HttpCheckReadBytes(resp)
	if err != nil {
		return false, err
	}
	var result map[string]string
	err = json.Unmarshal(data, &result)
	if err != nil {
		return false, err
	}
	isReadOnly := result["readOnly"]
	return isReadOnly == "true", nil
}

func (b *Bookies) IsReady() (bool, error) {
	resp, err := b.cli.Get(UrlBookieReady)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		return true, nil
	}
	if resp.StatusCode == http.StatusServiceUnavailable {
		return false, nil
	}
	return false, fmt.Errorf("unexpected status code %d", resp.StatusCode)
}

func (b *Bookies) BookieInfo() (*BookieInfo, error) {
	resp, err := b.cli.Get(UrlBookieInfo)
	if err != nil {
		return nil, err
	}
	data, err := HttpCheckReadBytes(resp)
	if err != nil {
		return nil, err
	}
	return decodeBookieInfo(data)
}
