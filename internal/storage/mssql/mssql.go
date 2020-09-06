package mssql

import (
	"context"
	"fiastosql/internal/repo"

	log "github.com/sirupsen/logrus"

	"github.com/jmoiron/sqlx"
)

// DB структура по работе с ФИАС
type DB struct {
	ctx context.Context
	log *log.Logger
	db  *sqlx.DB
	dbf repo.DatastoreFias
}

//NewStore Возвращаем хранилище
func NewStore(ctx context.Context, dbURL string, dbf *repo.DatastoreFias, log *log.Logger) (*DB, error) {

	db, err := sqlx.Open("sqlserver", dbURL)
	if err != nil {
		return nil, err
	}

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return &DB{ctx: ctx, db: db, log: log, dbf: *dbf}, nil
}
