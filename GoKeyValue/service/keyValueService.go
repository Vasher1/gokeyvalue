package service

import (
	"errors"
	"sync"
)

var (
	store map[string]interface{}
	mutex sync.RWMutex
)

func Initialise() {
	store = make(map[string]interface{})
}

func Add(key string, value interface{}) (err error) {
	mutex.RLock()
	var _, exists = store[key]
	mutex.RUnlock()

	if exists {
		err = errors.New("Key already exists")
		//fmt.Printf("Key already exists\n")
	} else {
		mutex.Lock()
		store[key] = value
		mutex.Unlock()
		//fmt.Printf("Added : %v, %v\n", key, value)
	}

	return
}

func Read(key string) (err error, value interface{}) {
	mutex.RLock()
	value, exists := store[key]
	mutex.RUnlock()

	if exists {
		//fmt.Printf("Value : %v\n", value)
	} else {
		err = errors.New("Key not found")
		//fmt.Printf("Key not found\n")
	}

	return
}

func ReadAll() (err error, returnStore map[string]interface{}) {
	mutex.RLock()
	var length = len(store)
	mutex.RUnlock()

	if length <= 0 {
		err = errors.New("Store is empty")
		//fmt.Printf("Store is empty\n")
	}
	returnStore = store

	return
}

func Remove(key string) (err error) {
	mutex.RLock()
	var _, exists = store[key]
	mutex.RUnlock()

	if exists {
		delete(store, key)
		//fmt.Printf("Value deleted\n")
	} else {
		err = errors.New("Value not found")
		//fmt.Printf("Value not found\n")
	}

	return
}
