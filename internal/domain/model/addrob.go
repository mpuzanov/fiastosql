package model

import (
	"reflect"
	"time"
)

// Addrob ..
type Addrob struct {
	ACTSTATUS  int       `db:"actstatus"`  //int
	AOGUID     string    `db:"aoguid"`     //varchar(36)
	AOID       string    `db:"aoid"`       //varchar(36)
	AOLEVEL    int       `db:"aolevel"`    //int
	AREACODE   string    `db:"areacode"`   //varchar(3)
	AUTOCODE   string    `db:"autocode"`   //varchar(1)
	CENTSTATUS int       `db:"centstatus"` //int
	CITYCODE   string    `db:"citycode"`   //varchar(3)
	CODE       string    `db:"code"`       //varchar(17)
	CURRSTATUS int       `db:"currstatus"` //int
	ENDDATE    time.Time `db:"enddate"`    //date
	FORMALNAME string    `db:"formalname"` //varchar(120)
	IFNSFL     string    `db:"ifnsfl"`     //varchar(4)
	IFNSUL     string    `db:"ifnsul"`     //varchar(4)
	NEXTID     string    `db:"nextid"`     //varchar(36)
	OFFNAME    string    `db:"offname"`    //varchar(120)
	OKATO      string    `db:"okato"`      //varchar(11)
	OKTMO      string    `db:"oktmo"`      //varchar(11)
	OPERSTATUS int       `db:"operstatus"` //int
	PARENTGUID string    `db:"parentguid"` //varchar(36)
	PLACECODE  string    `db:"placecode"`  //varchar(3)
	PLAINCODE  string    `db:"plaincode"`  //varchar(15)
	POSTALCODE string    `db:"postalcode"` //varchar(6)
	PREVID     string    `db:"previd"`     //varchar(36)
	REGIONCODE string    `db:"regioncode"` //varchar(2)
	SHORTNAME  string    `db:"shortname"`  //varchar(10)
	STARTDATE  time.Time `db:"startdate"`  //date
	STREETCODE string    `db:"streetcode"` //varchar(4)
	TERRIFNSFL string    `db:"terrifnsfl"` //varchar(4)
	TERRIFNSUL string    `db:"terrifnsul"` //varchar(4)
	UPDATEDATE time.Time `db:"updatedate"` //date
	CTARCODE   string    `db:"ctarcode"`   //varchar(3)
	EXTRCODE   string    `db:"extrcode"`   //varchar(4)
	SEXTCODE   string    `db:"sextcode"`   //varchar(3)
	LIVESTATUS int       `db:"livestatus"` //int
	NORMDOC    string    `db:"normdoc"`    //varchar(36)
	PLANCODE   string    `db:"plancode"`   //varchar(4)
	CADNUM     string    `db:"cadnum"`     //varchar(100)
	DIVTYPE    int       `db:"divtype"`    //int
}

// Addrobs ..
type Addrobs struct {
	Dbf []interface{} //Addrob
}

// GetFields ..
func GetFields(s interface{}) map[string]interface{} {

	fields := make(map[string]interface{})
	v := reflect.ValueOf(s)
	typeOfS := v.Type()

	for i := 0; i < v.NumField(); i++ {
		//fmt.Printf("Field: %s\tValue: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
		fields[typeOfS.Field(i).Name] = v.Field(i).Interface()
	}
	return fields
}

// GetColumns ..
func GetColumns(s interface{}) []string {

	v := reflect.ValueOf(s)
	typeOfS := v.Type()

	names := make([]string, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		//fmt.Printf("Field: %s\tValue: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
		names[i] = typeOfS.Field(i).Name
	}
	return names
}

// GetValues ..
func GetValues(s interface{}) []interface{} {

	v := reflect.ValueOf(s)
	values := make([]interface{}, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		//fmt.Printf("Field: %s\tValue: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
		values[i] = v.Field(i).Interface()
	}
	return values
}
