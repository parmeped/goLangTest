package repository

var authorised = []User{
	User{"TestUser", "pass"},
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
