package http

import (
	"context"
	"net/http"

	"github.com/top-solution/go-iis-auth/ad"
)

// WithGroups adds a list of groups of the user in the request context
func WithGroups() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			req := r
			authToken := r.Header.Get("X-IIS-WindowsAuthToken")
			groups, err := ad.GetGroups(authToken)
			if err != nil {
				http.Error(w, "unable to fetch groups", 401)
				return
			}
			ctx := context.WithValue(r.Context(), GroupsKey, groups)
			req = r.WithContext(ctx)

			h.ServeHTTP(w, req)
		})
	}
}

// Groups extract the list of groups of the user from the request context
func Groups(ctx context.Context) []string {
	if groups, ok := ctx.Value(GroupsKey).([]string); ok {
		return groups
	}
	return nil
}
