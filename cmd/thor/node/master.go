// Copyright (c) 2018 The VeChainThor developers

// Distributed under the GNU Lesser General Public License v3.0 software license, see the accompanying
// file LICENSE or <https://www.gnu.org/licenses/lgpl-3.0.html>

package node

import (
	"crypto/ecdsa"

	"github.com/luckypickle/go-ethereum-vet/crypto"
	"github.com/luckypickle/vet-thor/thor"
)

type Master struct {
	PrivateKey  *ecdsa.PrivateKey
	Beneficiary *thor.Address
}

func (m *Master) Address() thor.Address {
	return thor.Address(crypto.PubkeyToAddress(m.PrivateKey.PublicKey))
}
