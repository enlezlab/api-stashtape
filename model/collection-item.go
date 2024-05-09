package model

type CollectionItemList struct {
	Title       string
	Description string
}

type CollectionItem struct {
	CollectionId string
	Timestamp    string
	List         []CollectionItemList
}
