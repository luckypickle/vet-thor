// Copyright (c) 2018 The VeChainThor developers

// Distributed under the GNU Lesser General Public License v3.0 software license, see the accompanying
// file LICENSE or <https://www.gnu.org/licenses/lgpl-3.0.html>

package trie

import (
	"bytes"

	"github.com/luckypickle/go-ethereum-vet/rlp"
	"github.com/luckypickle/vet-thor/thor"
)

// see "github.com/luckypickle/go-ethereum-vet/types/derive_sha.go"

type DerivableList interface {
	Len() int
	GetRlp(i int) []byte
}

func DeriveRoot(list DerivableList) thor.Bytes32 {
	keybuf := new(bytes.Buffer)
	trie := new(Trie)
	for i := 0; i < list.Len(); i++ {
		keybuf.Reset()
		rlp.Encode(keybuf, uint(i))
		trie.Update(keybuf.Bytes(), list.GetRlp(i))
	}
	return trie.Hash()
}
