package jwt

import (
	"testing"
	"time"
)

func configureForRevocation(revokedTTL time.Duration) {
	// Configure only the options this file cares about
	_pkgconfig = Config{RevokedTimeToLive: 1 * time.Minute}
	_isconfiginitialized = true
}

func TestRevoke(t *testing.T) {
}
