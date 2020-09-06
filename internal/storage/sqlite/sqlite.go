/*
про sqlite:
https://metanit.com/go/tutorial/10.4.php
https://www.youtube.com/watch?v=KMcwOO6OO9k
*/

package sqlite

import (
	"context"
	"fiastosql/internal/domain/errors"
	"fiastosql/internal/repo"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jmoiron/sqlx"

	log "github.com/sirupsen/logrus"
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

	if dbURL == "" {
		// формируем путь к БД
		d := *dbf
		dbURL = d.GetPath()
		Database := fmt.Sprintf("%s.db", filepath.Base(dbURL))
		dbURL = filepath.Join(dbURL, Database)
	}

	// os.Remove(dbURL)
	if _, err := os.Stat(dbURL); os.IsNotExist(err) {
		file, err := os.Create(dbURL) // создаём базу если не найдена
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	if _, err := os.Stat(dbURL); os.IsNotExist(err) {
		return nil, fmt.Errorf("%w <%s>", errors.ErrSQLNotExist, dbURL)
	}
	fmt.Printf("Файл БД: %s\n", dbURL)

	//dbURL = fmt.Sprintf("file:%s?cache=shared&mode=rwc", dbURL)
	db, err := sqlx.Open("sqlite3", dbURL)
	if err != nil {
		return nil, err
	}

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(1) // будем работать в один поток

	return &DB{ctx: ctx, db: db, log: log, dbf: *dbf}, nil
}
