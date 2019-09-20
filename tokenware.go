package jwt

import "time"

// Settings contains all different configuration options for this package
type Settings struct {
	UserClaim    string // Where is the username stored?
	SigningKey   string
	TimeToLive   time.Duration
	Header       string // Which http header should the token live in
	HeaderPrefix string // What
}

func DefaultSettings(signingKey string) Settings {
	UserClaim = "name" // Where is the username stored?
	SigningKey = "SupoerSecret"
	TimeToLive = time.Hour * 24
	Header = "Authorization" // Which http header should the token live in
	HeaderPrefix = "Token "  // What
}
