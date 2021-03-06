package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Payment struct {
	ID      int `gorm:"primary_key" form:"id" json:"id"`
	PlaceID int `gorm:"not null" form:"placeId" json:"placeId"`
	Cost    int `gorm:"not null" form:"cost" json:"cost"`
}

func SamplePayments() []Payment {
	payments := make([]Payment, 0, 10)
	for i := 0; i < 10; i++ {
		payments = append(payments,
			Payment{
				PlaceID: i + 1,
				Cost:    (i + 1) * 1000,
			})
	}
	return payments
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
