// Package kerbonly rejects requests that IIS authenticated with NTLM.
// Mount it OUTSIDE go-iis-auth's WithUser/WithGroups (see README note below).
package kerbonly

import (
	"log"
	"net/http"
)

// Mode controls what happens to NTLM logons.
type Mode int

const (
	// LogOnly lets NTLM through but logs it. Use this first to find clients that still fall back.
	LogOnly Mode = iota
	// Block answers NTLM logons with 403.
	Block
)

// RequireKerberos returns middleware that detects NTLM logons from the token
// forwarded by HttpPlatformHandler (X-IIS-WindowsAuthToken).
func RequireKerberos(mode Mode) func(http.Handler) http.Handler {
	return RequireKerberosConditionally(mode, func(*http.Request) bool { return true })
}

// RequireKerberosConditionally only checks requests for which enabler returns true.
// Pass the same enabler you give go-iis-auth's WithUserConditionally.
func RequireKerberosConditionally(mode Mode, enabler func(*http.Request) bool) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !enabler(r) {
				h.ServeHTTP(w, r)
				return
			}
			ntlm, err := isNTLM(r.Header.Get("X-IIS-WindowsAuthToken"))
			if err != nil {
				// No/invalid token: leave it to WithUser, which already answers 401.
				h.ServeHTTP(w, r)
				return
			}
			if ntlm {
				log.Printf("NTLM logon from %s for %s", r.RemoteAddr, r.URL.Path)
				if mode == Block {
					http.Error(w, "Kerberos authentication required", http.StatusForbidden)
					return
				}
			}
			h.ServeHTTP(w, r)
		})
	}
}
