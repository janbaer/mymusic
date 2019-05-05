package model

// ExactSearch - Defines the fields for an exact search
type ExactSearch struct {
	Artist string
	Title  string
}

// NewExactSearch - Creates a new ExactSearch struct
func NewExactSearch(artist string, title string) *ExactSearch {
	return &ExactSearch{
		Artist: artist,
		Title:  title,
	}
}
