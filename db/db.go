package db

import (
	"github.com/Masato4556/kurayami-go/entity"
	"github.com/jinzhu/gorm"
)

//DB初期化
func Init() {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベース開けず！（dbInit）")
	}
	autoMigration()
	defer db.Close()
}

func autoMigration() {
	db.AutoMigrate(&entity.Todo{})
}
