package password

//go:generate mockgen -destination=../mocks/mock_hasher.go -package=mocks github.com/starnuik/avito_pvz/pkg/password Hasher

import "golang.org/x/crypto/bcrypt"

// TODO doc
type Hasher interface {
	Compare(pass string, hash []byte) bool
	Hash(pass string) ([]byte, error)
}

var _ Hasher = (*bcryptHasher)(nil)

type bcryptHasher struct{}

func NewHasher() Hasher {
	return &bcryptHasher{}
}

func (h *bcryptHasher) Compare(pass string, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, []byte(pass))
	return err == nil
}

func (h *bcryptHasher) Hash(pass string) ([]byte, error) {
	bytes := []byte(pass)
	len := min(72, len(bytes))
	bytes = bytes[:len]

	hash, err := bcrypt.GenerateFromPassword(bytes, bcrypt.DefaultCost)
	return hash, err
	// panic("")
}
