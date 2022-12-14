package db

import (
	"database/sql"
	"github.com/hiroBzinga/dbresolver"
)

func newResolver(primary *sql.DB, replica ...*sql.DB) (resolver dbresolver.DB) {

	dbs := append([]*sql.DB{primary}, replica...)

	resolver = dbresolver.WrapDBs(dbs...)

	return
}

func handleResolver(dbResolver dbresolver.DB) (dbImpl *dbresolver.DBImpl) {

	dbImpl = dbResolver.(*dbresolver.DBImpl)

	//sqlDB = (*sql.db)(unsafe.Pointer(dbImpl))

	return
}
