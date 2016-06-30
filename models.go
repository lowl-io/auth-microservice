package main

type User struct {
	id       int    `gorm:"not null;primary_key"`
	name     string `gorm:"type:varchar(60);unique_index"`
	email    string `gorm:"type:varchar(255)"`
	password string `gorm:"type:varchar(60)"`
}

type UserStatus bool

const (
	isActive UserStatus = true
	isDeleted UserStatus = true
	isBlocked UserStatus = true
)