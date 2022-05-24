package util

import "github.com/google/uuid"

func NewUUIDStr() string {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return ""
	}
	return newUUID.String()
}
