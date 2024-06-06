package utils

import (
	"github.com/Gabriel-Macedogmc/report-logs-golang/internal/interfaces"
	"golang.org/x/crypto/bcrypt"
)

type Bcrypt struct{}

func NewBcrypt() interfaces.BcryptInterface {
	return &Bcrypt{}
}

func (b *Bcrypt) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}

func (b *Bcrypt) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
