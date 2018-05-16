package models

import "github.com/jinzhu/gorm"
import _ "github.com/mattn/go-sqlite3"

type Payment struct {
	Id      int `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	PlaceId int `gorm:"not null" form:"placeId" json:"placeId"`
	Cost    int `gorm:"not null" form:"cost" json:"cost"`
}

func SamplePayments() []Payment {
	payment := make([]Payment, 0, 10)
	for i := 0; i < 10; i++ {
		payment = append(payment, Payment{PlaceId: i + 1, Cost: (i + 1) * 1000})
	}
	return payment
}

func InitPaymentTable() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./data.db")
	db.LogMode(true)

	if err != nil {
		panic(err)
	}

	if !db.HasTable(&Payment{}) {
		db.CreateTable(&Payment{})
		db.Set("gorm.table_options", "ENGINE=InnoDB").CreateTable(&Payment{})

		// insert seeds
		for _, payment := range SamplePayments() {
			db.Create(&payment)
		}
	}

	return db
}
