package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
)

var people = make([]Person, 0)
var lastID = 0
var notFoundResponse = Response{Status: "ERROR", Message: "Not found"}
var deletedResponse = Response{Status: "OK", Message: "Deleted"}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/people", CreateHandler(GetPeople)).Methods("GET")
	router.HandleFunc("/people", CreateHandler(CreatePerson)).Methods("POST")
	router.HandleFunc("/people/{id}", CreateHandler(EditPerson)).Methods("PUT")
	router.HandleFunc("/people/{id}", CreateHandler(GetPerson)).Methods("GET")
	router.HandleFunc("/people/{id}", CreateHandler(DeletePerson)).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":1337", router))
}

func CreateHandler(impl func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		impl(w, r)
	}
}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range people {
		var id, _ = strconv.Atoi(params["id"])
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func EditPerson(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	var id, _ = strconv.Atoi(params["id"])
	var newPerson Person
	_ = json.NewDecoder(request.Body).Decode(&newPerson)
	newPerson.ID = id
	for index, item := range people {
		if item.ID == id {
			people = append(append(people[:index], newPerson), people[index+1:]...)
			json.NewEncoder(writer).Encode(newPerson)
			return
		}
	}
	json.NewEncoder(writer).Encode(notFoundResponse)
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = nextID()
	people = append(people, person)
	json.NewEncoder(w).Encode(person)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id, _ = strconv.Atoi(params["id"])
	for index, item := range people {
		if item.ID == id {
			people = append(people[:index], people[index+1:]...)
			json.NewEncoder(w).Encode(deletedResponse)
			return
		}
	}
	json.NewEncoder(w).Encode(notFoundResponse)
}

func nextID() int {
	lastID++
	return lastID
}
