package repo

import (
	"fiastosql/internal/domain/model"
)

//DatastoreFias интерфейс БД адресного классификатора
type DatastoreFias interface {
	GetAddrobs() (*model.Addrobs, error)
	//Чтение домов из ФИАС
	GetHouses() (*model.Houses, error)
	//Чтение квартир из ФИАС
	GetRooms() (*model.Rooms, error)
	//Возвращаем путь к базе (к файлам)
	GetPath() string
}

//DatastoreDB ...
type DatastoreDB interface {
	// Переносим ADDROB
	ToAddrob() error
	// Переносим HOUSE
	ToHouse() error
	// Переносим ROOM
	ToRoom() error
	// Создаём таблицы в БД
	CreateTables() error
	// Создаём индексы в БД
	CreateIndex() error
}
