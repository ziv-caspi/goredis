package database

import (
	"fmt"
	"time"
)

type Database struct {
	dict              map[string]string
	expirationManager *ExpirationManager
}

func (db *Database) Get(key string) (value string, ok bool) {
	val, ok := db.dict[key]
	return val, ok
}

func (db *Database) Set(key string, value string, ttl time.Duration) {
	db.dict[key] = value
	if ttl != 0 {
		db.expirationManager.RegisterNewKeyExpirationState(key, time.Now().Add(ttl))
	}
}

func NewDatabase() *Database {

	expiration := NewExpirationManager()
	database := Database{dict: make(map[string]string), expirationManager: expiration}
	removeExpiredKeys := func() {
		for {
			key := <-expiration.NotifyExpiredKeysChannel
			delete(database.dict, key)
			fmt.Println("removed expired key:", key)
		}
	}
	go removeExpiredKeys()
	return &database
}
