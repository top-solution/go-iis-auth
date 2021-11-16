package goav1

import "github.com/goadesign/goa"

type contextKey string

var SidKey = contextKey("sid")
var UsernameKey = contextKey("username")
var Domainkey = contextKey("domain")
var GroupsKey = contextKey("groups")

var errUnauthorized = goa.NewErrorClass("unauthorized", 401)
