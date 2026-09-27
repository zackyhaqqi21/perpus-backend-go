package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"perpusbackendgo/model"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=zacky dbname=perpus_db port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Gagal konek ke database: " + err.Error())
	}

	db.AutoMigrate(&model.Book{})

	DB = db
}
