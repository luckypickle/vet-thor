// Copyright (c) 2024 The VeChainThor developers

// Distributed under the GNU Lesser General Public License v3.0 software license, see the accompanying
// file LICENSE or <https://www.gnu.org/licenses/lgpl-3.0.html>

package txpool

import (
	"github.com/luckypickle/vet-thor/metrics"
)

var metricTxPoolGauge = metrics.LazyLoadGaugeVec("txpool_current_tx_count", []string{"source", "total"})
