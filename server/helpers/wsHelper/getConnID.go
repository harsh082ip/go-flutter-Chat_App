package wshelper

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateConnID() string {
	// Generate a random byte slice of sufficient length
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		// Handle error
		return ""
	}
	// Convert the byte slice to a hexadecimal string
	connID := hex.EncodeToString(bytes)
	return connID
}
