package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"keyvaluepairexercise/service"
	"net/http"
	"strings"
)

type keyvalueHttpController struct{}

func (controller keyvalueHttpController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/keyvalue" {
		switch r.Method {
		case http.MethodGet:
			getAll(w, r)
		case http.MethodPut:
			put(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		id := strings.TrimPrefix(r.URL.Path, "/keyvalue/")

		switch r.Method {
		case http.MethodGet:
			get(id, w)
		case http.MethodDelete:
			delete(id, w)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	}
}

func newKeyValueHttpController() *keyvalueHttpController {
	service.Initialise()

	return &keyvalueHttpController{}
}

func getAll(w http.ResponseWriter, r *http.Request) {
	err, store := service.ReadAll()

	if handlePotentialError(err, w) {
		return
	}

	encodeResponseAsJSON(store, w)
}

func get(key string, w http.ResponseWriter) {
	err, value := service.Read(key)
	if handlePotentialError(err, w) {
		return
	}

	encodeResponseAsJSON(value, w)
}

func put(w http.ResponseWriter, r *http.Request) {
	request, err := parseRequest(r)
	if handlePotentialError(err, w) {
		return
	}

	err = service.Add(request.Key, request.Value)
	if handlePotentialError(err, w) {
		return
	}

	w.WriteHeader(http.StatusOK)
}

func delete(key string, w http.ResponseWriter) {
	err := service.Remove(key)

	if handlePotentialError(err, w) {
		return
	}

	w.WriteHeader(http.StatusOK)
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}

func parseRequest(r *http.Request) (AddRequest, error) {
	var parsedRequest AddRequest

	data, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(data, &parsedRequest)

	if err != nil {
		return parsedRequest, err
	}

	return parsedRequest, nil
}

func handlePotentialError(err error, w http.ResponseWriter) (errorFound bool) {
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))

		return true
	}

	return false
}

type AddRequest struct {
	Key   string
	Value interface{}
}
