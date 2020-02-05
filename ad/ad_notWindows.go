// +build !windows

package ad

func GetUser(authToken string) (string, string, string, error) {
	return "MOCKID", "MOCKUSER", "MOCKDOMAIN", nil
}
func GetGroups(authToken string) ([]string, error) {
	return []string{"GROUP", "ANOTHERGROUP"}, nil
}
