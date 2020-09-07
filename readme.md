# Утилита извлекает данные из классификатора ФИАС (формата dbf) и записывает в требуемую БД(Sqlite, MSSQL, PostgreSql, MySQL) 

## Запуск

Примеры:  

	./fias-to-sql -d H:\\fias\\fias18 -r 18 -v
	./fias-to-sql -d H:\\fias\\fias18 -r 18 -v --db_url postgres://postgres:123@localhost:5432/fias?sslmode=disable
	./fias-to-sql -d H:\\fias\\fias18 -r 18 -v --db_url sqlserver://sa:123@localhost?database=fias
	./fias-to-sql -d H:\\fias\\fias18 -r 18 -v --db_url root:dnypr1@/fias    

dir (-d) - путь к базе с дбф-файлами (там будет создана база sqlite)  
region (-r) - номер региона  
verbose (-v) - подробный вывод  

*Строки подключения к БД:*

    postgres://postgres:123@localhost:5432/fias?sslmode=disable
    sqlserver://sa:123@localhost:1432?database=fias