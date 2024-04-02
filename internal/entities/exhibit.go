package entities

// This structure is dedicated to detailed injsonation about exhibits.
type Exhibit struct {
	ExhibitId          int    `json:"exhibitId"`
	ExhibitName        string `json:"exhibitName"`
	PathToExhibitImage string `json:"pathToExhibitImage"`
	StartDate          int64  `json:"startDate"`
	EndDate            int64  `json:"endDate"`
	Type               string `json:"type"`
	Status             int    `json:"status"`
	Description        string `json:"description"`
	Location           string `json:"location"`
	Website            string `json:"website"`
}

// This structure is dedicated to short injsonation about exhibits.
type AllExhibits struct {
	ExhibitId          int    `json:"exhibitId"`
	ExhibitName        string `json:"exhibitName"`
	PathToExhibitImage string `json:"pathToExhibitImage"`
	StartDate          int64  `json:"startDate"`
	EndDate            int64  `json:"endDate"`
	Status             int    `json:"status"`
}

// This structure is dedicated to creating exhibits.
type CreateExhibit struct {
	ExhibitName string `json:"exhibitName" binding:"required"`
	StartDate   int64  `json:"startDate" binding:"required"`
	EndDate     int64  `json:"endDate" binding:"required"`
	Type        string `json:"type" binding:"required"`
	Status      int    `json:"status" binding:"required"`
	Description string `json:"description" binding:"required"`
	Location    string `json:"location"`
	Website     string `json:"website"`
	Authors     []int  `json:"authors"`
	Items       []int  `json:"items"`
}

// This structure is dedicated to updating exhibits.
type UpdateExhibit struct {
	ExhibitId   int
	ExhibitName string `json:"exhibitName" binding:"required"`
	StartDate   int64  `json:"startDate" binding:"required"`
	EndDate     int64  `json:"endDate" binding:"required"`
	Type        string `json:"type" binding:"required"`
	Status      int    `json:"status" binding:"required"`
	Description string `json:"description" binding:"required"`
	Location    string `json:"location"`
	Website     string `json:"website"`
	Authors     []int  `json:"authors"`
	Items       []int  `json:"items"`
}

// This structure is dedicated to deleting exhibits.
type DeleteExhibit struct {
	ExhibitId   int
	IsRemoved   bool `json:"isRemoved"`
	DateRemoved int64
}
