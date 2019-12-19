package repository

type DataReader interface {
	Read() []string
}

type DB struct {
	Tables [1]Table
}

type Table struct {
	Name string
	Data [3]string
}

func New() *DB {
	return initializer()
}

func initializer() *DB {
	array := [3]string{"test1", "test2", "test3"}
	var table = Table{"testing", array}
	var tables [1]Table
	tables[0] = table

	var db = DB{tables}
	return &db
}

func Read() []string {
	var readTest []string
	readTest[0] = "Are"
	readTest[1] = "You"
	readTest[2] = "Reading"
	readTest[3] = "This"

	return readTest
}
