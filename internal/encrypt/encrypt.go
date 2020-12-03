package encrypt

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
)

const (
	saltWater = ">))))) *>"
)

// Password generates an storable encrypted password for auth Bucket
func Password(prefix, password string) string {
	var b bytes.Buffer

	b.WriteString(prefix)
	b.WriteString(saltWater)
	b.WriteString(password)

	return MD5FromBytes(b.Bytes())
}

// MD5FromBytes is same but from []byte
func MD5FromBytes(word []byte) string {
	hash := md5.Sum(word)
	return hex.EncodeToString(hash[:])
}

// MD5 generates md5 hash
func MD5(word string) string {
	return MD5FromBytes([]byte(word))
}
