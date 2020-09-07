package postgres

import (
	"fiastosql/internal/domain/model"
	"fmt"
	"strings"

	"github.com/lib/pq"
	"github.com/vbauerster/mpb"
	"github.com/vbauerster/mpb/decor"
)

// ToAddrob перенос таблицы ADDROB
func (s *DB) ToAddrob() error {

	if err := s.insertTo("addrob"); err != nil {
		return err
	}

	return nil
}

// ToHouse ...
func (s *DB) ToHouse() error {

	if err := s.insertTo("house"); err != nil {
		return err
	}

	return nil
}

// ToRoom ...
func (s *DB) ToRoom() error {

	if err := s.insertTo("room"); err != nil {
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
	case "addrob":
		dt, err := s.dbf.GetAddrobs()
		if err != nil {
			return err
		}
		rowCount = len(dt.Dbf)
		columns = model.GetColumns(model.Addrob{})
		data = dt.Dbf
	case "house":
		dt, err := s.dbf.GetHouses()
		if err != nil {
			return err
		}
		rowCount = len(dt.Dbf)
		columns = model.GetColumns(model.House{})
		data = dt.Dbf
	case "room":
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
	for i := 0; i < len(columns); i++ {
		columns[i] = strings.ToLower(columns[i])
	}

	tx := s.db.MustBegin()
	stmt, err := tx.PrepareContext(s.ctx, pq.CopyIn(tableName, columns...))
	if err != nil {
		return err
	}
	defer stmt.Close()
	s.log.Debugf("добавляем в %s", tableName)

	//=======ProgressBar ========================
	name := fmt.Sprintf("%-6s", tableName)
	p := mpb.New(mpb.WithWidth(64))
	bar := p.AddBar(int64(rowCount),
		mpb.PrependDecorators(
			decor.Name(name, decor.WC{W: len(name) + 1, C: decor.DidentRight}),
			decor.CountersNoUnit("%6d / %-6d", decor.WCSyncWidth),
		),
		mpb.AppendDecorators(decor.Percentage(decor.WC{W: 5})),
	)
	//===========================================

	for i := 0; i < rowCount; i++ {
		values := model.GetValues(data[i])
		if _, err := stmt.ExecContext(s.ctx, values...,
		); err != nil {
			return err
		}
		bar.Increment()
	}
	p.Wait()
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	err = stmt.Close()
	if err != nil {
		return err
	}
	tx.Commit()
	s.log.Debugf("добавили в %s : %6d", tableName, rowCount)
	return nil
}
