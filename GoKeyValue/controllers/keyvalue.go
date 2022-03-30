package controllers

import (
	"encoding/json"
	"io"
	"keyvaluepairexercise/service"
	"net/http"
	"strings"
)

type keyvalueController struct{}

func (controller keyvalueController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/keyvalue" {
		switch r.Method {
		case http.MethodGet:
			getAll(w, r)
		case http.MethodPost:
			post(w, r)
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

func getAll(w http.ResponseWriter, r *http.Request) {
	err, store := service.ReadAll()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	encodeResponseAsJSON(store, w)
}

func get(id string, w http.ResponseWriter) {
	err, value := service.Read(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	encodeResponseAsJSON(value, w)
}

func post(w http.ResponseWriter, r *http.Request) {
	request, err := parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse request object"))
		return
	}

	err = service.Add(request.Key, request.Value)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func delete(id string, w http.ResponseWriter) {
	err := service.Remove(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func newKeyValueController() *keyvalueController {
	service.Initialise()

	return &keyvalueController{}
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}

type AddRequest struct {
	Key   string
	Value interface{}
}

func parseRequest(r *http.Request) (AddRequest, error) {
	dec := json.NewDecoder(r.Body)
	var request AddRequest
	err := dec.Decode(&request)
	if err != nil {
		return AddRequest{}, err
	}
	return request, nil
}
