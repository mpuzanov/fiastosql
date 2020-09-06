package errors

// StrError для выдачи ошибок
type StrError string

func (ee StrError) Error() string {
	return string(ee)
}

var (
	//ErrNotExist Файл или каталог не существует
	ErrNotExist = StrError("Файл или каталог не существует")
	//ErrSQLNotExist Файл или каталог не существует
	ErrSQLNotExist = StrError("Файл БД не существует")
	//ErrNotRegion Не задан номер региона для выгрузки из DBF
	ErrNotRegion = StrError("Не задан номер региона для выгрузки из DBF")

	// ErrBadLoginDBConfiguration .
	ErrBadLoginDBConfiguration = StrError("Ошибка аутентификации при подключении к БД")
	// ErrNoDBAffected ошибка "Действие не затронуло ни одной строки"
	ErrNoDBAffected = StrError("Действие не затронуло ни одной строки")
	// ErrRecordNotFound ошибка "Запись не найдена"
	ErrRecordNotFound = StrError("Запись не найдена")
)
