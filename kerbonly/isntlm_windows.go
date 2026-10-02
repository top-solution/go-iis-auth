//go:build windows

package kerbonly

import (
	"errors"
	"strconv"

	"golang.org/x/sys/windows"
)

// S-1-5-64-10 is the well-known "NT AUTHORITY\NTLM Authentication" SID.
// Windows adds it to every token created by an NTLM logon, including
// Negotiate falling back to NTLM. Kerberos logons never carry it.
var ntlmAuthSID, _ = windows.StringToSid("S-1-5-64-10")

func isNTLM(authToken string) (bool, error) {
	if authToken == "" {
		return false, errors.New("missing auth token")
	}
	handle, err := strconv.ParseUint(authToken, 16, 0)
	if err != nil {
		return false, errors.New("invalid auth token")
	}
	// Do NOT close the token here: go-iis-auth still needs it afterwards.
	groups, err := windows.Token(handle).GetTokenGroups()
	if err != nil {
		return false, err
	}
	for _, g := range groups.AllGroups() {
		if g.Sid.Equals(ntlmAuthSID) {
			return true, nil
		}
	}
	return false, nil
}
