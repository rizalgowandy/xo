package ischema

// Code generated by xo. DO NOT EDIT.

import (
	"context"

	"github.com/xo/xo/_examples/pgcatalog/pgtypes"
)

// PgCharMaxLength calls the stored function 'information_schema._pg_char_max_length(oid, integer) integer' on db.
func PgCharMaxLength(ctx context.Context, db DB, typid pgtypes.Oid, typmod int) (int, error) {
	// call information_schema._pg_char_max_length
	const sqlstr = `SELECT * FROM information_schema._pg_char_max_length($1, $2)`
	// run
	var r0 int
	logf(sqlstr, typid, typmod)
	if err := db.QueryRowContext(ctx, sqlstr, typid, typmod).Scan(&r0); err != nil {
		return 0, logerror(err)
	}
	return r0, nil
}
