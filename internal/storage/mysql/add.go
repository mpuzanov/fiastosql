package mysql

import (
	"fiastosql/internal/domain/model"
	"fmt"
	"strings"

	"github.com/vbauerster/mpb"
	"github.com/vbauerster/mpb/decor"
)

// ToAddrob перенос таблицы ADDROB
func (s *DB) ToAddrob() error {

	if err := s.insertTo("addrob"); err != nil {
		return err
	}

	// data, err := s.dbf.GetAddrobs()
	// if err != nil {
	// 	return err
	// }
	// rowCount := len(data.Dbf)

	// columns := model.GetColumns(model.Addrob{})
	// values := strings.Repeat("?,", len(columns))
	// values = strings.TrimSuffix(values, ",")
	// tsql := "insert into Addrob(" + strings.Join(columns, ",") + ") values (" + values + ")"
	// //s.log.Debug(tsql)

	// tx := s.db.MustBegin()
	// stmt, err := tx.PrepareContext(s.ctx, tsql)
	// if err != nil {
	// 	return err
	// }
	// defer stmt.Close()
	// s.log.Debugf("добавляем в ADDROB")

	// //=======ProgressBar ========================
	// name := fmt.Sprintf("%-6s", "ADDROB")
	// p := mpb.New(mpb.WithWidth(64))
	// bar := p.AddBar(int64(rowCount),
	// 	mpb.PrependDecorators(
	// 		decor.Name(name, decor.WC{W: len(name) + 1, C: decor.DidentRight}),
	// 		decor.CountersNoUnit("%6d / %-6d", decor.WCSyncWidth),
	// 	),
	// 	mpb.AppendDecorators(decor.Percentage(decor.WC{W: 5})))
	// //===========================================

	// for i := 0; i < rowCount; i++ {
	// 	values := model.GetValues(data.Dbf[i])
	// 	if _, err := stmt.ExecContext(s.ctx, values...); err != nil {
	// 		return err
	// 	}
	// 	bar.Increment()
	// }
	// p.Wait()

	// err = stmt.Close()
	// if err != nil {
	// 	return err
	// }
	// tx.Commit()
	// s.log.Debugf("добавили в ADDROB: %6d", rowCount)

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

	placeHolders := strings.Repeat("?,", len(columns))
	placeHolders = fmt.Sprintf("(%s)", strings.TrimSuffix(placeHolders, ","))
	query := fmt.Sprintf("INSERT INTO %s ("+strings.Join(columns, ",")+") VALUES ", tableName)
	insertLimit := 1500

	s.log.Debugf("добавляем в %s", tableName)

	valueStrings := make([]string, 0, insertLimit)
	bulkValues := []interface{}{}
	record := make([]interface{}, len(columns))
	y := 0

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

	tx := s.db.MustBegin()
	for i := 0; i < rowCount; i++ {
		y++
		bar.Increment()
		valueStrings = append(valueStrings, placeHolders)
		record = model.GetValues(data[i])
		for y := 0; y < len(columns); y++ {
			bulkValues = append(bulkValues, record[y])
		}

		if y > insertLimit {
			stmt := query + strings.Join(valueStrings, ",")
			if _, err := tx.ExecContext(s.ctx, stmt, bulkValues...); err != nil {
				return err
			}
			valueStrings = valueStrings[:0]
			bulkValues = bulkValues[:0]
			y = 0
		}
	}
	stmt := query + strings.Join(valueStrings, ",")
	if _, err := tx.ExecContext(s.ctx, stmt, bulkValues...); err != nil {
		return err
	}
	p.Wait()

	tx.Commit()

	s.log.Debugf("добавили в %s : %6d", tableName, rowCount)
	return nil
}
