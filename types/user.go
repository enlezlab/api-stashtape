package types

type Entries struct {
	UID         string
	Title       string
	Description string
}

type User struct {
	USER_ID string
	ENTRIES []Entries
}
