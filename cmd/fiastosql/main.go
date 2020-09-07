package main

import (
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"

	"fiastosql/internal/storage"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	flag "github.com/spf13/pflag"
)

var (
	path, cfgPath, region, dbURL string
	verbose                      bool
)

func init() {
	flag.StringVarP(&path, "dir", "d", "", "путь к каталогу с файлами классификатора ФИАС")
	flag.StringVarP(&region, "region", "r", "", "код региона для импорта")
	flag.StringVar(&dbURL, "db_url", "", "строка подключения к БД приёмнику")
	flag.BoolVarP(&verbose, "verbose", "v", false, "подробный вывод")
	flag.Parse()

	if os.Getenv("REGION") != "" {
		region = os.Getenv("REGION")
	}
	if os.Getenv("DIR") != "" {
		path = os.Getenv("DIR")
	}
	if os.Getenv("DB_URL") != "" {
		dbURL = os.Getenv("DB_URL")
	}
}

func main() {
	if path == "" || region == "" {
		flag.PrintDefaults()
		return
	}

	logger := logrus.New()
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	logger.SetFormatter(customFormatter)
	if verbose {
		logger.SetLevel(logrus.DebugLevel)
	}

	fmt.Printf("Путь к dbf-файлам: %s, Регион: %s\n", path, region)

	dbf, err := storage.NewStorageDBF(path, region, logger)
	if err != nil {
		log.Fatalf("error NewStorageDBF: %s", err)
	}

	db, err := storage.NewStorageDB(dbURL, dbf, logger)
	if err != nil {
		log.Fatalf("error NewStorageDB: %s", err)
	}

	// начинаем обработку файлов
	start := time.Now()

	err = db.ToAddrob()
	if err != nil {
		logger.Error(err)
	}

	err = db.ToHouse()
	if err != nil {
		logger.Error(err)
	}

	err = db.ToRoom()
	if err != nil {
		logger.Error(err)
	}

	fmt.Printf("Заполнили таблицы за: %v\n", time.Since(start))

	err = db.CreateIndex()
	if err != nil {
		logger.Error(err)
	}

	fmt.Printf("Выполнено за: %v\n", time.Since(start))
}
