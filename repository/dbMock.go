package repository

var authorised = []User{
	User{"TestUser",
		"pass",
		[]Message{Message{"TestUser2", "TestUser", "Hi, this is an unread Message"}},
		[]Message{Message{"TestUser2", "TestUser", "Hi, this was already read"}},
	},
	User{"TestUser2",
		"pass2",
		[]Message{Message{"TestUser", "TestUser2", "Hi, this is an unread Message"}},
		[]Message{Message{"TestUser", "TestUser2", "Hi, this was already read"}},
	},
}

// var testArray = []string{
// 	"test1",
// 	"test2",
// 	"test3",
// }

// the mocked data returned
type Repo struct {
	Authorized []User
	TestArray  []string
}

func GetRepo() *Repo {
	repo := Repo{authorised, nil}
	return &repo
}
