package controller

import (
	"encoding/json"
	"net/http"

	"github.com/avinashga23golearning/model"
	"github.com/avinashga23golearning/persistence"

	"github.com/gorilla/mux"
)

//PersonController handles person
type PersonController struct{}

var personPersistenceManager *persistence.PersonPersistenceManager

func init() {
	personPersistenceManager = persistence.NewPersonPersistenceManager()
}

//NewPersonController creates and returns new person controller
func NewPersonController() *PersonController {
	return &PersonController{}
}

//GetPersonByID gets person by id
func (PersonController) GetPersonByID(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	id := vars["id"]
	person, notFound := personPersistenceManager.GetPersonByID(id)

	if notFound {
		rw.WriteHeader(http.StatusNotFound)
	} else {
		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		json.NewEncoder(rw).Encode(person)
	}
}

//CreatePerson creates person
func (PersonController) CreatePerson(rw http.ResponseWriter, req *http.Request) {
	person := model.Person{}

	json.NewDecoder(req.Body).Decode(&person)
	id := personPersistenceManager.CreatePerson(person)

	rw.Header().Add("Location", "http://localhost:8080/person/"+id)
	rw.WriteHeader(http.StatusCreated)
}

//DeletePersonByID deletes person by id
func (PersonController) DeletePersonByID(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	personPersistenceManager.DeletePerson(id)
	rw.WriteHeader(http.StatusOK)
}

//UpdatePerson updates person
func (PersonController) UpdatePerson(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	var person model.Person
	id := vars["id"]
	json.NewDecoder(req.Body).Decode(&person)
	person.ID = id

	personPersistenceManager.UpdatePerson(person)
	rw.WriteHeader(http.StatusOK)
}
