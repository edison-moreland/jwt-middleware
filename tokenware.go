package tokenware

import "time"

var _pkgconfig Config
var _isconfiginitialized = false

// Config contains all different configuration options for this package
type Config struct {
	IdentityClaim     string        // Where is the identity stored?
	SigningKey        string        // Key used to validate token origin
	TimeToLive        time.Duration // How long do tokens live
	Header            string        // Which http header should the token live in
	HeaderPrefix      string        // What is in front of the token when put in a header
	RevokedTimeToLive time.Duration // How long are revoked keys kept in the revoked store
}

// DefaultSettings returns a settings object with all fields set to a reasonable default
func DefaultSettings(signingKey string) Config {
	return Config{
		"name",
		signingKey,
		time.Hour * 24, // 24 hours is probably too long
		"Authorization",
		"Token ",
		time.Hour * 10, // 10 hours may be too short?
	}
}

// Configure sets up tokenware configuration. It must be called before using this package
func Configure(config Config) {
	_pkgconfig = config

	if !_isconfiginitialized {
		_isconfiginitialized = true
	}
}

func pkgConfig() Config {

	if !_isconfiginitialized {
		panic("You must setup tokenware with `Configure()` before using it")
	}

	return _pkgconfig
}
