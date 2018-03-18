package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/avinashga23golearning/model"

	"github.com/julienschmidt/httprouter"
)

//PersonController handles person
type PersonController struct{}

var personMap map[int]model.Person
var counter int

func init() {
	personMap = make(map[int]model.Person)
}

//NewPersonController creates and returns new person controller
func NewPersonController() *PersonController {
	return &PersonController{}
}

//GetPersonByID gets person by id
func (PersonController) GetPersonByID(rw http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id, _ := strconv.Atoi(params.ByName("id"))
	if person, exists := personMap[id]; exists {
		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		json.NewEncoder(rw).Encode(person)
	} else {
		rw.WriteHeader(http.StatusNotFound)
	}
}

//CreatePerson creates person
func (PersonController) CreatePerson(rw http.ResponseWriter, req *http.Request, params httprouter.Params) {
	person := model.Person{}

	json.NewDecoder(req.Body).Decode(&person)
	counter++
	person.ID = counter
	personMap[counter] = person

	rw.Header().Add("Location", "http://localhost:8080/person/"+strconv.Itoa(person.ID))
	rw.WriteHeader(http.StatusCreated)
}

//DeletePersonByID deletes person by id
func (PersonController) DeletePersonByID(rw http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id, _ := strconv.Atoi(params.ByName("id"))
	delete(personMap, id)
	rw.WriteHeader(http.StatusOK)
}

//UpdatePerson updates person
func (PersonController) UpdatePerson(rw http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id, _ := strconv.Atoi(params.ByName("id"))
	if person, exists := personMap[id]; exists {
		json.NewDecoder(req.Body).Decode(&person)
		person.ID = id
		personMap[id] = person

		rw.WriteHeader(http.StatusOK)
	} else {
		rw.WriteHeader(http.StatusPreconditionFailed)
	}

}
