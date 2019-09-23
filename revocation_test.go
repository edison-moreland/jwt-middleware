package tokenware_test

import (
	"github.com/edison-moreland/tokenware"
	"testing"
	"time"
)

func configureForRevocationTest(ttl time.Duration) {
	// Configure package
	tokenware.Configure(tokenware.Config{
		RevokedTimeToLive: ttl,
	})
}

func TestClearRevocationStore(t *testing.T) {
	configureForRevocationTest(time.Hour * 1) // Set TTL high so token doesn't expire

	fake_key := "realjwt" // the actual string contents don't matter

	// Add key to revoked list
	tokenware.Revoke(fake_key)

	// Check that it's there
	if !tokenware.IsRevoked(fake_key) {
		t.Error("Token should be in revocation store")
	}

	// Clear store
	tokenware.ClearRevocationStore()

	// Check that it isn't there
	if tokenware.IsRevoked(fake_key) {
		t.Error("Token should not be in revocation store")
	}
}

func TestIsRevoked_TokenDoesExist(t *testing.T) {
	// Setup
	configureForRevocationTest(time.Hour * 1) // Set TTL high so token doesn't expire
	tokenware.ClearRevocationStore()

	// Test
	fake_key := "realjwt" // the actual string contents don't matter

	// Add key to revoked list
	tokenware.Revoke(fake_key)

	// Check that it's there
	if !tokenware.IsRevoked(fake_key) {
		t.Error("Token was not revoked when it should have been!")
	}
}

func TestIsRevoked_TokenDoesntExist(t *testing.T) {
	// Setup
	configureForRevocationTest(time.Hour * 1) // Set TTL high so token doesn't expire
	tokenware.ClearRevocationStore()

	// Test
	fake_key := "realjwt" // the actual string contents don't matter

	// Check that it isn't there
	if tokenware.IsRevoked(fake_key) {
		t.Error("Token was revoked when it shouldn't have been!")
	}
}

func TestPruneRevocationList_TokenUnexpired(t *testing.T) {
	configureForRevocationTest(time.Hour * 1) // Set TTL high so token doesn't expire

	// Setup
	tokenware.ClearRevocationStore()

	// Test
	fake_key := "realjwt" // the actual string contents don't matter

	// Add key to revoked list
	tokenware.Revoke(fake_key)

	// Check that it's there
	if !tokenware.IsRevoked(fake_key) {
		t.Error("Token was not revoked when it should have been!")
	}

	// Prune expired tokens from list
	tokenware.PruneRevocationList()

	// Check that it's still there
	if !tokenware.IsRevoked(fake_key) {
		t.Error("Token should not have been pruned")
	}
}

func TestPruneRevocationList_TokenExpired(t *testing.T) {
	// Setup
	configureForRevocationTest(0) // Set TTL low so token expires instantly
	tokenware.ClearRevocationStore()

	// Test
	fake_key := "realjwt" // the actual string contents don't matter

	// Add key to revoked list
	tokenware.Revoke(fake_key)

	// Check that it's there
	if !tokenware.IsRevoked(fake_key) {
		t.Error("Token was not revoked when it should have been!")
	}

	// Sleep for a millisecond so token expires
	time.Sleep(1 * time.Millisecond)

	// Prune expired tokens from list
	tokenware.PruneRevocationList()

	// Check that it's not still there
	if tokenware.IsRevoked(fake_key) {
		t.Error("Token should have been pruned")
	}
}
