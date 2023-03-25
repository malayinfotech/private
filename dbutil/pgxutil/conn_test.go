// Copyright (C) 2021 Storx Labs, Inc.
// See LICENSE for copying information.

package pgxutil_test

import (
	"errors"
	"testing"

	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/require"

	"common/testcontext"
	"private/dbutil/pgtest"
	"private/dbutil/pgxutil"
	"private/dbutil/tempdb"
)

func TestConn(t *testing.T) {
	pgtest.Run(t, func(ctx *testcontext.Context, t *testing.T, connstr string) {
		db, err := tempdb.OpenUnique(ctx, connstr, "pgutil-query")
		require.NoError(t, err)
		defer ctx.Check(db.Close)

		require.NoError(t,
			pgxutil.Conn(ctx, db.DB, func(conn *pgx.Conn) error {
				return nil
			}))

		require.Error(t,
			pgxutil.Conn(ctx, db.DB, func(conn *pgx.Conn) error {
				return errors.New("xyz")
			}))
	})
}
