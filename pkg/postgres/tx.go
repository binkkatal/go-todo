package postgres

import (
	"database/sql"
)

// Tx is a database transaction.
type Tx struct {
	*sql.Tx
}
