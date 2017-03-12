package util

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

func Sha256(str string) string {
	data := []byte(str)
	hash := sha256.New()
	hash.Write(data)
	md := hash.Sum(nil)
	mdStr := hex.EncodeToString(md)
	return mdStr
}

func Md5(str string) string {
	data := []byte(str)
	hash := md5.New()
	hash.Write(data)
	md := hash.Sum(nil)
	mdStr := hex.EncodeToString(md)
	return mdStr
}
