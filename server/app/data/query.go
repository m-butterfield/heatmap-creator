package data

import (
	"time"
)

type Query struct {
	ID        string `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	QueryText string `gorm:"not null"`
	UserID    string `gorm:"not null"`
	User      *User
	CreatedAt time.Time `gorm:"not null;default:now()"`
}

func (s *ds) CreateQuery(query *Query) error {
	if tx := s.db.Create(query); tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (s *ds) GetQueryCountForUser(id string, since *time.Time) (int, error) {
	var count int64
	tx := s.db.
		Model(&Query{}).
		Where("user_id = ?", id).
		Where("created_at > ?", since).
		Count(&count)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(count), nil
}
