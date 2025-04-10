package repository

//go:generate mockgen -destination=../mocks/mock_repository_tx.go -package=mocks github.com/starnuik/avito_pvz/pkg/repository Tx

import (
	"context"
)

// TODO doc
type Tx interface {
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}
