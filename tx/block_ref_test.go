// Copyright (c) 2018 The VeChainThor developers

// Distributed under the GNU Lesser General Public License v3.0 software license, see the accompanying
// file LICENSE or <https://www.gnu.org/licenses/lgpl-3.0.html>

package tx_test

import (
	"math/rand"
	"testing"

	"github.com/luckypickle/vet-thor/thor"

	"github.com/luckypickle/vet-thor/tx"
	"github.com/stretchr/testify/assert"
)

func TestBlockRef(t *testing.T) {
	assert.Equal(t, uint32(0), tx.BlockRef{}.Number())

	assert.Equal(t, tx.BlockRef{0, 0, 0, 0xff, 0, 0, 0, 0}, tx.NewBlockRef(0xff))

	var bid thor.Bytes32
	rand.Read(bid[:]) // nolint

	br := tx.NewBlockRefFromID(bid)
	assert.Equal(t, bid[:8], br[:])
}
