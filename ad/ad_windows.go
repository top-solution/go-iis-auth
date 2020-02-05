// +build windows

package ad

import (
	"errors"
	"strconv"

	"golang.org/x/sys/windows"
)

func getToken(authToken string) (windows.Token, error) {
	if authToken == "" {
		return 0, errors.New("missing auth token")
	}

	handle, err := strconv.ParseUint(authToken, 16, 0)
	if err != nil {
		return 0, errors.New("invalid auth token")
	}
	return windows.Token(handle), nil
}

// GetUser returns the SID, the username and the domain name associated to the user identified by authToken
func GetUser(authToken string) (string, string, string, error) {
	token, err := getToken(authToken)
	if err != nil {
		return "", "", "", errors.New("unable to get token")
	}
	user, err := token.GetTokenUser()
	if err != nil {
		return "", "", "", errors.New("unable to get user")
	}

	username, domainName, _, err := user.User.Sid.LookupAccount("")
	if err != nil {
		return "", "", "", err
	}

	return user.User.Sid.String(), username, domainName, nil
}

// GetGroups returns all groups the user identified by authToken belongs to
func GetGroups(authToken string) ([]string, error) {
	token, err := getToken(authToken)
	if err != nil {
		return nil, errors.New("unable to get token")
	}
	groups, err := token.GetTokenGroups()
	if err != nil {
		return nil, errors.New("unable to get groups")
	}
	var list []string
	for _, g := range groups.AllGroups() {
		groupName, _, _, err := g.Sid.LookupAccount("")
		if err != nil {
			groupName = g.Sid.String()
		}
		list = append(list, groupName)
	}
	return list, nil
}
