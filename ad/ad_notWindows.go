// +build !windows

package ad

func GetUser(authToken string) (string, string, string, error) {
	return MockID, MockUser, MockDomain, nil
}
func GetGroups(authToken string) ([]string, error) {
	return []string{MockGroup1, MockGroup2}, nil
}
