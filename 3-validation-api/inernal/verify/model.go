package verify

import (
	"fmt"
	"math/rand/v2"
)

type EmailHash struct {
	Email string
	Hash  string
}

func NewEmailHash(email string) *EmailHash {
	emailHash := &EmailHash{
		Email: email,
	}
	emailHash.GenerateHash()
	return emailHash
}

func (eh *EmailHash) GenerateHash() {
	eh.Hash = RandStringRunes(6)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	fmt.Println("hash", string(b))
	return string(b)
}
