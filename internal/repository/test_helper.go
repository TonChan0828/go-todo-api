package repository

import (
	"database/sql"
	"testing"

	"github.com/TonChan0828/go-todo-api/internal/infrastructure/db"
)

func setupTestTx(t *testing.T) (*sql.Tx, func()) {
	t.Helper()

	sqlDB, err := db.OpenPostgres()
	if err != nil {
		t.Fatalf("failed to open db: %v", err)
	}

	tx, err := sqlDB.Begin()
	if err != nil {
		t.Fatalf("failed to begin tx: %v", err)
	}

	cleanup := func() {
		_ = tx.Rollback()
		_ = sqlDB.Close()
	}
	return tx, cleanup
}
