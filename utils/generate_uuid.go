package utils

import "math/rand"

func GenerateJobUuid() string {
	return generateJobUuidHead(2) + "-" + generateJobUuidTail(5)
}
func generateJobUuidHead(n int) string {
	var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
func generateJobUuidTail(n int) string {
	var letters = []rune("0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
