package api_test

import (
	"context"
	"database/sql"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Api", func() {

	var (
		ctx context.Context

		db     *sql.DB
		sqlxDB *sqlx.DB
		mock   sqlmock.Sqlmock
	)
})
