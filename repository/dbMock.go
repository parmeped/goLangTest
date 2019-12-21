package repository

var authorised = []User{
	User{"TestUser",
		"pass",
		[]Message{Message{"unAuthUser", "TestUser", "Hi, this is an unread Auth Message"}},
		[]Message{Message{"unAuthUser", "TestUser", "Hi, this was already auth read"}},
	},
	User{"unAuthUser",
		"superPass",
		[]Message{Message{"TestUser", "unAuthUser", "Hi, this is an unread Message"}},
		[]Message{Message{"TestUser", "unAuthUser", "Hi, this was already read"}},
	},
}

var testArray = []string{
	"test1",
	"test2",
	"test3",
}

// the mocked data returned
type Repo struct {
	Authorized []User
	TestArray  []string
}

func GetRepo() *Repo {
	repo := Repo{authorised, testArray}
	return &repo
}
