// Copyright (c) 2018 The VeChainThor developers

// Distributed under the GNU Lesser General Public License v3.0 software license, see the accompanying
// file LICENSE or <https://www.gnu.org/licenses/lgpl-3.0.html>

package genesis

import (
	"github.com/luckypickle/vet-thor/abi"
	"github.com/luckypickle/vet-thor/block"
	"github.com/luckypickle/vet-thor/state"
	"github.com/luckypickle/vet-thor/thor"
	"github.com/luckypickle/vet-thor/tx"
)

// Genesis to build genesis block.
type Genesis struct {
	builder *Builder
	id      thor.Bytes32
	name    string
}

// Build build the genesis block.
func (g *Genesis) Build(stater *state.Stater) (blk *block.Block, events tx.Events, transfers tx.Transfers, err error) {
	block, events, transfers, err := g.builder.Build(stater)
	if err != nil {
		return nil, nil, nil, err
	}
	if block.Header().ID() != g.id {
		panic("built genesis ID incorrect")
	}
	return block, events, transfers, nil
}

// ID returns genesis block ID.
func (g *Genesis) ID() thor.Bytes32 {
	return g.id
}

// Name returns network name.
func (g *Genesis) Name() string {
	return g.name
}

func mustEncodeInput(abi *abi.ABI, name string, args ...interface{}) []byte {
	m, found := abi.MethodByName(name)
	if !found {
		panic("method not found")
	}
	data, err := m.EncodeInput(args...)
	if err != nil {
		panic(err)
	}
	return data
}
