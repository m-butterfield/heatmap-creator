package data

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID           string    `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Username     string    `gorm:"type:citext;not null;unique" json:"username"`
	DailyQueries int       `gorm:"not null"`
	CreatedAt    time.Time `gorm:"not null;default:now()"`
}

func (s *ds) CreateUser(user *User) error {
	if user.DailyQueries == 0 {
		user.DailyQueries = 20
	}
	if tx := s.db.Create(user); tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (s *ds) GetUser(username string) (*User, error) {
	user := &User{}
	tx := s.db.First(&user, "username = $1", username)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, tx.Error
	}
	return user, nil
}
