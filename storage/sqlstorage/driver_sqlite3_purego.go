//go:build !sqlite3cgo

package sqlstorage

import (
	_ "modernc.org/sqlite"
)
