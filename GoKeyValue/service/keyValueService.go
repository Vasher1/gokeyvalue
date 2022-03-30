package service

import (
	"errors"
	"fmt"
)

var (
	store map[string]interface{}
)

func Initialise() {
	store = make(map[string]interface{})
}

func Add(key string, value interface{}) (err error) {
	var _, exists = store[key]

	if exists {
		err = errors.New("Key already exists")
		fmt.Printf("Key already exists\n")
	} else {
		store[key] = value
		fmt.Printf("Added : %v, %v\n", key, value)
	}

	return
}

func Read(key string) (err error, value interface{}) {
	value, exists := store[key]

	if exists {
		fmt.Printf("Value : %v\n", value)
	} else {
		err = errors.New("Key not found")
		fmt.Printf("Key not found\n")
	}

	return
}

func ReadAll() (err error, returnStore map[string]interface{}) {
	if len(store) <= 0 {
		err = errors.New("Store is empty")
		fmt.Printf("Store is empty\n")
	}

	returnStore = store

	return
}

func Remove(key string) (err error) {
	var _, exists = store[key]

	if exists {
		delete(store, key)
		fmt.Printf("Value deleted\n")
	} else {
		err = errors.New("Value not found")
		fmt.Printf("Value not found\n")
	}

	return
}
