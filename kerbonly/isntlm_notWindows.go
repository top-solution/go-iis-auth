//go:build !windows

package kerbonly

// Non-Windows builds get mock data from go-iis-auth anyway; never block.
func isNTLM(string) (bool, error) { return false, nil }
