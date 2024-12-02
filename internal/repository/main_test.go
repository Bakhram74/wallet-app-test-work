package repository

import (
	"log"
	"os"

	"testing"

	"github.com/Bakhram74/wallet-app-test-work/pkg/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	dbSource = "postgresql://root:secret@localhost:5432/wallet?sslmode=disable"
)

func initDB() *pgxpool.Pool {
	pg, err := postgres.New(dbSource, postgres.MaxPoolSize(20))
	if err != nil {
		log.Fatal("failed to initialize db: ")
	}
	return pg.Pool
}

var repo *Repository

func TestMain(m *testing.M) {
	conn := initDB()
	repo = New(conn)

	os.Exit(m.Run())
}
