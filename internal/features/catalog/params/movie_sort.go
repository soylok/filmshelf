package params

type MovieSort string

const (
	MovieSortIDAsc      MovieSort = "id_asc"
	MovieSortIDDesc     MovieSort = "id_desc"
	MovieSortTitleAsc   MovieSort = "title_asc"
	MovieSortTitleDesc  MovieSort = "title_desc"
	MovieSortYearAsc    MovieSort = "release_year_asc"
	MovieSortYearDesc   MovieSort = "release_year_desc"
	MovieSortRatingAsc  MovieSort = "rating_asc"
	MovieSortRatingDesc MovieSort = "rating_desc"
)
