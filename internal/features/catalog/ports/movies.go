package ports

type MovieListItem struct {
	ID             int64
	Title          string
	ReleaseYear    *int32
	Rating         *float64
	PosterFilename *string
	Genre          *string
}
