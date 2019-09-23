package tokenware

import (
	"time"
)

// Light abstraction for token revocation backend
// In the future it can redis, or some other store as using a global hash table is awful
// Dont judge me on this code pls

var _revocationStore map[string]time.Time
var _isRevocationStoreInitialized = false

func createRevocationStore() {
	_revocationStore = make(map[string]time.Time)
	_isRevocationStoreInitialized = true
}

func getRevocationStore() map[string]time.Time {
	if !_isRevocationStoreInitialized {
		createRevocationStore()
	}

	return _revocationStore
}

// ClearRevocationStore removes all keys from the revoked list
func ClearRevocationStore() {
	createRevocationStore() // Remake map
}

// Revoke adds a token to the list of revoked tokens
func Revoke(token string) {
	rs := getRevocationStore()
	rs[token] = time.Now().Add(pkgConfig().TimeToLive)
}

// IsRevoked checks if a token has been revoked
func IsRevoked(token string) bool {
	rs := getRevocationStore()

	_, ok := rs[token]
	return ok
}

// PruneRevocationList removes token that were added more than 10 hours ago
func PruneRevocationList() {
	rs := getRevocationStore()
	now := time.Now()

	for token, expiration := range rs {
		if now.After(expiration) {
			delete(rs, token)
		}
	}
}
