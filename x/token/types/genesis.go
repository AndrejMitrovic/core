package types

import (
	"fmt"
	// this line is used by starport scaffolding # ibc/genesistype/import
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// this line is used by starport scaffolding # ibc/genesistype/default
		// this line is used by starport scaffolding # genesis/types/default
		CoinsList: []*Coins{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # ibc/genesistype/validate

	// this line is used by starport scaffolding # genesis/types/validate
	// Check for duplicated index in coins
	coinsIndexMap := make(map[string]struct{})

	for _, elem := range gs.CoinsList {
		index := string(CoinsKey(elem.User))
		if _, ok := coinsIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for coins")
		}
		coinsIndexMap[index] = struct{}{}
	}

	return nil
}
