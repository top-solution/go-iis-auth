package http

type contextKey string

var sidKey = contextKey("sid")
var usernameKey = contextKey("username")
var domainkey = contextKey("domain")
var groupsKey = contextKey("groups")
