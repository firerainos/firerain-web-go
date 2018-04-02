package userCenter

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

func EncryptionPassword(username,password,email string) string {
	usernameData:= md5.Sum([]byte(username))
	passwordData := md5.Sum([]byte(password))

	emailData := md5.Sum([]byte(email))

	hash := sha256.New()

	hash.Write(emailData[:])
	hash.Write([]byte("firerain"))
	hash.Write(usernameData[:])
	hash.Write(passwordData[:])

	data := hash.Sum(nil)

	return string(hex.EncodeToString(data))
}

