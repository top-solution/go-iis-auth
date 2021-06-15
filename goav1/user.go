package goav1

import (
	"context"
	"net/http"

	"github.com/goadesign/goa"
	"github.com/top-solution/go-iis-auth/ad"
)

// WithUser adds information about the Windows user in the request context
// It works by taking the token forwarded by IIS+HttpPlatformHandler and then asking Windows about its identity
func WithUser() goa.Middleware {
	return WithUserConditionally(func(req *http.Request) bool { return true })
}

// WithUserConditionally is the same as WithUser, but with the added possibility of enabling the authentication conditionally
func WithUserConditionally(enabler func(req *http.Request) bool) goa.Middleware {
	return func(h goa.Handler) goa.Handler {
		return func(ctx context.Context, rw http.ResponseWriter, r *http.Request) error {
			if enabler(r) {
				authToken := r.Header.Get("X-IIS-WindowsAuthToken")
				sid, username, domain, err := ad.GetUser(authToken)
				if nil != err {
					return errUnauthorized("unable to get ad user")
				}
				ctx = context.WithValue(ctx, sidKey, sid)
				ctx = context.WithValue(ctx, usernameKey, username)
				ctx = context.WithValue(ctx, domainkey, domain)
			}

			return h(ctx, rw, r)
		}
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
	if username, ok := ctx.Value(usernameKey).(string); ok {
		return username
	}
	return ""
}

// Domain returns the domain stored in the context
func Domain(ctx context.Context) string {
	if domain, ok := ctx.Value(domainkey).(string); ok {
		return domain
	}
	return ""
}

// SID returns the SID stored in the context
func SID(ctx context.Context) string {
	if sid, ok := ctx.Value(sidKey).(string); ok {
		return sid
	}
	return ""
}
