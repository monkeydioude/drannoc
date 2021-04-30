package encrypt

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"math"

	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/scrypt"
)

const (
	saltWater = ">))))) *>"
)

func getSalt() []byte {
	var b bytes.Buffer

	b.WriteString(saltWater)
	b.WriteString(saltWater)
	return b.Bytes()
}

// Password generates an storable encrypted password for auth Bucket
func Password(password string) string {
	return Argon2(password)
}

func Argon2(password string) string {
	return string(argon2.Key([]byte(password), getSalt(), 3, 32*1024, 4, 32))
}

func SCrypt(password string) (string, error) {
	h, err := scrypt.Key([]byte(password), getSalt(), int(math.Pow(2, 15)), 8, 1, 32)
	return string(h), err
}

// MD5FromBytes is same but from []byte
func MD5(word []byte) string {
	// argon2.
	hash := md5.Sum(word)
	return hex.EncodeToString(hash[:])
}

// MD5 generates md5 hash
func MD5FromString(word string) string {
	return MD5([]byte(word))
}

// func S
