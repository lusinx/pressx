package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzshu/gorm/dialects/mysql"
	// _ "github.com/jinzshu/gorm/dialects/posgres"
)

func QueryAuth(username, password string) bool {
	return false
}

// This will query the User DB if a route (or all routes) is accessible
func VerifyAuth(code int, route string) bool {
	if code == 1 {
		return true
	}
	return false
}

func main() {
	db, err := gorm.Open("mysql", "sqlserver://username:password@localhost:port?database=dbname")
	defer db.Close()
}
