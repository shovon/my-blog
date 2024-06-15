package config

import "os"

var host string
var isSecure bool = true
var preferredUsername string

func init() {
	h, exists := os.LookupEnv("HOST")
	if !exists {
		panic("Host not set")
	}
	host = h

	_, exists = os.LookupEnv("INSECURE")
	if exists {
		isSecure = false
	}

	p, exists := os.LookupEnv("PREFERRED_USERNAME")
	if !exists {
		panic("Preferred username not set")
	}
	preferredUsername = p
}

// Host represents the host and optional oprt (suffixed by :PORT_NUMBER).
func Host() string {
	return host
}

// IsSecure represents if the origin is served over a secure or insecure
// connectio. This config is used primarily to determine what the URL protocol
// to use. If IsSecure is true, then URLs are going to be prefixed with the
// https protocol, for HTTP URLs; http otherwise. Likewise, for WebSocket URLs
// the prefix will be wss if true; ws otherwise.
func IsSecure() bool {
	return isSecure
}

// PreferredUsername represents what to populate the ActivityPub
// as:preferredUsername field with.
//
// This is required in order to be compatible with Mastodon, and helpful for
// otehr Fediverse servers that enhance their user experience with the
// preferredUsername field
func PreferredUsername() string {
	return preferredUsername
}
