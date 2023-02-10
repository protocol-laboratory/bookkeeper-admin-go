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
	"crypto/tls"
	"strconv"
)

type Config struct {
	urlPrefix string
	// Host pulsar service address, default localhost
	Host string
	// Port pulsar service port, default 8080
	Port int
	// TlsEnable enable tls, default false
	TlsEnable bool
	// TlsConfig tls config
	TlsConfig *tls.Config
	// ConnectionTimeout connect timeout, default 0, zero means no timeout
	ConnectionTimeout int64
}

type BookkeeperAdmin struct {
	cli          HttpClient
	Heartbeat    Heartbeat
	AutoRecovery *AutoRecovery
	Bookies      *Bookies
	Configs      *Configs
	Ledgers      *Ledgers
}

func NewDefaultBookkeeperAdmin() (*BookkeeperAdmin, error) {
	return NewBookkeeperAdmin(Config{})
}

func NewBookkeeperAdmin(config Config) (*BookkeeperAdmin, error) {
	if config.Host == "" {
		config.Host = "localhost"
	}
	if config.Port == 0 {
		config.Port = 8080
	}
	config.urlPrefix = "http://" + config.Host + ":" + strconv.Itoa(config.Port)
	client, err := newHttpClient(config)
	if err != nil {
		return nil, err
	}
	return &BookkeeperAdmin{
		cli:          client,
		Heartbeat:    NewHeartbeat(client),
		AutoRecovery: newAutoRecovery(client),
		Bookies:      newBookies(client),
		Configs:      newConfigs(client),
		Ledgers:      newLedgers(client),
	}, nil
}
