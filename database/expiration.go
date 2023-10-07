package database

import (
	"strconv"
	"time"
)

type ExpirationManager struct {
	keyToExpirationTimeMap   map[string]string
	expirationTimeToKeysMap  map[string][]string
	NotifyExpiredKeysChannel chan string
}

func NewExpirationManager() *ExpirationManager {
	manager := ExpirationManager{keyToExpirationTimeMap: make(map[string]string), expirationTimeToKeysMap: make(map[string][]string), NotifyExpiredKeysChannel: make(chan string)}
	go runExpiredKeysCycle(&manager)
	return &manager
}

func (expiration *ExpirationManager) RegisterNewKeyExpirationState(key string, newExpiration time.Time) {
	currentIndex, exists := expiration.keyToExpirationTimeMap[key]
	if exists {
		keys := expiration.expirationTimeToKeysMap[currentIndex]
		newKeys := make([]string, len(keys))
		for _, v := range keys {
			if v != key {
				newKeys = append(newKeys, v)
			}
		}
		expiration.expirationTimeToKeysMap[currentIndex] = newKeys
	}

	newExpirationIndex := convertExpirationTimeToString(newExpiration)
	keys := expiration.expirationTimeToKeysMap[newExpirationIndex]
	expiration.expirationTimeToKeysMap[newExpirationIndex] = append(keys, key)
}

func (expiration *ExpirationManager) findExpiredKeys() []string {
	currentExpiredIndex := convertExpirationTimeToString(time.Now())
	return expiration.expirationTimeToKeysMap[currentExpiredIndex]
}

func runExpiredKeysCycle(manager *ExpirationManager) {
	for {
		time.Sleep(1 * time.Second)
		expired := manager.findExpiredKeys()
		for _, key := range expired {
			manager.NotifyExpiredKeysChannel <- key
		}
	}
}

// translates time to an indexable string, the resolution of the index is defined here
func convertExpirationTimeToString(time time.Time) string {
	return strconv.FormatInt(time.Unix(), 10)
}
