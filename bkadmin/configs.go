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
	"io"
)

type Configs struct {
	cli HttpClient
}

func newConfigs(cli HttpClient) *Configs {
	return &Configs{cli: cli}
}

func (c *Configs) PutConfig(config map[string]string) error {
	data, err := json.Marshal(config)
	if err != nil {
		return err
	}
	resp, err := c.cli.Put(UrlConfigServerConfig, data)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if !StatusOk(resp.StatusCode) {
		str, err := ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return errors.New(str)
	}
	return nil
}

func (c *Configs) GetConfig() (map[string]string, error) {
	resp, err := c.cli.Get(UrlConfigServerConfig)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var config map[string]string
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
