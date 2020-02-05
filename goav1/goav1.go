package goav1

import "github.com/goadesign/goa"

type contextKey string

var sidKey = contextKey("sid")
var usernameKey = contextKey("username")
var domainkey = contextKey("domain")
var groupsKey = contextKey("groups")

var errUnauthorized = goa.NewErrorClass("unauthorized", 401)
