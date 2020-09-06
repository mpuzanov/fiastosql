package dbf

import (
	"context"
	"fiastosql/internal/domain/errors"
	"fiastosql/internal/domain/model"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/LindsayBradford/go-dbf/godbf"
	log "github.com/sirupsen/logrus"
)

//DBF ...
type DBF struct {
	ctx          context.Context
	log          *log.Logger
	path         string // каталог с файлами
	region       string // СС – код субъекта Российской Федерации  – региона
	fileEncoding string // кодировка файлов "CP866" "CP1251" "UTF8"
}

//NewStore Возвращаем хранилище
func NewStore(ctx context.Context, path, region string, log *log.Logger) (*DBF, error) {

	fileEncoding := "CP866"
	//log.Debugf("Path: %s, Region: %s, FileEncoding: %s", path, region, fileEncoding)
	return &DBF{ctx: ctx,
		path:         path,
		region:       region,
		fileEncoding: fileEncoding,
		log:          log,
	}, nil
}

//strDateFormat форматируем дату YYYYMMDD => YYYY-MM-DD
func strDateFormat(s string) string {
	return s[0:4] + "-" + s[4:6] + "-" + s[6:8]
}

// GetPath ..
func (s *DBF) GetPath() string {
	return s.path
}

// GetAddrobs ..
func (s *DBF) GetAddrobs() (*model.Addrobs, error) {
	data := model.Addrobs{}

	path := filepath.Join(s.path, fmt.Sprintf("ADDROB%s.DBF", s.region))

	s.log.Debugf("файл: %s", path)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("%w <%s>", errors.ErrNotExist, path)
	}

	dbf, err := godbf.NewFromFile(path, s.fileEncoding)
	if err != nil {
		return nil, err
	}
	rowCount := dbf.NumberOfRecords()
	//rowCount = 1000
	for i := 0; i < rowCount; i++ {
		row := dbf.GetRowAsSlice(i)

		zap := model.Addrob{}
		v, _ := strconv.Atoi(row[0])
		zap.ACTSTATUS = v
		zap.AOGUID = row[1]
		zap.AOID = row[2]
		v, _ = strconv.Atoi(row[3])
		zap.AOLEVEL = v
		zap.AREACODE = row[4]
		zap.AUTOCODE = row[5]
		v, _ = strconv.Atoi(row[6])
		zap.CENTSTATUS = v
		zap.CITYCODE = row[7]
		zap.CODE = row[8]
		v, _ = strconv.Atoi(row[9])
		zap.CURRSTATUS = v
		t, _ := time.Parse("20060102", row[10])
		zap.ENDDATE = t //strDateFormat(row[10])
		zap.FORMALNAME = row[11]
		zap.IFNSFL = row[12]
		zap.IFNSUL = row[13]
		zap.NEXTID = row[14]
		zap.OFFNAME = row[15]
		zap.OKATO = row[16]
		zap.OKTMO = row[17]
		v, _ = strconv.Atoi(row[18])
		zap.OPERSTATUS = v
		zap.PARENTGUID = row[19]
		zap.PLACECODE = row[20]
		zap.PLAINCODE = row[21]
		zap.POSTALCODE = row[22]
		zap.PREVID = row[23]
		zap.REGIONCODE = row[24]
		zap.SHORTNAME = row[25]
		t, _ = time.Parse("20060102", row[26])
		zap.STARTDATE = t
		zap.STREETCODE = row[27]
		zap.TERRIFNSFL = row[28]
		zap.TERRIFNSUL = row[29]
		t, _ = time.Parse("20060102", row[30])
		zap.UPDATEDATE = t //strDateFormat(row[30])
		zap.CTARCODE = row[31]
		zap.EXTRCODE = row[32]
		zap.SEXTCODE = row[33]
		v, _ = strconv.Atoi(row[34])
		zap.LIVESTATUS = v
		zap.NORMDOC = row[35]
		zap.PLANCODE = row[36]
		zap.CADNUM = row[37]
		v, _ = strconv.Atoi(row[38])
		zap.DIVTYPE = v
		data.Dbf = append(data.Dbf, zap)
		//s.log.Debug(zap)
	}
	return &data, nil
}

//GetHouses дома по улице
func (s *DBF) GetHouses() (*model.Houses, error) {
	data := model.Houses{}

	path := filepath.Join(s.path, fmt.Sprintf("HOUSE%s.DBF", s.region))

	s.log.Debugf("файл: %s", path)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("%w <%s>", errors.ErrNotExist, path)
	}

	dbf, err := godbf.NewFromFile(path, s.fileEncoding)
	if err != nil {
		return nil, err
	}
	rowCount := dbf.NumberOfRecords()
	//rowCount = 1000
	for i := 0; i < rowCount; i++ {

		zap := model.House{}
		zap.AOGUID = dbf.FieldValue(i, 0)
		zap.BUILDNUM = dbf.FieldValue(i, 1)
		t, _ := time.Parse("20060102", dbf.FieldValue(i, 2))
		zap.ENDDATE = t
		v, _ := strconv.Atoi(dbf.FieldValue(i, 3))
		zap.ESTSTATUS = v
		zap.HOUSEGUID = dbf.FieldValue(i, 4)
		zap.HOUSEID = dbf.FieldValue(i, 5)
		zap.HOUSENUM = dbf.FieldValue(i, 6)
		v, _ = strconv.Atoi(dbf.FieldValue(i, 7))
		zap.STATSTATUS = v
		zap.IFNSFL = dbf.FieldValue(i, 8)
		zap.IFNSUL = dbf.FieldValue(i, 9)
		zap.OKATO = dbf.FieldValue(i, 10)
		zap.OKTMO = dbf.FieldValue(i, 11)
		zap.POSTALCODE = dbf.FieldValue(i, 12)
		t, _ = time.Parse("20060102", dbf.FieldValue(i, 13))
		zap.STARTDATE = t
		zap.STRUCNUM = dbf.FieldValue(i, 14)
		v, _ = strconv.Atoi(dbf.FieldValue(i, 15))
		zap.STRSTATUS = v
		zap.TERRIFNSFL = dbf.FieldValue(i, 16)
		zap.TERRIFNSUL = dbf.FieldValue(i, 17)
		t, _ = time.Parse("20060102", dbf.FieldValue(i, 18))
		zap.UPDATEDATE = t
		zap.NORMDOC = dbf.FieldValue(i, 19)
		v, _ = strconv.Atoi(dbf.FieldValue(i, 20))
		zap.COUNTER = v
		zap.CADNUM = dbf.FieldValue(i, 21)
		v, _ = strconv.Atoi(dbf.FieldValue(i, 22))
		zap.DIVTYPE = v
		data.Dbf = append(data.Dbf, zap)
		//s.log.Debug(zap)
	}

	return &data, nil
}

//GetRooms Квартиры по дому
func (s *DBF) GetRooms() (*model.Rooms, error) {
	data := model.Rooms{}

	path := filepath.Join(s.path, fmt.Sprintf("ROOM%s.DBF", s.region))

	s.log.Debugf("файл: %s", path)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("%w <%s>", errors.ErrNotExist, path)
	}

	dbf, err := godbf.NewFromFile(path, s.fileEncoding)
	if err != nil {
		return nil, err
	}
	rowCount := dbf.NumberOfRecords()
	//rowCount = 1000
	for i := 0; i < rowCount; i++ {

		zap := model.Room{}
		zap.ROOMID = dbf.FieldValue(i, 0)
		zap.ROOMGUID = dbf.FieldValue(i, 1)
		zap.HOUSEGUID = dbf.FieldValue(i, 2)
		zap.REGIONCODE = dbf.FieldValue(i, 3)
		zap.FLATNUMBER = dbf.FieldValue(i, 4)
		v, _ := strconv.Atoi(dbf.FieldValue(i, 5))
		zap.FLATTYPE = v
		zap.ROOMNUMBER = dbf.FieldValue(i, 6)
		zap.ROOMTYPE = dbf.FieldValue(i, 7)
		zap.CADNUM = dbf.FieldValue(i, 8)
		zap.ROOMCADNUM = dbf.FieldValue(i, 9)
		zap.POSTALCODE = dbf.FieldValue(i, 10)
		t, _ := time.Parse("20060102", dbf.FieldValue(i, 11))
		zap.UPDATEDATE = t //UPDATEDATE
		zap.PREVID = dbf.FieldValue(i, 12)
		zap.NEXTID = dbf.FieldValue(i, 13)
		v, _ = strconv.Atoi(dbf.FieldValue(i, 14))
		zap.OPERSTATUS = v
		t, _ = time.Parse("20060102", dbf.FieldValue(i, 15))
		zap.STARTDATE = t //STARTDATE
		t, _ = time.Parse("20060102", dbf.FieldValue(i, 16))
		zap.ENDDATE = t //ENDDATE
		v, _ = strconv.Atoi(dbf.FieldValue(i, 17))
		zap.LIVESTATUS = v
		zap.NORMDOC = dbf.FieldValue(i, 18)

		data.Dbf = append(data.Dbf, zap)
		//s.log.Debug(zap)
	}

	return &data, nil
}
