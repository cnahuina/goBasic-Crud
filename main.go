package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//Creamos la entidad Persona
//omitempty que no me traiga en blanco
type Person struct {
	ID string `json:"id,omitempty"`
	FirstName string `json:"firstname,omitempty"`
	LastName string `json:"lastname,omitempty"`
	Address *Address `json:"address,omitempty"`
}

type Address struct {
	City string `json:"city"`
	State string `json:"state"`
}


var people []Person

func GetPeopleEndPoint(w http.ResponseWriter, req *http.Request){
	json.NewEncoder(w).Encode(people)
}
func GetPersonEndPoint(w http.ResponseWriter, req *http.Request){

}
func CreatePersonEndPoint(w http.ResponseWriter, r *http.Request){

}
func DeletePersonEndPoint(w http.ResponseWriter, r *http.Request){

}

func main () {

	//created and asignated
	router := mux.NewRouter()

	people = append(people, Person{ID:"1",FirstName: "Carlos" , LastName: "Garcia",Address: &Address{City: "Lima",State:"California"}})
	people = append(people, Person{ID:"2",FirstName: "Emmanuel" , LastName: "Garcia"})

	//list endPoints
	router.HandleFunc("/people", GetPeopleEndPoint).Methods("GET")

	router.HandleFunc("/people/{id}", GetPersonEndPoint).Methods("GET")

	router.HandleFunc("/people/{id}", CreatePersonEndPoint).Methods("POST")

	router.HandleFunc("/people/{id}", DeletePersonEndPoint).Methods("DELETE")

	//Escuchar al servidor , anteponiendo un log por si ocurre algun error en el servidor cuando inicie
	log.Fatal(http.ListenAndServe(":3000", router))
}

