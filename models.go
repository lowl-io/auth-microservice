package main

type User struct {
	ID       int       `gorm:"not null;type:serial;primary_key"`
	Name     string    `gorm:"not null;type:varchar(60);unique_index"`
	Email    string    `gorm:"type:varchar(255);unique_index"`
	Password string    `gorm:"not null;type:varchar(60)"`
	Status   string
}

type DataBaseConfig struct {
	Dialect      string `json:"dialect"`
	DataBaseInfo string `json:"databaseinfo"`
}

type JWTKeyConfig struct {
	Key string `json:"key"`
}

type Token struct {
	Token string `json:"token"`
}