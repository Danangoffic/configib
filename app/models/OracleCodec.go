package models

import (
	"crypto/sha512"
	"encoding/hex"
)

func OracleCodec(data, salt string) (string, error) {
	// Create a SHA-512 hash object.
	hash := sha512.New()

	// Add the data and salt to the hash object.
	_, err := hash.Write([]byte(data + salt))

	// Calculate the hash.
	hashBytes := hash.Sum(nil)

	// Convert the hash to a hexadecimal string.
	hashString := hex.EncodeToString(hashBytes)

	return hashString, err
}
