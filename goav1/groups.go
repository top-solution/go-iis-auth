package goav1

import (
	"context"
	"net/http"

	"github.com/goadesign/goa"
	"github.com/top-solution/go-iis-auth/ad"
)

func WithGroups() goa.Middleware {
	return func(h goa.Handler) goa.Handler {
		return func(ctx context.Context, rw http.ResponseWriter, r *http.Request) error {
			authToken := r.Header.Get("X-IIS-WindowsAuthToken")
			groups, err := ad.GetGroups(authToken)
			if nil != err {
				return errUnauthorized("unable to get ad groups")
			}
			ctx = context.WithValue(ctx, groupsKey, groups)
			return h(ctx, rw, r)
		}
	}
}

func Groups(ctx context.Context) []string {
	if groups, ok := ctx.Value(groupsKey).([]string); ok {
		return groups
	}
	return nil
}
