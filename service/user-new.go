package service

import (
	"stashtape/data"
	"stashtape/model"

	"github.com/google/uuid"
)

func NewUser(entries model.Entries) string {

	userId := uuid.New().String()

	entriesMap := []model.Entries{}

	entriesMap = append(entriesMap, entries)

	res := data.NewEntry("USER_CONTENT", model.User{
		USER_ID: userId,
		ENTRIES: entriesMap,
	})

	return res
}
