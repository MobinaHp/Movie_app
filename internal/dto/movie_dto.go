package dto


type CreateMovieRequest struct {
	Title       string   `json:"title" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Genres      []string `json:"genres"`
	ReleaseDate string   `json:"release_date" binding:"required"`
}

type MovieResponse struct {
	ID          uint     `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Genres      []string `json:"genres"`
	ReleaseDate string   `json:"release_date"`
}
