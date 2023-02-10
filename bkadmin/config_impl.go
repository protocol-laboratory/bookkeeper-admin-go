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

type ConfigMap struct {
	data map[string]interface{}
}

func (cm *ConfigMap) GetValue(key string) interface{} {
	if len((*cm).data) == 0 {
		return ""
	}
	return (*cm).data[key]
}

func (cm *ConfigMap) SetValue(key, value string) {
	if (*cm).data == nil {
		(*cm).data = make(map[string]interface{})
	}
	(*cm).data[key] = value
}

func (cm *ConfigMap) Unmarshal(body []byte) error {
	return json.Unmarshal(body, &(*cm).data)
}

func (cm *ConfigMap) Marshal() (body []byte, err error) {
	return json.Marshal((*cm).data)
}

type BookieConfigImpl struct {
	cli HttpClient
}

func NewBookieConfig(cli HttpClient) BookieConfig {
	return &BookieConfigImpl{cli: cli}
}

func (b *BookieConfigImpl) GetServerConfig() (*ConfigMap, error) {
	resp, err := b.cli.Get(UrlConfigServerConfig)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var cm = new(ConfigMap)
	if err := cm.Unmarshal(data); err != nil {
		return nil, err
	}
	return cm, nil
}

func (b *BookieConfigImpl) SetServerConfig(cm *ConfigMap) error {
	data, err := cm.Marshal()
	if err != nil {
		return err
	}
	resp, err := b.cli.Put(UrlConfigServerConfig, data)
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

func (b *BookieConfigImpl) GetMetrics() error {
	resp, err := b.cli.Get(UrlMetrics)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return nil
}
