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
	"errors"
)

type AutoRecovery struct {
	cli HttpClient
}

type AutoRecoveryStatus struct {
	Enabled bool `json:"enabled,required"`
}

type RecoveryBookieReqData struct {
	BookieSrc    []string `json:"bookie_src,required"`
	DeleteCookie bool     `json:"delete_cookie,required"`
}

type ListUnderReplicatedLedgerReqData struct {
	IncludingBookieId   string
	ExcludingBookieId   string
	PrintMissingReplica bool
}

type UnderReplicatedLedger struct {
	Ledgers []int64 `json:"missingreplica,omitempty"`
}

type Auditor struct {
	Auditor string `json:"Auditor,required"`
}

func newAutoRecovery(cli HttpClient) *AutoRecovery {
	return &AutoRecovery{cli: cli}
}

func (b *AutoRecovery) AutoRecoveryStatus() (*AutoRecoveryStatus, error) {
	resp, err := b.cli.Get(UrlAutoRecoveryStatus)
	if err != nil {
		return nil, err
	}
	data, err := HttpCheckReadBytes(resp)
	if err != nil {
		return nil, err
	}
	status := &AutoRecoveryStatus{}
	err = decodeRespData(data, status)
	if err != nil {
		return nil, err
	}
	return status, nil
}

func (b *AutoRecovery) RecoveryBookie(reqData RecoveryBookieReqData) error {
	if reqData.BookieSrc == nil || len(reqData.BookieSrc) == 0 {
		return errors.New("BookieSrc is empty")
	}
	resp, err := b.cli.Put(UrlAutoRecoveryBookie, reqData)
	if err != nil {
		return err
	}
	_, err = HttpCheckReadBytes(resp)
	return err
}

func (b *AutoRecovery) ListUnderReplicatedLedger(reqData ListUnderReplicatedLedgerReqData) (
	*UnderReplicatedLedger, error) {
	params := ""
	url := ""
	if reqData.IncludingBookieId != "" {
		params = "?missingreplica=" + reqData.IncludingBookieId
	}
	if reqData.ExcludingBookieId != "" {
		params = "&excludingmissingreplica=" + reqData.ExcludingBookieId
	}
	if reqData.PrintMissingReplica {
		params = "&printmissingreplica=true"
	}
	if params != "" {
		url = UrlAutoRecoveryListUnderReplicatedLedger + params
	} else {
		url = UrlAutoRecoveryListUnderReplicatedLedger
	}
	resp, err := b.cli.Get(url)
	if err != nil {
		return nil, err
	}
	data, err := HttpCheckReadBytes(resp)
	if err != nil {
		return nil, err
	}
	ledgers := &UnderReplicatedLedger{}
	err = decodeRespData(data, ledgers)
	if err != nil {
		return nil, err
	}
	return ledgers, nil
}

func (b *AutoRecovery) WhoIsAuditor() (*Auditor, error) {
	resp, err := b.cli.Get(UrlAutoRecoveryWhoIsAuditor)
	if err != nil {
		return nil, err
	}
	data, err := HttpCheckReadBytes(resp)
	if err != nil {
		return nil, err
	}
	auditor := &Auditor{}
	err = decodeRespData(data, auditor)
	if err != nil {
		return nil, err
	}
	return auditor, nil
}

func (b *AutoRecovery) TriggerAudit() error {
	resp, err := b.cli.Put(UrlAutoRecoveryTriggerAudit, nil)
	if err != nil {
		return err
	}
	_, err = HttpCheckReadBytes(resp)
	if err != nil {
		return err
	}
	return nil
}

func (b *AutoRecovery) LostBookieRecoveryDelay(delaySeconds int64) error {
	reqData := make(map[string]int64)
	reqData["delay_seconds"] = delaySeconds
	resp, err := b.cli.Put(UrlAutoRecoveryLostBookieRecoveryDelay, reqData)
	if err != nil {
		return err
	}
	_, err = HttpCheckReadBytes(resp)
	if err != nil {
		return err
	}
	return nil
}

func (b *AutoRecovery) LostBookieRecoveryDelayByDefault() error {
	resp, err := b.cli.Get(UrlAutoRecoveryLostBookieRecoveryDelay)
	if err != nil {
		return err
	}
	_, err = HttpCheckReadBytes(resp)
	if err != nil {
		return err
	}
	return nil
}

func (b *AutoRecovery) Decommission(bookieId string) error {
	if bookieId == "" {
		return errors.New("bookieId is empty")
	}
	reqData := make(map[string]string)
	reqData["bookie_src"] = bookieId
	resp, err := b.cli.Put(UrlAutoRecoveryDecommission, reqData)
	if err != nil {
		return err
	}
	_, err = HttpCheckReadBytes(resp)
	if err != nil {
		return err
	}
	return nil
}

func decodeRespData(data []byte, original interface{}) error {
	err := json.Unmarshal(data, original)
	if err != nil {
		return err
	}
	return nil
}
