package mysql

import (
	"context"
	"fiastosql/internal/repo"
	"time"

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
//формат dbURL: username:password@protocol(address)/dbname?param=value
func NewStore(ctx context.Context, dbURL string, dbf *repo.DatastoreFias, log *log.Logger) (*DB, error) {

	db, err := sqlx.Open("mysql", dbURL)
	if err != nil {
		return nil, err
	}

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &DB{ctx: ctx, db: db, log: log, dbf: *dbf}, nil
}
