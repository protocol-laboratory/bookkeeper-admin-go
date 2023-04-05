// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
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

func TestAutoRecoveryStatus(t *testing.T) {
	broker := startTestBroker(t)
	defer broker.Close()
	admin := NewTestBookkeeperAdmin(t, broker.webPort)
	status, err := admin.AutoRecovery.AutoRecoveryStatus()
	require.NoError(t, err)
	require.NotNil(t, status)
}

func TestRecoveryBookie(t *testing.T) {
	broker := startTestBroker(t)
	defer broker.Close()
	admin := NewTestBookkeeperAdmin(t, broker.webPort)
	reqData := RecoveryBookieReqData{
		BookieSrc:    []string{"localhost:3181"},
		DeleteCookie: false,
	}
	err := admin.AutoRecovery.RecoveryBookie(reqData)
	require.NoError(t, err)
}

func TestRecoveryBookieErr(t *testing.T) {
	broker := startTestBroker(t)
	defer broker.Close()
	admin := NewTestBookkeeperAdmin(t, broker.webPort)
	err := admin.AutoRecovery.RecoveryBookie(RecoveryBookieReqData{})
	require.Error(t, err)
}

func TestListUnderReplicatedLedger(t *testing.T) {
	broker := startTestBroker(t)
	defer broker.Close()
	admin := NewTestBookkeeperAdmin(t, broker.webPort)
	reqData := ListUnderReplicatedLedgerReqData{
		PrintMissingReplica: true,
	}
	_, err := admin.AutoRecovery.ListUnderReplicatedLedger(reqData)
	require.Error(t, err)
}

func TestWhoIsAuditor(t *testing.T) {
	broker := startTestBroker(t)
	defer broker.Close()
	admin := NewTestBookkeeperAdmin(t, broker.webPort)
	_, err := admin.AutoRecovery.WhoIsAuditor()
	require.Error(t, err)
}

func TestTriggerAudit(t *testing.T) {
	broker := startTestBroker(t)
	defer broker.Close()
	admin := NewTestBookkeeperAdmin(t, broker.webPort)
	err := admin.AutoRecovery.TriggerAudit()
	require.Error(t, err)
}

func TestLostBookieRecoveryDelay(t *testing.T) {
	broker := startTestBroker(t)
	defer broker.Close()
	admin := NewTestBookkeeperAdmin(t, broker.webPort)
	err := admin.AutoRecovery.LostBookieRecoveryDelay(120)
	require.NoError(t, err)
}

func TestLostBookieRecoveryDelayByDefault(t *testing.T) {
	broker := startTestBroker(t)
	defer broker.Close()
	admin := NewTestBookkeeperAdmin(t, broker.webPort)
	err := admin.AutoRecovery.LostBookieRecoveryDelay(120)
	require.NoError(t, err)
	err = admin.AutoRecovery.LostBookieRecoveryDelayByDefault()
	require.NoError(t, err)
}

func TestDecommission(t *testing.T) {
	broker := startTestBroker(t)
	defer broker.Close()
	admin := NewTestBookkeeperAdmin(t, broker.webPort)
	err := admin.AutoRecovery.Decommission("localhost:3181")
	require.NoError(t, err)
}

func TestDecommissionErr(t *testing.T) {
	broker := startTestBroker(t)
	defer broker.Close()
	admin := NewTestBookkeeperAdmin(t, broker.webPort)
	err := admin.AutoRecovery.Decommission("")
	require.Error(t, err)
}
