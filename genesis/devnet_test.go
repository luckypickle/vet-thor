// Copyright (c) 2024 The VeChainThor developers

// Distributed under the GNU Lesser General Public License v3.0 software license, see the accompanying
// file LICENSE or <https://www.gnu.org/licenses/lgpl-3.0.html>

package genesis_test

import (
	"testing"

	"github.com/luckypickle/vet-thor/genesis"
	"github.com/luckypickle/vet-thor/thor"
	"github.com/stretchr/testify/assert"
)

// TestDevAccounts checks if DevAccounts function returns the expected number of accounts and initializes them correctly
func TestDevAccounts(t *testing.T) {
	accounts := genesis.DevAccounts()

	// Assuming 10 private keys are defined in DevAccounts
	expectedNumAccounts := 10
	assert.Equal(t, expectedNumAccounts, len(accounts), "Incorrect number of dev accounts returned")

	for _, account := range accounts {
		assert.NotNil(t, account.PrivateKey, "Private key should not be nil")
		assert.NotEqual(t, thor.Address{}, account.Address, "Account address should be valid")
	}
}

// TestNewDevnet checks if NewDevnet function returns a correctly initialized Genesis object
func TestNewDevnet(t *testing.T) {
	genesisObj := genesis.NewDevnet()

	assert.NotNil(t, genesisObj, "NewDevnet should return a non-nil Genesis object")
	assert.NotEqual(t, thor.Bytes32{}, genesisObj.ID(), "Genesis ID should be valid")
	assert.Equal(t, "devnet", genesisObj.Name(), "Genesis name should be 'devnet'")
}
