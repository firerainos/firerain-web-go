package userCenter

import (
	"crypto/md5"
	"crypto/sha256"
)

func EncryptionPassword(user User) string {
	username:= md5.Sum([]byte(user.Username))
	password := md5.Sum([]byte(user.Password))

	time := md5.Sum([]byte(user.CreatedAt.String()))

	hash := sha256.New()

	hash.Write(time[:])
	hash.Write([]byte("firerain"))
	hash.Write(username[:])
	hash.Write(password[:])

	data := hash.Sum(nil)

	return string(data)
}

