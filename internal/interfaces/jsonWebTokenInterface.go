package interfaces

type JsonWebToken interface {
	GenerateToken(userId uint) (map[string]interface{}, error)
	VerifyToken(tokenString string) (bool, error)
}
