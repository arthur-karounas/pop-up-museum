package entities

// This structure is dedicated to compare status codes and their localizations.
type Status struct {
	StatusCode int    `json:"statusCode"`
	Locale     string `json:"locale"`
}
