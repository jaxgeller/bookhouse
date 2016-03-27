package main

import "github.com/jinzhu/gorm"

type Book struct {
	gorm.Model
	Title  string
	Author string
	URL    string `gorm:"type:varchar(2083);unique_index;not null"`
	IMG    string `gorm:"type:varchar(2083);index"`
}
