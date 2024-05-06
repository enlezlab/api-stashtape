package util

import "github.com/google/uuid"

func UserIdGen() string {
	userId := uuid.New()
	return userId.String()
}
