package database

type Database struct {
	dict map[string]string
}

func (db *Database) Get(key string) (value string, ok bool) {
	val, ok := db.dict[key]
	return val, ok
}

func (db *Database) Set(key string, value string) {
	db.dict[key] = value
}

func NewDatabase() *Database {
	return &Database{dict: make(map[string]string)}
}
