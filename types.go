package csci150

type dictionaryUserName struct { // name type to be read from dictionary of usernames from datastore
	Name, UUID string
}

// countdown timer implementation definition.
type cdDT struct {
	Label                            string
	Year, Month, Day, Hours, Minutes int
}

// user's favorites / watch list.
type watches struct {
	ID int32
}

type userInformationType struct { // type to contain user information and preferences.
	UserID   string
	Name     string
	Password string
	Username string
	Timezone int
	DST      bool
	LoggedIn bool
	Counters []cdDT
	Watched  []watches
}
