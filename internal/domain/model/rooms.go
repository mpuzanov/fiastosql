package model

import "time"

// Room ..
type Room struct {
	ROOMID     string    `db:"roomid"`     //varchar(36),
	ROOMGUID   string    `db:"roomguid"`   //varchar(36),
	HOUSEGUID  string    `db:"houseguid"`  //varchar(36),
	REGIONCODE string    `db:"regioncode"` //varchar(2),
	FLATNUMBER string    `db:"flatnumber"` //varchar(50),
	FLATTYPE   int       `db:"flattype"`   //int,
	ROOMNUMBER string    `db:"roomnumber"` //varchar(50),
	ROOMTYPE   string    `db:"roomtype"`   //varchar(2),
	CADNUM     string    `db:"cadnum"`     //varchar(100),
	ROOMCADNUM string    `db:"roomcadnum"` //varchar(100),
	POSTALCODE string    `db:"postalcode"` //varchar(6),
	UPDATEDATE time.Time `db:"updatedate"` //date,
	PREVID     string    `db:"previd"`     //varchar(36),
	NEXTID     string    `db:"nextid"`     //varchar(36),
	OPERSTATUS int       `db:"operstatus"` //int,
	STARTDATE  time.Time `db:"startdate"`  //date,
	ENDDATE    time.Time `db:"enddate"`    //date,
	LIVESTATUS int       `db:"livestatus"` //int,
	NORMDOC    string    `db:"normdoc"`    //varchar(36)
}

// Rooms ..
type Rooms struct {
	Dbf []interface{} //Room
}
