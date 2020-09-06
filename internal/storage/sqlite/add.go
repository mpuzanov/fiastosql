package sqlite

import (
	"fmt"

	"github.com/vbauerster/mpb/v5"
	"github.com/vbauerster/mpb/v5/decor"
)

// ToAddrob перенос таблицы ADDROB из дбф в sqlite
func (s *DB) ToAddrob() error {

	data, err := s.dbf.GetAddrobs()
	if err != nil {
		return err
	}
	rowCount := len(data.Dbf)
	s.log.Debugf("ADDROB: %6d", rowCount)

	tsql := `
	INSERT OR REPLACE INTO ADDROB(
		ACTSTATUS,
		AOGUID,
		AOID,
		AOLEVEL,
		AREACODE,
		AUTOCODE,
		CENTSTATUS,
		CITYCODE,
		CODE,
		CURRSTATUS,
		ENDDATE,
		FORMALNAME,
		IFNSFL,
		IFNSUL,
		NEXTID,
		OFFNAME,
		OKATO,
		OKTMO,
		OPERSTATUS,
		PARENTGUID,
		PLACECODE,
		PLAINCODE,
		POSTALCODE,
		PREVID,
		REGIONCODE,
		SHORTNAME,
		STARTDATE,
		STREETCODE,
		TERRIFNSFL,
		TERRIFNSUL,
		UPDATEDATE,
		CTARCODE,
		EXTRCODE,
		SEXTCODE,
		LIVESTATUS,
		NORMDOC,
		PLANCODE,
		CADNUM,
		DIVTYPE
	) values(
		:actstatus,
		:aoguid,
		:aoid,
		:aolevel,
		:areacode,
		:autocode,
		:centstatus,
		:citycode,
		:code,
		:currstatus,
		:enddate,
		:formalname,
		:ifnsfl,
		:ifnsul,
		:nextid,
		:offname,
		:okato,
		:oktmo,
		:operstatus,
		:parentguid,
		:placecode,
		:plaincode,
		:postalcode,
		:previd,
		:regioncode,
		:shortname,
		:startdate,
		:streetcode,
		:terrifnsfl,
		:terrifnsul,
		:updatedate,
		:ctarcode,
		:extrcode,
		:sextcode,
		:livestatus,
		:normdoc,
		:plancode,
		:cadnum,
		:divtype
	)
	`

	tx := s.db.MustBegin()

	stmt, err := tx.PrepareNamedContext(s.ctx, tsql)
	if err != nil {
		return err
	}

	name := fmt.Sprintf("%-6s", "ADDROB")
	p := mpb.New(mpb.WithWidth(64))
	bar := p.AddBar(int64(rowCount),
		mpb.PrependDecorators(
			decor.Name(name, decor.WC{W: len(name) + 1, C: decor.DidentRight}),
			decor.CountersNoUnit("%6d / %-6d", decor.WCSyncWidth),
		),
		mpb.AppendDecorators(decor.Percentage(decor.WC{W: 5})))

	countInsert := 0
	for i := 0; i < rowCount; i++ {
		zap := data.Dbf[i]
		if _, err := stmt.ExecContext(s.ctx, zap); err != nil {
			return err
		}
		countInsert++
		bar.Increment()
	}
	p.Wait()
	tx.Commit()

	s.log.Debugf("вставили в ADDROB: %6d", countInsert)

	return nil
}

// ToHouse ...
func (s *DB) ToHouse() error {

	data, err := s.dbf.GetHouses() // .getDbfTable("HOUSE")
	if err != nil {
		return err
	}
	rowCount := len(data.Dbf)
	s.log.Debugf("HOUSE: %6d", rowCount)

	tsql := `
	INSERT OR REPLACE INTO HOUSE(
		AOGUID,
		BUILDNUM,
		ENDDATE,
		ESTSTATUS,
		HOUSEGUID,
		HOUSEID,
		HOUSENUM,
		STATSTATUS,
		IFNSFL,
		IFNSUL,
		OKATO,
		OKTMO,
		POSTALCODE,
		STARTDATE,
		STRUCNUM,
		STRSTATUS,
		TERRIFNSFL,
		TERRIFNSUL,
		UPDATEDATE,
		NORMDOC,
		COUNTER,
		CADNUM,
		DIVTYPE
	) values(
		:aoguid,
		:buildnum,
		:enddate,
		:eststatus,
		:houseguid,
		:houseid,
		:housenum,
		:statstatus,
		:ifnsfl,
		:ifnsul,
		:okato,
		:oktmo,
		:postalcode,
		:startdate,
		:strucnum,
		:strstatus,
		:terrifnsfl,
		:terrifnsul,
		:updatedate,
		:normdoc,
		:counter,
		:cadnum,
		:divtype
	)
	`
	tx := s.db.MustBegin()
	stmt, err := tx.PrepareNamedContext(s.ctx, tsql)
	if err != nil {
		return err
	}
	name := fmt.Sprintf("%-6s", "HOUSE")
	p := mpb.New(mpb.WithWidth(64))
	bar := p.AddBar(int64(rowCount),
		mpb.PrependDecorators(
			decor.Name(name, decor.WC{W: len(name) + 1, C: decor.DidentRight}),
			decor.CountersNoUnit("%6d / %-6d", decor.WCSyncWidth),
		),
		mpb.AppendDecorators(decor.Percentage(decor.WC{W: 5})))

	countInsert := 0
	for i := 0; i < rowCount; i++ {
		zap := data.Dbf[i]
		if _, err := stmt.ExecContext(s.ctx, zap); err != nil {
			return err
		}
		countInsert++
		bar.Increment()
	}
	p.Wait()
	tx.Commit()

	s.log.Debugf("вставили в House : %6d", countInsert)

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

	tsql := `
	INSERT OR REPLACE INTO ROOM(
		ROOMID,
		ROOMGUID,
		HOUSEGUID,
		REGIONCODE,
		FLATNUMBER,
		FLATTYPE,
		ROOMNUMBER,
		ROOMTYPE,
		CADNUM,
		ROOMCADNUM,
		POSTALCODE,
		UPDATEDATE,
		PREVID,
		NEXTID,
		OPERSTATUS,
		STARTDATE,
		ENDDATE,
		LIVESTATUS,
		NORMDOC
	) values(
		:roomid,
		:roomguid,
		:houseguid,
		:regioncode,
		:flatnumber,
		:flattype,
		:roomnumber,
		:roomtype,
		:cadnum,
		:roomcadnum,
		:postalcode,
		:updatedate,
		:previd,
		:nextid,
		:operstatus,
		:startdate,
		:enddate,
		:livestatus,
		:normdoc
	)
	`
	tx := s.db.MustBegin()
	stmt, err := tx.PrepareNamedContext(s.ctx, tsql)
	if err != nil {
		return err
	}
	name := fmt.Sprintf("%-6s", "ROOM")
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
		zap := data.Dbf[i]

		if _, err := stmt.ExecContext(s.ctx, zap); err != nil {
			return err
		}
		countInsert++
		bar.Increment()
	}
	p.Wait()
	tx.Commit()

	s.log.Debugf("вставили в Room  : %6d", countInsert)

	return nil
}
