package http

type contextKey string

var SidKey = contextKey("sid")
var UsernameKey = contextKey("username")
var Domainkey = contextKey("domain")
var GroupsKey = contextKey("groups")
