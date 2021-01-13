package security

import "github.com/google/uuid"

// GenerateUUID generate uuid
func GenerateUUID() string {
	id, err := uuid.NewRandom()
	if err != nil {
		return ""
	}
	return id.String()
}
