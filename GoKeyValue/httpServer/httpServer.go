package httpServer

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"keyvaluepairexercise/keyValueService"
	"net/http"
	"strings"
)

func HttpListen() {
	http.HandleFunc("/keyvalue", handler)
	http.HandleFunc("/keyvalue/", pathHandler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getAll(w, r)
	case http.MethodPut:
		put(w, r)
	default:
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
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

// Methods

func getAll(w http.ResponseWriter, r *http.Request) {
	err, store := keyValueService.ReadAll()

	if handlePotentialError(err, w) {
		return
	}

	encodeResponseAsJSON(store, w)
}

func get(key string, w http.ResponseWriter) {
	err, value := keyValueService.Read(key)
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

	err = keyValueService.Add(request.Key, request.Value)
	if handlePotentialError(err, w) {
		return
	}

	w.WriteHeader(http.StatusOK)
}

func delete(key string, w http.ResponseWriter) {
	err := keyValueService.Remove(key)

	if handlePotentialError(err, w) {
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Helpers

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}

func parseRequest(r *http.Request) (addRequest, error) {
	var parsedRequest addRequest

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

// Structs

type addRequest struct {
	Key   string
	Value interface{}
}
