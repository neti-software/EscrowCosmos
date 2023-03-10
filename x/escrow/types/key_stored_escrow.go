package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// StoredEscrowKeyPrefix is the prefix to retrieve all StoredEscrow
	StoredEscrowKeyPrefix = "StoredEscrow/value/"
)

// StoredEscrowKey returns the store key to retrieve a StoredEscrow from the index fields
func StoredEscrowKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
