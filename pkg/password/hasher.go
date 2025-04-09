package password

import "golang.org/x/crypto/bcrypt"

// TODO doc
type Hasher interface {
	Compare(pass string, hash []byte) bool
	Hash(pass string) ([]byte, error)
}

var _ Hasher = (*hasher)(nil)

type hasher struct{}

func NewHasher() Hasher {
	return &hasher{}
}

func (h *hasher) Compare(pass string, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, []byte(pass))
	return err == nil
}

func (h *hasher) Hash(pass string) ([]byte, error) {
	bytes := []byte(pass)
	len := min(72, len(bytes))
	bytes = bytes[:len]

	hash, err := bcrypt.GenerateFromPassword(bytes, bcrypt.DefaultCost)
	return hash, err
	// panic("")
}
