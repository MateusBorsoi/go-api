package models

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Nome      string    `gorm:"not null"`
	Email     string    `gorm:"unique;not null"`
	Senha     string    `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Senha), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	u.Senha = string(hashedPassword)
	return nil
}

func (u *User) CheckPassword(senha string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Senha), []byte(senha))

}
