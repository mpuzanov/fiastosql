package sqlite

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
		
		CREATE TABLE IF NOT EXISTS ADDROB(
			ACTSTATUS int,
			AOGUID nvarchar(36),
			AOID nvarchar(36),
			AOLEVEL int,
			AREACODE nvarchar(3),
			AUTOCODE nvarchar(1),
			CENTSTATUS int,
			CITYCODE nvarchar(3),
			CODE nvarchar(17),
			CURRSTATUS int,
			ENDDATE date,
			FORMALNAME	nvarchar(120),
			IFNSFL nvarchar(4),
			IFNSUL nvarchar(4),
			NEXTID nvarchar(36),
			OFFNAME nvarchar(120),
			OKATO nvarchar(11),
			OKTMO nvarchar(11),
			OPERSTATUS int,
			PARENTGUID nvarchar(36),
			PLACECODE nvarchar(3),
			PLAINCODE nvarchar(15),
			POSTALCODE nvarchar(6),
			PREVID nvarchar(36),
			REGIONCODE nvarchar(2),
			SHORTNAME nvarchar(10),
			STARTDATE date,
			STREETCODE nvarchar(4),
			TERRIFNSFL nvarchar(4),
			TERRIFNSUL nvarchar(4),
			UPDATEDATE date,
			CTARCODE nvarchar(3),
			EXTRCODE nvarchar(4),
			SEXTCODE nvarchar(3),
			LIVESTATUS int,
			NORMDOC nvarchar(36),
			PLANCODE nvarchar(4),
			CADNUM nvarchar(100),
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

		CREATE TABLE IF NOT EXISTS HOUSE(
			AOGUID nvarchar(36),
			BUILDNUM nvarchar(10),
			ENDDATE date,
			ESTSTATUS int,
			HOUSEGUID nvarchar(36),
			HOUSEID nvarchar(36),
			HOUSENUM nvarchar(20),
			STATSTATUS int,
			IFNSFL nvarchar(4),
			IFNSUL nvarchar(4),
			OKATO nvarchar(11),
			OKTMO nvarchar(11),
			POSTALCODE nvarchar(6),
			STARTDATE date,
			STRUCNUM nvarchar(10),
			STRSTATUS int,
			TERRIFNSFL nvarchar(4),
			TERRIFNSUL nvarchar(4),
			UPDATEDATE date,
			NORMDOC nvarchar(36),
			COUNTER int,
			CADNUM nvarchar(100),
			DIVTYPE int		
	  );
	  `

	if _, err := s.db.ExecContext(s.ctx, sql); err != nil {
		return err
	}

	return nil
}

func (s *DB) createTableROOM() error {
	sql := `
		DROP TABLE IF EXISTS ROOM;

		CREATE TABLE IF NOT EXISTS ROOM(
			ROOMID nvarchar(36),
			ROOMGUID nvarchar(36),
			HOUSEGUID nvarchar(36),
			REGIONCODE nvarchar(2),
			FLATNUMBER nvarchar(50),
			FLATTYPE int,
			ROOMNUMBER nvarchar(50),
			ROOMTYPE nvarchar(2),
			CADNUM nvarchar(100),
			ROOMCADNUM nvarchar(100),
			POSTALCODE nvarchar(6),
			UPDATEDATE date,
			PREVID nvarchar(36),
			NEXTID nvarchar(36),
			OPERSTATUS int,
			STARTDATE date,
			ENDDATE date,
			LIVESTATUS int,
			NORMDOC nvarchar(36)
	  );
	  `

	if _, err := s.db.ExecContext(s.ctx, sql); err != nil {
		return err
	}

	return nil
}

// CreateIndex ...
func (s *DB) CreateIndex() error {

	tsql := `
	CREATE INDEX idx_find_streets ON ADDROB (AOLEVEL, PARENTGUID, ACTSTATUS);
	CREATE INDEX idx_find_houses ON HOUSE (AOGUID, ENDDATE);
	CREATE INDEX idx_find_room ON ROOM (HOUSEGUID, ENDDATE);
	`
	s.log.Debug("Создаём индексы...")
	if _, err := s.db.ExecContext(s.ctx, tsql); err != nil {
		return err
	}
	s.log.Debug("Индексы создали")
	return nil
}
