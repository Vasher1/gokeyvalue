package controllers

import (
	"net/http"
)

func RegisterControllers() {
	controller := newKeyValueController()

	http.Handle("/keyvalue", controller)
	http.Handle("/keyvalue/", controller)
}
