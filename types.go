package csci150

type dictionaryUserName struct { // name type to be read from dictionary of usernames from datastore
	Name, UUID string
}

// countdown timer implementation definition.
type upComming struct {
	Title                                string
	ID, Year, Month, Day, Hours, Minutes int
}

// list of top rated movies.
type topRatedPop struct {
	Title  string
	ID     int
	Rating float32
}

// definition for maovie / tv / game detail.
type movieTvGameInformation struct {
	ID                    int
	Image, Description, ReleaseDate    string
	TVSeasons, TVEpisodes int
	Genres               []string
	UID					int
}

// user's favorites / watch list.
type watches struct {
	ID int32
}

// all upcomming movies list.
type cdUpcomming []upComming

// all top or popular rated movies.
type topPopRated []topRatedPop

// type to contain user information and preferences.
type userInformationType struct {
	UserID      string
	Name        string
	Password    string
	Username    string
	Timezone    int
	DST         bool
	LoggedIn    bool
	Counters    cdUpcomming
	Top, Pop    topPopRated
	Watched     []watches
	MovieTvGame movieTvGameInformation
}
