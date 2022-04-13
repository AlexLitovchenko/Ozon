package DataBase

import (
	"test/structure"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=ajk354rmlet dbname=OZON port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Не удалось подключиться к базе данных")
	}

	//Миграция нашей структуры в бд
	db.AutoMigrate(&structure.URLS{})
	return db
}
