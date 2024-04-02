package entities

// This structure is dedicated to detailed injsonation about reservations.
type Reservation struct {
	ReservationId int    `json:"reservationId"`
	UserId        int    `json:"userId"`
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	MiddleName    string `json:"middleName"`
	Date          int64  `json:"date"`
	AutoBrand     string `json:"autoBrand"`
	AutoNumber    string `json:"autoNumber"`
	Status        int    `json:"status"`
}

// This structure is dedicated to creating reservations.
type CreateReservation struct {
	UserId     int
	FirstName  string `json:"firstName" binding:"required"`
	LastName   string `json:"lastName" binding:"required"`
	MiddleName string `json:"middleName" binding:"required"`
	Date       int64  `json:"date" binding:"required"`
	AutoBrand  string `json:"autoBrand"`
	AutoNumber string `json:"autoNumber"`
}

// This structure is dedicated to updating reservations.
type UpdateReservation struct {
	ReservationId int
	Status        int `json:"status"`
}
