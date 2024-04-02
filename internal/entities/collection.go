package entities

// This structure is dedicated to creating and deleting authors.
type Collection struct {
	Id     int `json:"id"`
	UserId int `json:"userId"`
	ItemId int `json:"itemId" binding:"required"`
}
