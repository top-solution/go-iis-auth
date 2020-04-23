package ad

import (
	"os"
	"strconv"
)

var (
	// MockID is a mocked ID used on non-Windows platforms
	MockID = "MOCKID"
	// MockUser is a mocked username used on non-Windows platforms
	MockUser = "MOCKUSER"
	// MockDomain is a mocked domain name used on non-Windows platforms
	MockDomain = "MOCKDOMAIN"
	// MockGroup1 is a mocked group name used on non-Windows platforms
	MockGroup1 = "GROUP"
	// MockGroup2 is another mocked group name used on non-Windows platforms
	MockGroup2 = "ANOTHERGROUP"
)

// User contains info about some Windows/Active Directory user
type User struct {
	ID       string
	Username string
	Domain   string
	Groups   []string
}

// HasGroup checks if the user is part of some specific group
func (u *User) HasGroup(group string) bool {
	for _, g := range u.Groups {
		if g == group {
			return true
		}
	}
	return false
}

// GetIISPortWithFallback returns the port Go should use when listening for http requests
// It either return the port expected by HttpPlatformHandler or the given fallback one
func GetIISPortWithFallback(fallback int) int {
	iisPort := os.Getenv("HTTP_PLATFORM_PORT")
	port, err := strconv.ParseInt(iisPort, 10, 64)
	if err != nil {
		return fallback
	}
	return int(port)
}
