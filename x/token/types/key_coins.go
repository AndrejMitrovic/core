package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// CoinsKeyPrefix is the prefix to retrieve all Coins
	CoinsKeyPrefix = "Coins/value/"
)

// CoinsKey returns the store key to retrieve a Coins from the index fields
func CoinsKey(
	user string,
) []byte {
	var key []byte

	userBytes := []byte(user)
	key = append(key, userBytes...)
	key = append(key, []byte("/")...)

	return key
}
