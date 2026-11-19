package posters

type PosterStorage struct {
	postersDir string
}

func NewPosterStorage(postersDir string) *PosterStorage {
	return &PosterStorage{
		postersDir: postersDir,
	}
}
