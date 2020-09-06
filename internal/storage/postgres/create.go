package postgres

import "fmt"

// CreateTables ...
func (s *DB) CreateTables() error {

	if err := s.createTableADDROB(); err != nil {
		return fmt.Errorf("%w create table ADDROB", err)
	}
	if err := s.createTableHOUSE(); err != nil {
		return fmt.Errorf("%w create table HOUSE", err)
	}
	if err := s.createTableROOM(); err != nil {
		return fmt.Errorf("%w create table ROOM", err)
	}
	return nil
}

func (s *DB) createTableADDROB() error {
	sql := `
		DROP TABLE IF EXISTS ADDROB;
		
		CREATE TABLE ADDROB(
			ACTSTATUS int,
			AOGUID varchar(36),
			AOID varchar(36),
			AOLEVEL int,
			AREACODE varchar(3),
			AUTOCODE varchar(1),
			CENTSTATUS int,
			CITYCODE varchar(3),
			CODE varchar(17),
			CURRSTATUS int,
			ENDDATE date,
			FORMALNAME	varchar(120),
			IFNSFL varchar(4),
			IFNSUL varchar(4),
			NEXTID varchar(36),
			OFFNAME varchar(120),
			OKATO varchar(11),
			OKTMO varchar(11),
			OPERSTATUS int,
			PARENTGUID varchar(36),
			PLACECODE varchar(3),
			PLAINCODE varchar(15),
			POSTALCODE varchar(6),
			PREVID varchar(36),
			REGIONCODE varchar(2),
			SHORTNAME varchar(10),
			STARTDATE date,
			STREETCODE varchar(4),
			TERRIFNSFL varchar(4),
			TERRIFNSUL varchar(4),
			UPDATEDATE date,
			CTARCODE varchar(3),
			EXTRCODE varchar(4),
			SEXTCODE varchar(3),
			LIVESTATUS int,
			NORMDOC varchar(36),
			PLANCODE varchar(4),
			CADNUM varchar(100),
			DIVTYPE int
		);
		`

	if _, err := s.db.ExecContext(s.ctx, sql); err != nil {
		return err
	}

	return nil
}

func (s *DB) createTableHOUSE() error {
	sql := `
		DROP TABLE IF EXISTS HOUSE;

		CREATE TABLE HOUSE(
			AOGUID varchar(36),
			BUILDNUM varchar(10),
			ENDDATE date,
			ESTSTATUS int,
			HOUSEGUID varchar(36),
			HOUSEID varchar(36),
			HOUSENUM varchar(20),
			STATSTATUS int,
			IFNSFL varchar(4),
			IFNSUL varchar(4),
			OKATO varchar(11),
			OKTMO varchar(11),
			POSTALCODE varchar(6),
			STARTDATE date,
			STRUCNUM varchar(10),
			STRSTATUS int,
			TERRIFNSFL varchar(4),
			TERRIFNSUL varchar(4),
			UPDATEDATE date,
			NORMDOC varchar(36),
			COUNTER int,
			CADNUM varchar(100),
			DIVTYPE int		
	  );`

	if _, err := s.db.ExecContext(s.ctx, sql); err != nil {
		return err
	}

	return nil
}

func (s *DB) createTableROOM() error {
	sql := `
		DROP TABLE IF EXISTS ROOM;

		CREATE TABLE ROOM(
			ROOMID varchar(36),
			ROOMGUID varchar(36),
			HOUSEGUID varchar(36),
			REGIONCODE varchar(2),
			FLATNUMBER varchar(50),
			FLATTYPE int,
			ROOMNUMBER varchar(50),
			ROOMTYPE varchar(2),
			CADNUM varchar(100),
			ROOMCADNUM varchar(100),
			POSTALCODE varchar(6),
			UPDATEDATE date,
			PREVID varchar(36),
			NEXTID varchar(36),
			OPERSTATUS int,
			STARTDATE date,
			ENDDATE date,
			LIVESTATUS int,
			NORMDOC varchar(36)
	  );`

	if _, err := s.db.ExecContext(s.ctx, sql); err != nil {
		return err
	}

	return nil
}

// CreateIndex ...
func (s *DB) CreateIndex() error {
	sql := `
	CREATE INDEX idx_find_streets ON ADDROB (AOLEVEL, PARENTGUID, ACTSTATUS);
	CREATE INDEX idx_find_houses ON HOUSE (AOGUID, ENDDATE);
	CREATE INDEX idx_find_room ON ROOM (HOUSEGUID, ENDDATE);
	`
	s.log.Debug("Создаём индексы...")
	if _, err := s.db.ExecContext(s.ctx, sql); err != nil {
		return err
	}
	s.log.Debug("Индексы создали")
	return nil
}
