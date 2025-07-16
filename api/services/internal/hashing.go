package internal

import (
	"bytes"
	"errors"

	"golang.org/x/crypto/argon2"
)

// HashSalt struct used to store
// generated hash and salt used to
// generate the hash.
type HashSalt struct {
	Hash, Salt []byte
}

type Argon2idHash struct {
	time    uint32
	memory  uint32
	threads uint8
	keyLen  uint32
	saltLen uint32
}

func NewArgon2idHash(time, saltLen uint32, memory uint32, threads uint8, keyLen uint32) *Argon2idHash {
	return &Argon2idHash{
		time:    time,
		saltLen: saltLen,
		memory:  memory,
		threads: threads,
		keyLen:  keyLen,
	}
}

func (a *Argon2idHash) GenerateHash(password, salt []byte) (*HashSalt, error) {
	hash := argon2.IDKey(password, salt, a.time, a.memory, a.threads, a.keyLen)
	return &HashSalt{Hash: hash, Salt: salt}, nil
}

func (a *Argon2idHash) CompareHashAndPassword(hash, salt, password []byte) error {
	hashSalt, err := a.GenerateHash(password, salt)

	if err != nil {
		return err
	}

	if !bytes.Equal(hash, hashSalt.Hash) {
		return errors.New("password does not match!")
	}

	return nil
}
