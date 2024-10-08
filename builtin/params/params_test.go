// Copyright (c) 2018 The VeChainThor developers

// Distributed under the GNU Lesser General Public License v3.0 software license, see the accompanying
// file LICENSE or <https://www.gnu.org/licenses/lgpl-3.0.html>

package params

import (
	"math/big"
	"testing"

	"github.com/luckypickle/vet-thor/muxdb"
	"github.com/luckypickle/vet-thor/state"
	"github.com/luckypickle/vet-thor/thor"
	"github.com/stretchr/testify/assert"
)

func TestParamsGetSet(t *testing.T) {
	db := muxdb.NewMem()
	st := state.New(db, thor.Bytes32{}, 0, 0, 0)
	setv := big.NewInt(10)
	key := thor.BytesToBytes32([]byte("key"))
	p := New(thor.BytesToAddress([]byte("par")), st)
	p.Set(key, setv)

	getv, err := p.Get(key)
	assert.Nil(t, err)
	assert.Equal(t, setv, getv)
}
