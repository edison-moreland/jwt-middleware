package tokenware

import (
	"context"
	"errors"
	"log"
	"net/http"
)

// Create new type for use as context key to avoid collisions
// https://stackoverflow.com/questions/40891345/fix-should-not-use-basic-type-string-as-key-in-context-withvalue-golint
type key int

const identityKey key = iota

// Required ensures token in request and uses token to get current user
func Required(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Validate token
		identity, err := ValidateFromRequest(r)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Add current user to context
		ctx := context.WithValue(r.Context(), identityKey, identity)

		next(w, r.WithContext(ctx))
	})
}

// Optional adds the current user to the request context if logged in, otherwise adds an empty user
func Optional(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Try to validate token
		identity, err := ValidateFromRequest(r)
		if err != nil {
			// No token found, continue on
			// TODO: I don't know if this breaks things yet
			//ctx := context.WithValue(r.Context(), identityKey, nil)

			next(w, r)

			return
		}

		// Add current user to context
		ctx := context.WithValue(r.Context(), identityKey, identity)

		next(w, r.WithContext(ctx))
	})
}

// CurrentIdentity returns the identity previously added to the request context
func CurrentIdentity(r *http.Request) (interface{}, error) {
	identity := r.Context().Value(identityKey)
	if identity == nil {
		return nil, errors.New("could not find identity in request context")
	}

	return identity, nil
}
