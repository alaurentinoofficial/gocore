package hashs

type HashService interface {
	Hash(value string) string
	Validate(hashedValue string, value string) bool
}
