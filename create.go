package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func createDatabase() {
	db, err := gorm.Open("postgres", "user=postgres dbname=ambition password=ambition")
	check(err)
	db.DB()

	db.CreateTable(&Event{})
	db.CreateTable(&EventTime{})

}
