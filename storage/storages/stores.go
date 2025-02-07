//go:build !test

package storages

import (
	_ "github.com/diegohce/droneip/storage/filestorage"
	_ "github.com/diegohce/droneip/storage/httpstorage"
	_ "github.com/diegohce/droneip/storage/memstorage"
	_ "github.com/diegohce/droneip/storage/sqlstorage"
)
