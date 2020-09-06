package model

import "time"

// House ..
type House struct {
	AOGUID     string    `db:"aoguid"`     //varchar(36),
	BUILDNUM   string    `db:"buildnum"`   //varchar(10),
	ENDDATE    time.Time `db:"enddate"`    //date,
	ESTSTATUS  int       `db:"eststatus"`  //int,
	HOUSEGUID  string    `db:"houseguid"`  //varchar(36),
	HOUSEID    string    `db:"houseid"`    //varchar(36),
	HOUSENUM   string    `db:"housenum"`   //varchar(20),
	STATSTATUS int       `db:"statstatus"` //int,
	IFNSFL     string    `db:"ifnsfl"`     //varchar(4),
	IFNSUL     string    `db:"ifnsul"`     //varchar(4),
	OKATO      string    `db:"okato"`      //varchar(11),
	OKTMO      string    `db:"oktmo"`      //varchar(11),
	POSTALCODE string    `db:"postalcode"` //varchar(6),
	STARTDATE  time.Time `db:"startdate"`  //date,
	STRUCNUM   string    `db:"strucnum"`   //varchar(10),
	STRSTATUS  int       `db:"strstatus"`  //int,
	TERRIFNSFL string    `db:"terrifnsfl"` //varchar(4),
	TERRIFNSUL string    `db:"terrifnsul"` //varchar(4),
	UPDATEDATE time.Time `db:"updatedate"` //date,
	NORMDOC    string    `db:"normdoc"`    //varchar(36),
	COUNTER    int       `db:"counter"`    //int,
	CADNUM     string    `db:"cadnum"`     //varchar(100),
	DIVTYPE    int       `db:"divtype"`    //int
}

// Houses ..
type Houses struct {
	Dbf []interface{}   //House
}
