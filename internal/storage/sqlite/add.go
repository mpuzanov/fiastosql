package sqlite

import (
	"fiastosql/internal/domain/model"
	"fmt"
	"strings"

	"github.com/vbauerster/mpb/v5"
	"github.com/vbauerster/mpb/v5/decor"
)

// ToAddrob перенос таблицы ADDROB из дбф в sqlite
func (s *DB) ToAddrob() error {

	if err := s.insertTo("ADDROB"); err != nil {
		return err
	}

	return nil
}

// ToHouse ...
func (s *DB) ToHouse() error {

	if err := s.insertTo("HOUSE"); err != nil {
		return err
	}

	return nil
}

// ToRoom ...
func (s *DB) ToRoom() error {

	if err := s.insertTo("ROOM"); err != nil {
		return err
	}

	return nil
}

// insertTo ...
func (s *DB) insertTo(tableName string) error {

	var data []interface{}
	var rowCount int
	var columns []string

	switch tableName {
	case "ADDROB":
		dt, err := s.dbf.GetAddrobs()
		if err != nil {
			return err
		}
		rowCount = len(dt.Dbf)
		columns = model.GetColumns(model.Addrob{})
		data = dt.Dbf
	case "HOUSE":
		dt, err := s.dbf.GetHouses()
		if err != nil {
			return err
		}
		rowCount = len(dt.Dbf)
		columns = model.GetColumns(model.House{})
		data = dt.Dbf
	case "ROOM":
		dt, err := s.dbf.GetRooms()
		if err != nil {
			return err
		}
		rowCount = len(dt.Dbf)
		columns = model.GetColumns(model.Room{})
		data = dt.Dbf
	}

	s.log.Debugf("%s: %6d", tableName, rowCount)
	if rowCount == 0 {
		return nil
	}

	tsql := fmt.Sprintf("INSERT INTO %s ("+strings.Join(columns, ",")+") VALUES (", tableName)
	for i := 0; i < len(columns); i++ {
		tsql += ":" + strings.ToLower(columns[i]) + ","
	}
	tsql = strings.TrimSuffix(tsql, ",")
	tsql += ")"
	//s.log.Debugf(tsql)

	tx := s.db.MustBegin()
	stmt, err := tx.PrepareNamedContext(s.ctx, tsql)
	if err != nil {
		return err
	}
	name := fmt.Sprintf("%-6s", tableName)
	p := mpb.New(mpb.WithWidth(64))
	bar := p.AddBar(int64(rowCount),
		mpb.PrependDecorators(
			decor.Name(name, decor.WC{W: len(name) + 1, C: decor.DidentRight}),
			decor.CountersNoUnit("%6d / %-6d", decor.WCSyncWidth),
		),
		mpb.AppendDecorators(decor.Percentage(decor.WC{W: 5})),
	)
	countInsert := 0
	for i := 0; i < rowCount; i++ {
		if _, err := stmt.ExecContext(s.ctx, data[i]); err != nil {
			return err
		}
		countInsert++
		bar.Increment()
	}
	p.Wait()
	tx.Commit()

	s.log.Debugf("вставили в %s  : %6d", tableName, countInsert)

	return nil
}
