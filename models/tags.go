package models

import "github.com/jinzhu/gorm"
import _ "github.com/mattn/go-sqlite3"

type Tag struct {
	Id   int    `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Name string `gorm:"not null" form:"tagname" json:"tagname"`
}

func SampleTags() []Tag {
	tagName := []string{"電気代", "食費", "家賃", "水道代", "ガス代", "交通費", "酒代", "飲み物代", "雑費", "菓子代", "電化製品", "服代", "本代", "通信代", "医療費"}
	tag := make([]Tag, 0, 15)
	for i := 0; i < 10; i++ {
		tag = append(tag, Tag{Id: i + 1, Name: tagName[i]})
	}
	return tag
}

func InitTagTable() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./data.db")
	db.LogMode(true)

	if err != nil {
		panic(err)
	}

	if !db.HasTable(&Tag{}) {
		db.CreateTable(&Tag{})
		db.Set("gorm.table_options", "ENGINE=InnoDB").CreateTable(&Tag{})

		// insert seeds
		for _, tag := range SampleTags() {
			db.Create(&tag)
		}
	}

	return db
}
