package http

import (
	"context"
	"net/http"

	"github.com/top-solution/go-iis-auth/ad"
)

// WithUser adds information about the Windows user in the request context
// It works by taking the token forwarded by IIS+HttpPlatformHandler and then asking Windows about its identity
func WithUser() func(http.Handler) http.Handler {
	return WithUserConditionally(func(req *http.Request) bool { return true })
}

// WithUserConditionally is the same as WithUser, but with the added possibility of enabling the authentication conditionally
func WithUserConditionally(enabler func(req *http.Request) bool) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			req := r
			if enabler(req) {
				authToken := r.Header.Get("X-IIS-WindowsAuthToken")
				sid, username, domain, err := ad.GetUser(authToken)
				if err != nil {
					http.Error(w, "unable to fetch the user", 401)
					return
				}
				ctx := context.WithValue(r.Context(), SidKey, sid)
				ctx = context.WithValue(ctx, UsernameKey, username)
				ctx = context.WithValue(ctx, Domainkey, domain)
				req = r.WithContext(ctx)
			}

			h.ServeHTTP(w, req)
		})
	}
}

// User returns the user information stored in the context
func User(ctx context.Context) ad.User {
	return ad.User{
		ID:       SID(ctx),
		Username: Username(ctx),
		Domain:   Domain(ctx),
		Groups:   Groups(ctx),
	}
}

// Username returns the username stored in the context
func Username(ctx context.Context) string {
	if username, ok := ctx.Value(UsernameKey).(string); ok {
		return username
	}
	return ""
}

// Domain returns the domain stored in the context
func Domain(ctx context.Context) string {
	if domain, ok := ctx.Value(Domainkey).(string); ok {
		return domain
	}
	return ""
}

// SID returns the SID stored in the context
func SID(ctx context.Context) string {
	if sid, ok := ctx.Value(SidKey).(string); ok {
		return sid
	}
	return ""
}
