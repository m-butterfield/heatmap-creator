package data

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"time"
)

func Connect() (Store, error) {
	s, err := getDS()
	if err != nil {
		return nil, err
	}
	return s, nil
}

func getDS() (*ds, error) {
	var logLevel logger.LogLevel
	if os.Getenv("SQL_LOGS") == "true" {
		logLevel = logger.Info
	} else {
		logLevel = logger.Silent
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_SOCKET")), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		return nil, err
	}
	return &ds{db: db}, nil
}

type Store interface {
	CreateUser(user *User) error
	GetUser(username string) (*User, error)
	CreateAccessToken(user *User) (*AccessToken, error)
	DeleteAccessToken(id string) error
	GetAccessToken(id string) (*AccessToken, error)
	GetAccessTokenByQueryToken(queryToken string) (*AccessToken, error)
	CreateQuery(query *Query) error
	GetQueryCountForUser(id string, since *time.Time) (int, error)
}

type ds struct {
	db *gorm.DB
}
