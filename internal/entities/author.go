package entities

// This structure is dedicated to detailed injsonation about authors.
type Author struct {
	AuthorId          int    `json:"authorId"`
	AuthorName        string `json:"authorName"`
	Description       string `json:"description"`
	PathToAuthorImage string `json:"pathToAuthorImage"`
	Pseudonym         string `json:"pseudonym"`
	PhoneNumber       string `json:"phoneNumber"`
	Email             string `json:"email"`
}

// This structure is dedicated to short injsonation about authors.
type AllAuthors struct {
	AuthorId          int    `json:"authorId"`
	AuthorName        string `json:"authorName"`
	PathToAuthorImage string `json:"pathToAuthorImage"`
}

// This structure is dedicated to creating authors.
type CreateAuthor struct {
	AuthorName  string `json:"authorName" binding:"required"`
	Description string `json:"description" binding:"required"`
	Pseudonym   string `json:"pseudonym" binding:"required"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
	Email       string `json:"email" binding:"required"`
}

// This structure is dedicated to updating authors.
type UpdateAuthor struct {
	AuthorId    int
	AuthorName  string `json:"authorName" binding:"required"`
	Description string `json:"description" binding:"required"`
	Pseudonym   string `json:"pseudonym" binding:"required"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
	Email       string `json:"email" binding:"required"`
}

// This structure is dedicated to deleting authors.
type DeleteAuthor struct {
	AuthorId    int
	IsRemoved   bool `json:"isRemoved"`
	DateRemoved int64
}
