package entities

// This structure is dedicated to injsonation about gallery info.
type Info struct {
	MuseumPhoneNumber    string `json:"museumPhoneNumber"`
	Email                string `json:"email"`
	WorkSchedule         string `json:"workSchedule"`
	Address              string `json:"address"`
	OrganizerPhoneNumber string `json:"organizerPhoneNumber"`
	Website              string `json:"website"`
}

// This structure is dedicated to injsonation about FAQs.
type FAQ struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
