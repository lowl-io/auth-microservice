package main

type User struct {
	id       int
	name     string `gorm:"index"`
	email    string
	password string `gorm:"index"`
}

type UserStatus bool

const (
	isActive  UserStatus = false
	isDeleted UserStatus = false
	isBlocked UserStatus = false
)