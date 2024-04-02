package entities

// This structure is dedicated to injsonation about users.
type User struct {
	UserId          int    `json:"userId"`
	PathToUserImage string `json:"pathToUserImage"`
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	MiddleName      string `json:"middleName"`
}

// This structure is dedicated to creating users.
type CreateUser struct {
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
	PhoneNumber string `json:"phoneNumber"`
	FirstName   string `json:"firstName" binding:"required"`
	LastName    string `json:"lastName" binding:"required"`
	MiddleName  string `json:"middleName" binding:"required"`
	Biography   string `json:"biography"`
}

// This structure is dedicated to authentication and authorization.
type LoginUser struct {
	UserId   int
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     int
}
