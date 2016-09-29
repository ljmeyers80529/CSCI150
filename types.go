package csci150

type dictionaryUserName struct {      // name type to be read from dictionary of usernames from datastore 
    Name, UUID string
}

type userInformationType struct {   // type to contain user information and preferences.
	UserID   string
	Name     string
	Password string
	Username string
    Timezone int
    DST      bool
	loggedIn bool
}