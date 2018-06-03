package models

import (
	"math/rand"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type PaymentTag struct {
	ID        int `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	PaymentID int `gorm:"not null" form:"paymentId" json:"paymentId"`
	TagID     int `gorm:"not null" form:"tagId" json:"tagId"`
}

func SamplePaymentTags() []PaymentTag {
	random := rand.New(rand.NewSource(1))
	random.Seed(5)
	paymentTags := make([]PaymentTag, 0, 10)
	for i := 0; i < 10; i++ {
		paymentTags = append(paymentTags,
			PaymentTag{
				ID:        i + 1,
				PaymentID: random.Intn(10-1) + 1,
				TagID:     random.Intn(15-1) + 1,
			})
	}
	return paymentTags
}

func InitPaymentTagTable() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./data.db")
	db.LogMode(true)

	if err != nil {
		panic(err)
	}

	if !db.HasTable(&PaymentTag{}) {
		db.CreateTable(&PaymentTag{})
		db.Set("gorm.table_options", "ENGINE=InnoDB").CreateTable(&PaymentTag{})

		// insert seeds
		for _, paymentTag := range SamplePaymentTags() {
			db.Create(&paymentTag)
		}
	}

	return db
}
