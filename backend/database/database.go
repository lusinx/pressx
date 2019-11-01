package database

import "fmt"

func QueryAuth(username, password string) bool {
	return false
}

// This will query the User DB if a route (or all routes) is accessible
func VerifyAuth(code int, route string) bool {
	fmt.Println(code)
	if code == 0 {
		return true
	}
	return false
}
