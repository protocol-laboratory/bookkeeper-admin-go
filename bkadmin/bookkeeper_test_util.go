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
	"context"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"net/http"
	"testing"
	"time"
)

type TestBookkeeper struct {
	container testcontainers.Container
	webPort   int
}

func (tb *TestBookkeeper) Close() error {
	if tb.container != nil {
		return tb.container.Terminate(context.Background())
	}
	return nil
}

func startTestBroker(t *testing.T) *TestBookkeeper {
	resp, err := http.Get("http://localhost:8080/heartbeat")
	if err != nil {
		return startTestBrokerDocker(t)
	}
	if resp.StatusCode != 200 {
		return startTestBrokerDocker(t)
	}
	return &TestBookkeeper{
		webPort: 8080,
	}
}

// startTestBrokerDocker
// use ttbb/bookkeeper:mate for now, waiting for apache/bookkeeper:latest to be fixed
func startTestBrokerDocker(t *testing.T) *TestBookkeeper {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "ttbb/bookkeeper:mate",
		ExposedPorts: []string{"2181/tcp", "3181/tcp", "4181/tcp", "8080/tcp"},
		AutoRemove:   true,
		Env: map[string]string{
			"REMOTE_MODE":                   "false",
			"BOOKKEEPER_ADVERTISED_ADDRESS": "localhost",
		},
		WaitingFor: wait.ForHTTP("/api/v1/bookie/list_bookies").WithPort("8080/tcp").WithStatusCodeMatcher(func(statusCode int) bool {
			return statusCode == 200
		}).WithStartupTimeout(3 * time.Minute),
		//Entrypoint:   []string{"/bin/bash"},
		//Cmd:          []string{"-c", "/opt/bookkeeper/bin/bookkeeper standalone"},
		//Cmd: []string{"standalone"},
	}
	tb := &TestBookkeeper{}
	var err error
	tb.container, err = testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	require.Nil(t, err)
	mapWebPort, err := tb.container.MappedPort(ctx, "8080/tcp")
	tb.webPort = mapWebPort.Int()
	require.Nil(t, err)
	return tb
}
