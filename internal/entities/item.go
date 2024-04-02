package entities

// This structure is dedicated to detailed injsonation about items.
type Item struct {
	ItemId          int     `json:"itemId"`
	AuthorId        int     `json:"authorId"`
	ExhibitId       int     `json:"exhibitId"`
	ItemName        string  `json:"itemName"`
	PathToItemImage string  `json:"pathToItemImage"`
	Technique       string  `json:"technique"`
	Height          float64 `json:"height"`
	Length          float64 `json:"length"`
	Status          int     `json:"status"`
	Created         int64   `json:"created"`
}

// This structure is dedicated to short injsonation about items.
type AllItems struct {
	ItemId          int    `json:"itemId"`
	AuthorId        int    `json:"authorId"`
	ItemName        string `json:"itemName"`
	PathToItemImage string `json:"pathToItemImage"`
}

// This structure is dedicated to creating items.
type CreateItem struct {
	AuthorId  int     `json:"authorId" binding:"required"`
	ExhibitId int     `json:"exhibitId"`
	ItemName  string  `json:"itemName" binding:"required"`
	Technique string  `json:"technique" binding:"required"`
	Height    float64 `json:"height" binding:"required"`
	Length    float64 `json:"length" binding:"required"`
	Status    int     `json:"status" binding:"required"`
	Created   int64   `json:"created" binding:"required"`
}

// This structure is dedicated to updating items.
type UpdateItem struct {
	ItemId    int
	AuthorId  int     `json:"authorId" binding:"required"`
	ExhibitId int     `json:"exhibitId"`
	ItemName  string  `json:"itemName" binding:"required"`
	Technique string  `json:"technique" binding:"required"`
	Height    float64 `json:"height" binding:"required"`
	Length    float64 `json:"length" binding:"required"`
	Status    int     `json:"status" binding:"required"`
	Created   int64   `json:"created" binding:"required"`
}

// This structure is dedicated to deleting items.
type DeleteItem struct {
	ItemId      int
	IsRemoved   bool `json:"isRemoved"`
	DateRemoved int64
}
