package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//PersonController handles person
type PersonController struct{}

//NewPersonController creates and returns new person controller
func NewPersonController() *PersonController {
	return &PersonController{}
}

//GetPersonByID gets person by id
func (PersonController) GetPersonByID(rw http.ResponseWriter, req http.Request, params httprouter.Params) {

}

//CreatePerson creates person
func (PersonController) CreatePerson(rw http.ResponseWriter, req http.Request, params httprouter.Params) {

}

//DeletePersonByID deletes person by id
func (PersonController) DeletePersonByID(rw http.ResponseWriter, req http.Request, params httprouter.Params) {

}

//UpdatePerson updates person
func (PersonController) UpdatePerson(rw http.ResponseWriter, req http.Request, params httprouter.Params) {

}
