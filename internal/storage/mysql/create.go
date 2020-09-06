package mysql

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
	if _, err := s.db.ExecContext(s.ctx, `DROP TABLE IF EXISTS addrob;`); err != nil {
		return err
	}
	sql := `	
	CREATE TABLE IF NOT EXISTS addrob(
		actstatus INT,
		aoguid VARCHAR(36),
		aoid VARCHAR(36),
		aolevel INT,
		areacode VARCHAR(3),
		autocode VARCHAR(1),
		centstatus INT,
		citycode VARCHAR(3),
		code VARCHAR(17),
		currstatus INT,
		enddate DATE,
		formalname	VARCHAR(120),
		ifnsfl VARCHAR(4),
		ifnsul VARCHAR(4),
		nextid VARCHAR(36),
		offname VARCHAR(120),
		okato VARCHAR(11),
		oktmo VARCHAR(11),
		operstatus INT,
		parentguid VARCHAR(36),
		placecode VARCHAR(3),
		plaincode VARCHAR(15),
		postalcode VARCHAR(6),
		previd VARCHAR(36),
		regioncode VARCHAR(2),
		shortname VARCHAR(10),
		startdate DATE,
		streetcode VARCHAR(4),
		terrifnsfl VARCHAR(4),
		terrifnsul VARCHAR(4),
		updatedate DATE,
		ctarcode VARCHAR(3),
		extrcode VARCHAR(4),
		sextcode VARCHAR(3),
		livestatus INT,
		normdoc VARCHAR(36),
		plancode VARCHAR(4),
		cadnum VARCHAR(100),
		divtype INT
	);
	`

	if _, err := s.db.ExecContext(s.ctx, sql); err != nil {
		return err
	}

	return nil
}

func (s *DB) createTableHOUSE() error {
	if _, err := s.db.ExecContext(s.ctx, `DROP TABLE IF EXISTS house;`); err != nil {
		return err
	}
	sql := `
	CREATE TABLE IF NOT EXISTS house(
			aoguid VARCHAR(36),
			buildnum VARCHAR(10),
			enddate DATE,
			eststatus INT,
			houseguid VARCHAR(36),
			houseid VARCHAR(36),
			housenum VARCHAR(20),
			statstatus INT,
			ifnsfl VARCHAR(4),
			ifnsul VARCHAR(4),
			okato VARCHAR(11),
			oktmo VARCHAR(11),
			postalcode VARCHAR(6),
			startdate DATE,
			strucnum VARCHAR(10),
			strstatus INT,
			terrifnsfl VARCHAR(4),
			terrifnsul VARCHAR(4),
			updatedate DATE,
			normdoc VARCHAR(36),
			counter INT,
			cadnum VARCHAR(100),
			divtype INT		
	  );`

	if _, err := s.db.ExecContext(s.ctx, sql); err != nil {
		return err
	}

	return nil
}

func (s *DB) createTableROOM() error {
	if _, err := s.db.ExecContext(s.ctx, `DROP TABLE IF EXISTS room;`); err != nil {
		return err
	}
	sql := `
	CREATE TABLE IF NOT EXISTS room(
			roomid VARCHAR(36),
			roomguid VARCHAR(36),
			houseguid VARCHAR(36),
			regioncode VARCHAR(2),
			flatnumber VARCHAR(50),
			flattype INT,
			roomnumber VARCHAR(50),
			roomtype VARCHAR(2),
			cadnum VARCHAR(100),
			roomcadnum VARCHAR(100),
			postalcode VARCHAR(6),
			updatedate DATE,
			previd VARCHAR(36),
			nextid VARCHAR(36),
			operstatus INT,
			startdate DATE,
			enddate DATE,
			livestatus INT,
			normdoc VARCHAR(36)
	  );`

	if _, err := s.db.ExecContext(s.ctx, sql); err != nil {
		return err
	}

	return nil
}

// CreateIndex ...
func (s *DB) CreateIndex() error {

	s.log.Debug("Создаём индексы...")
	if _, err := s.db.ExecContext(s.ctx, `CREATE INDEX idx_find_streets ON addrob (aolevel, parentguid, actstatus);`); err != nil {
		return err
	}
	if _, err := s.db.ExecContext(s.ctx, `CREATE INDEX idx_find_houses ON house (aoguid, enddate);`); err != nil {
		return err
	}
	if _, err := s.db.ExecContext(s.ctx, `CREATE INDEX idx_find_room ON room (houseguid, enddate);`); err != nil {
		return err
	}
	s.log.Debug("Индексы создали")
	return nil
}
