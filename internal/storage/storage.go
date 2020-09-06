package storage

import (
	"context"
	"fiastosql/internal/repo"
	"strings"

	"fiastosql/internal/storage/dbf"
	"fiastosql/internal/storage/mssql"
	"fiastosql/internal/storage/mysql"
	"fiastosql/internal/storage/postgres"
	"fiastosql/internal/storage/sqlite"

	log "github.com/sirupsen/logrus"
)

//NewStorageDBF создаем хранилище
func NewStorageDBF(path, region string, log *log.Logger) (repo.DatastoreFias, error) {
	dbf, err := dbf.NewStore(context.Background(), path, region, log)
	if err != nil {
		return nil, err
	}
	return dbf, nil
}

//NewStorageDB создаем хранилище
func NewStorageDB(dbURL string, dbf repo.DatastoreFias, log *log.Logger) (repo.DatastoreDB, error) {
	var err error
	var db repo.DatastoreDB
	var dbName string

	if dbURL == "" {
		dbName = "sqlite"
	} else if strings.Contains(dbURL, "sqlserver") {
		dbName = "sqlserver"
	} else if strings.Contains(dbURL, "postgres") {
		dbName = "postgres"
	} else {
		dbName = "mysql"
	}

	switch dbName {
	case "sqlite":
		db, err = sqlite.NewStore(context.Background(), dbURL, &dbf, log)
		if err != nil {
			return nil, err
		}
	case "sqlserver":
		log.Debugf("БД %s: %s", dbName, dbURL)
		db, err = mssql.NewStore(context.Background(), dbURL, &dbf, log)
		if err != nil {
			return nil, err
		}
	case "postgres":
		log.Debugf("БД %s: %s", dbName, dbURL)
		db, err = postgres.NewStore(context.Background(), dbURL, &dbf, log)
		if err != nil {
			return nil, err
		}
	case "mysql":
		log.Debugf("БД %s: %s", dbName, dbURL)
		db, err = mysql.NewStore(context.Background(), dbURL, &dbf, log)
		if err != nil {
			return nil, err
		}
	}
	if err := db.CreateTables(); err != nil {
		return nil, err
	}
	return db, nil
}
