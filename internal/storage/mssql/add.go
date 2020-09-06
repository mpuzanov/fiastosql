package mssql

import (
	"fiastosql/internal/domain/model"
	"fmt"

	mssql "github.com/denisenkom/go-mssqldb"
	"github.com/vbauerster/mpb/v5"
	"github.com/vbauerster/mpb/v5/decor"
)

// ToAddrob перенос таблицы ADDROB из дбф
func (s *DB) ToAddrob() error {

	data, err := s.dbf.GetAddrobs()
	if err != nil {
		return err
	}
	rowCount := len(data.Dbf)
	s.log.Debugf("ADDROB: %6d", rowCount)

	tx := s.db.MustBegin()
	columns := model.GetColumns(model.Addrob{})
	stmt, err := tx.Prepare(mssql.CopyIn("ADDROB", mssql.BulkOptions{}, columns...))
	if err != nil {
		return err
	}

	s.log.Debugf("добавляем в ADDROB")
	//=======ProgresBar =========================
	name := fmt.Sprintf("%-6s", "ADDROB")
	p := mpb.New(mpb.WithWidth(64))
	bar := p.AddBar(int64(rowCount),
		mpb.PrependDecorators(
			decor.Name(name, decor.WC{W: len(name) + 1, C: decor.DidentRight}),
			decor.CountersNoUnit("%6d / %-6d", decor.WCSyncWidth),
		),
		mpb.AppendDecorators(decor.Percentage(decor.WC{W: 5})))
	//===========================================

	for i := 0; i < rowCount; i++ {
		zap := data.Dbf[i]
		values := model.GetValues(zap)
		if _, err := stmt.ExecContext(s.ctx, values...,
		); err != nil {
			return err
		}
		bar.Increment()
	}
	p.Wait()

	result, err := stmt.Exec()
	if err != nil {
		return err
	}
	err = stmt.Close()
	if err != nil {
		return err
	}
	tx.Commit()
	count, _ := result.RowsAffected()
	s.log.Debugf("вставили в ADDROB : %6d", count)

	return nil
}

// ToHouse ...
func (s *DB) ToHouse() error {

	data, err := s.dbf.GetHouses()
	if err != nil {
		return err
	}
	rowCount := len(data.Dbf)
	s.log.Debugf("HOUSE: %6d", rowCount)

	tx := s.db.MustBegin()
	columns := model.GetColumns(model.House{})
	stmt, err := tx.Prepare(mssql.CopyIn("HOUSE", mssql.BulkOptions{}, columns...,
	))
	s.log.Debugf("добавляем в HOUSE")

	//=======ProgressBar ========================
	name := fmt.Sprintf("%-6s", "HOUSE")
	p := mpb.New(mpb.WithWidth(64))
	bar := p.AddBar(int64(rowCount),
		mpb.PrependDecorators(
			decor.Name(name, decor.WC{W: len(name) + 1, C: decor.DidentRight}),
			decor.CountersNoUnit("%6d / %-6d", decor.WCSyncWidth),
		),
		mpb.AppendDecorators(decor.Percentage(decor.WC{W: 5})))
	//===========================================

	for i := 0; i < rowCount; i++ {
		values := model.GetValues(data.Dbf[i])
		if _, err := stmt.ExecContext(s.ctx, values...,
		); err != nil {
			return err
		}
		bar.Increment()
	}
	p.Wait()

	result, err := stmt.Exec()
	if err != nil {
		return err
	}
	err = stmt.Close()
	if err != nil {
		return err
	}
	tx.Commit()
	count, _ := result.RowsAffected()
	s.log.Debugf("вставили в House : %6d", count)
	return nil
}

// ToRoom ...
func (s *DB) ToRoom() error {

	data, err := s.dbf.GetRooms()
	if err != nil {
		return err
	}
	rowCount := len(data.Dbf)
	s.log.Debugf("ROOM: %6d", rowCount)

	tx := s.db.MustBegin()
	columns := model.GetColumns(model.Room{})
	stmt, err := tx.Prepare(mssql.CopyIn("ROOM", mssql.BulkOptions{}, columns...))
	if err != nil {
		return err
	}
	defer stmt.Close()
	s.log.Debugf("добавляем в ROOM")

	//=======ProgressBar ========================
	name := fmt.Sprintf("%-6s", "ROOM")
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
		values := model.GetValues(data.Dbf[i])
		if _, err := stmt.ExecContext(s.ctx, values...); err != nil {
			return err
		}
		bar.Increment()
	}
	p.Wait()
	result, err := stmt.Exec()
	if err != nil {
		return err
	}
	err = stmt.Close()
	if err != nil {
		return err
	}
	tx.Commit()
	count, _ := result.RowsAffected()
	s.log.Debugf("вставили в ROOM : %6d", count)
	return nil
}
