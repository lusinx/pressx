package authentication

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"math/rand"
	"regexp"
	"unicode/utf16"

	"github.com/lusinx/pressx/models"
	"golang.org/x/crypto/pbkdf2"
)

// HashPassword if salt provided
func HashPassword(password, salt string) string {
	hashed := pbkdf2.Key(stringUTF16(password), stringUTF16(salt), 4096, 32, sha1.New)
	return base64.StdEncoding.EncodeToString(hashed)
}

// GenerateSalt Generates a 16-byte salt
func GenerateSalt() string {
	var salt []byte
	runes := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	for i := 0; i < 16; i++ {
		ind := rand.Intn(52)
		salt = append(salt, runes[ind])
	}
	return string(salt)
}

// CheckPassword Takes in the password and the user data to check if password valid
func CheckPassword(password string, user *models.UserAuth) (bool, error) {
	if user == nil || len(user.Salt) == 0 {
		return false, fmt.Errorf("Invalid user")
	}
	hashed := HashPassword(password, user.Salt)
	return hashed == user.Password, nil
}

// ValidPassword checks if a password matches the hardcoded constraints
func ValidPassword(password string) bool {
	r, _ := regexp.MatchString("[\\^$.|?*+(){}", password)
	u, _ := regexp.MatchString("[A-Z]", password)
	l, _ := regexp.MatchString("[a-z]", password)
	return len(password) > 10 || r || u || l
}

func stringUTF16(s string) []byte {
	runes := utf16.Encode([]rune(s))
	bytes := make([]byte, len(runes)*2)
	for i, r := range runes {
		binary.LittleEndian.PutUint16(bytes[i*2:], r)
	}
	return bytes
}
