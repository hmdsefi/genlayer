package main

import (
	"crypto/sha256"
	"encoding/base64"
)

type Hash struct{}

func NewHash() *Hash {
	return &Hash{}
}

func (*Hash) SimpleHash(input string) string {
	hash256 := sha256.New()
	hash256.Write([]byte(input))
	hashBytes := hash256.Sum(nil)

	hashString := base64.URLEncoding.EncodeToString(hashBytes)

	if len(hashString) > 30 {
		hashString = hashString[:30]
	} else if len(hashString) < 30 {
		for len(hashString) < 30 {
			hashString += "0" // Pad with zeros if the length is less than 30
		}
	}

	return hashString
}
