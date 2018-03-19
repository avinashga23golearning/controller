package controller

import (
	"encoding/json"
	"net/http"

	"github.com/avinashga23golearning/model"
	"github.com/avinashga23golearning/persistence"

	"github.com/julienschmidt/httprouter"
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
func (PersonController) GetPersonByID(rw http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	person := personPersistenceManager.GetPersonByID(id)
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(person)
}

//CreatePerson creates person
func (PersonController) CreatePerson(rw http.ResponseWriter, req *http.Request, params httprouter.Params) {
	person := model.Person{}

	json.NewDecoder(req.Body).Decode(&person)
	id := personPersistenceManager.CreatePerson(person)

	rw.Header().Add("Location", "http://localhost:8080/person/"+id)
	rw.WriteHeader(http.StatusCreated)
}

//DeletePersonByID deletes person by id
func (PersonController) DeletePersonByID(rw http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	personPersistenceManager.DeletePerson(id)
	rw.WriteHeader(http.StatusOK)
}

//UpdatePerson updates person
func (PersonController) UpdatePerson(rw http.ResponseWriter, req *http.Request, params httprouter.Params) {
	var person model.Person
	id := params.ByName("id")
	json.NewDecoder(req.Body).Decode(&person)
	person.ID = id

	personPersistenceManager.UpdatePerson(person)
	rw.WriteHeader(http.StatusOK)
}
