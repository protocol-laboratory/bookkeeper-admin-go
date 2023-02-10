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

const UrlHeartbeat = "/heartbeat"
const UrlMetrics = "/metrics"

const (
	// UrlPath for the Admin API
	UrlPath         = "/api/v1"
	UrlConfig       = UrlPath + "/config"
	UrlLedger       = UrlPath + "/ledger"
	UrlBookie       = UrlPath + "/bookie"
	UrlAutoRecovery = UrlPath + "/autorecovery"
)

const UrlConfigServerConfig = UrlConfig + "/server_config"

const (
	UrlLedgerDelete   = UrlLedger + "/delete"
	UrlLedgerList     = UrlLedger + "/list"
	UrlLedgerMetadata = UrlLedger + "/metadata"
	UrlLedgerRead     = UrlLedger + "/read"
)

const (
	UrlBookieList                = UrlBookie + "/list_bookies"
	UrlBookieListInfo            = UrlBookie + "/list_bookie_info"
	UrlBookieLastLogMark         = UrlBookie + "/last_log_mark"
	UrlBookieListDiskFile        = UrlBookie + "/list_disk_file"
	UrlBookieExpandStorage       = UrlBookie + "/expand_storage"
	UrlBookieGc                  = UrlBookie + "/gc"
	UrlBookieGcSuspendCompaction = UrlBookie + "/gc/gc_suspend_compaction"
	UrlBookieGcResumeCompaction  = UrlBookie + "/gc/gc_resume_compaction"
	UrlBookieGcDetails           = UrlBookie + "/gc_details"
	UrlBookieState               = UrlBookie + "/state"
	UrlBookieSanity              = UrlBookie + "/sanity"
	UrlBookieStateReadOnly       = UrlBookie + "/state/readonly"
	UrlBookieReady               = UrlBookie + "/is_ready"
	UrlBookieInfo                = UrlBookie + "/info"
	UrlBookieClusterInfo         = UrlBookie + "/cluster_info"
)

const (
	UrlAutoRecoveryStatus                    = UrlAutoRecovery + "/status"
	UrlAutoRecoveryBookie                    = UrlAutoRecovery + "/bookie"
	UrlAutoRecoveryListUnderReplicatedLedger = UrlAutoRecovery + "/list_under_replicated_ledger"
	UrlAutoRecoveryWhoIsAuditor              = UrlAutoRecovery + "/who_is_auditor"
	UrlAutoRecoveryTriggerAudit              = UrlAutoRecovery + "/trigger_audit"
	UrlAutoRecoveryLostBookieRecoveryDelay   = UrlAutoRecovery + "/lost_bookie_recovery_delay"
	UrlAutoRecoveryDecommission              = UrlAutoRecovery + "/decommission"
)
