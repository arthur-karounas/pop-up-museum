package entities

// This structure is dedicated to detailed injsonation about appeals.
type Appeal struct {
	AppealId    int    `json:"appealId"`
	UserId      int    `json:"userId"`
	Initials    string `json:"initials"`
	Contact     string `json:"contact"`
	Issue       string `json:"issue"`
	Content     string `json:"content"`
	Status      int    `json:"status"`
	DateCreated int64  `json:"dateCreated"`
}

// This structure is dedicated to creating appeals.
type CreateAppeal struct {
	UserId      int
	Initials    string `json:"initials" binding:"required"`
	Contact     string `json:"contact" binding:"required"`
	Issue       string `json:"issue" binding:"required"`
	Content     string `json:"content" binding:"required"`
	DateCreated int64
}

// This structure is dedicated to updating appeals.
type UpdateAppeal struct {
	AppealId int
	Status   int `json:"status" binding:"required"`
}
