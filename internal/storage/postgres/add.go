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

	data, err := s.dbf.GetAddrobs()
	if err != nil {
		return err
	}
	rowCount := len(data.Dbf)

	tx := s.db.MustBegin()
	columns := model.GetColumns(model.Addrob{})
	for i := 0; i < len(columns); i++ {
		columns[i] = strings.ToLower(columns[i])
	}
	stmt, err := tx.PrepareContext(s.ctx, pq.CopyIn("addrob", columns...))
	if err != nil {
		return err
	}
	defer stmt.Close()
	s.log.Debugf("добавляем в ADDROB")

	//=======ProgressBar ========================
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
		values := model.GetValues(data.Dbf[i])
		if _, err := stmt.ExecContext(s.ctx, values...); err != nil {
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
	s.log.Debugf("добавили в ADDROB: %6d", rowCount)

	return nil
}

// ToHouse ...
func (s *DB) ToHouse() error {

	data, err := s.dbf.GetHouses()
	if err != nil {
		return err
	}
	rowCount := len(data.Dbf)

	tx := s.db.MustBegin()
	columns := model.GetColumns(model.House{})
	for i := 0; i < len(columns); i++ {
		columns[i] = strings.ToLower(columns[i])
	}
	stmt, err := tx.PrepareContext(s.ctx, pq.CopyIn("house", columns...))
	if err != nil {
		return err
	}
	defer stmt.Close()
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
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	err = stmt.Close()
	if err != nil {
		return err
	}
	tx.Commit()
	s.log.Debugf("добавили в HOUSE : %6d", rowCount)

	return nil
}

// ToRoom ...
func (s *DB) ToRoom() error {

	data, err := s.dbf.GetRooms()
	if err != nil {
		return err
	}
	rowCount := len(data.Dbf)

	tx := s.db.MustBegin()
	columns := model.GetColumns(model.Room{})
	for i := 0; i < len(columns); i++ {
		columns[i] = strings.ToLower(columns[i])
	}
	stmt, err := tx.PrepareContext(s.ctx, pq.CopyIn("room", columns...))
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
	s.log.Debugf("добавили в ROOM : %6d", rowCount)
	return nil
}

// func BulkInsertRoom(unsavedRows []*model.Room) error {
// 	valueStrings := make([]string, 0, len(unsavedRows))
// 	valueArgs := make([]interface{}, 0, len(unsavedRows)*3)
// 	i := 0
// 	for _, post := range unsavedRows {
// 		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d)", i*3+1, i*3+2, i*3+3))
// 		valueArgs = append(valueArgs, post.)
// 		valueArgs = append(valueArgs, post.Column2)
// 		valueArgs = append(valueArgs, post.Column3)
// 		i++
// 	}
// 	stmt := fmt.Sprintf("INSERT INTO ROOM (

// 	) VALUES %s", strings.Join(valueStrings, ","))
// 	_, err := db.Exec(stmt, valueArgs...)
// 	return err
// }

/*

//test batch inserts

sls := []Person{
	{FirstName: "Ardie", LastName: "Savea", Email: "asavea@ab.co.nz"},
	{FirstName: "Sonny Bill", LastName: "Williams", Email: "sbw@ab.co.nz"},
	{FirstName: "Ngani", LastName: "Laumape", Email: "nlaumape@ab.co.nz"},
}

_, err = db.NamedExec(`INSERT INTO person (first_name, last_name, email)
	VALUES (:first_name, :last_name, :email)`, sls)
*/

/*
func BulkInsert(unsavedRows []*ExampleRowStruct) error {
    valueStrings := make([]string, 0, len(unsavedRows))
    valueArgs := make([]interface{}, 0, len(unsavedRows) * 3)
    i := 0
    for _, post := range unsavedRows {
        valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d)", i*3+1, i*3+2, i*3+3))
        valueArgs = append(valueArgs, post.Column1)
        valueArgs = append(valueArgs, post.Column2)
        valueArgs = append(valueArgs, post.Column3)
        i++
    }
    stmt := fmt.Sprintf("INSERT INTO my_sample_table (column1, column2, column3) VALUES %s", strings.Join(valueStrings, ","))
    _, err := db.Exec(stmt, valueArgs...)
    return err
}
*/
