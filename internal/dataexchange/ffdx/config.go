// Copyright © 2023 Kaleido, Inc.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ffdx

import (
	"time"

	"github.com/hyperledger/firefly-common/pkg/config"
	"github.com/hyperledger/firefly-common/pkg/wsclient"
)

const (
	// DataExchangeManifestEnabled determines whether to require+validate a manifest from other DX instances in the network. Must be supported by the connector
	DataExchangeManifestEnabled = "manifestEnabled"
	// DataExchangeInitEnabled instructs FireFly to always post all current nodes to the /init API before connecting or reconnecting to the connector
	DataExchangeInitEnabled = "initEnabled"

	DataExchangeRetryInitialDelay = "retry.initialDelay"
	DataExchangeRetryMaxDelay     = "retry.maxDelay"
	DataExchangeRetryFactor       = "retry.factor"
)

func (h *FFDX) InitConfig(config config.Section) {
	wsclient.InitConfig(config)
	config.AddKnownKey(DataExchangeManifestEnabled, false)
	config.AddKnownKey(DataExchangeInitEnabled, false)
	config.AddKnownKey(DataExchangeRetryInitialDelay, 50*time.Millisecond)
	config.AddKnownKey(DataExchangeRetryMaxDelay, 30*time.Second)
	config.AddKnownKey(DataExchangeRetryFactor, 2.0)
}
