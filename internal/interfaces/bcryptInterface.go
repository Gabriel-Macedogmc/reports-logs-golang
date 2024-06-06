package interfaces

type BcryptInterface interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}
