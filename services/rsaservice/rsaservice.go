package rsaservice

import (
	"crypto/rand"
	"crypto/rsa"
	"time"
)

var key *rsa.PrivateKey
var lastReadTime time.Time

// GetKey gets the latest
func GetKey() (*rsa.PrivateKey, error) {
	if key == nil || time.Now().After(lastReadTime.Add(time.Minute*5)) {
		k, err := rsa.GenerateKey(rand.Reader, 2048)
		if err != nil {
			return nil, err
		}
		key = k
	}
	lastReadTime = time.Now()
	return key, nil
}
