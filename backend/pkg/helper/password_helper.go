package helper

type PasswordHelper interface {
	Hash(password string) (string, error)
	IsMatch(hashedPassword string, password string) bool
}
