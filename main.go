package main
//
//import (
//	"encoding/json"
//	"github.com/gorilla/mux"
//	"log"
//	"net/http"
//)
//
////Creamos la entidad Persona
////omitempty que no me traiga en blanco
//type Persona struct {
//	ID string `json:"id,omitempty"`
//	FirstName string `json:"firstname,omitempty"`
//	LastName string `json:"lastname,omitempty"`
//	Address *Address `json:"address,omitempty"`
//}
//
//type Address struct {
//	City string `json:"city"`
//	State string `json:"state"`
//}
//
//
//var people []Persona
//
//func GetPeopleEndPoint(w http.ResponseWriter, req *http.Request){
//	json.NewEncoder(w).Encode(people)
//}
//func GetPersonEndPoint(w http.ResponseWriter, req *http.Request){
//	params := mux.Vars(req)
//	for _ , item := range people{
//		if item.ID == params ["id"]{
//			json.NewEncoder(w).Encode(item)
//			return
//		}
//	}
//	json.NewEncoder(w).Encode(&Persona{})
//}
//func CreatePersonEndPoint(w http.ResponseWriter, req *http.Request){
//	params := mux.Vars(req)
//	var person Persona
//	_ = json.NewDecoder(req.Body).Decode(&person)
//	person.ID = params["id"]
//	people = append(people,person)
//	json.NewEncoder(w).Encode(people)
//
//}
//func DeletePersonEndPoint(w http.ResponseWriter, req *http.Request){
//
//	params := mux.Vars(req)
//	for index, item := range people{
//		if item.ID == params["id"]{
//			people= append(people[:index],people[index + 1:]...)
//			break
//		}
//	}
//	json.NewEncoder(w).Encode(people)
//
//}
//
//func main () {
//
//	//created and asignated
//	router := mux.NewRouter()
//
//	people = append(people, Persona{ID:"1",FirstName: "Carlos" , LastName: "Garcia",Address: &Address{City: "Lima",State:"California"}})
//	people = append(people, Persona{ID:"2",FirstName: "Emmanuel" , LastName: "Garcia"})
//
//	//list endPoints
//	router.HandleFunc("/people", GetPeopleEndPoint).Methods("GET")
//
//	router.HandleFunc("/people/{id}", GetPersonEndPoint).Methods("GET")
//
//	router.HandleFunc("/people/{id}", CreatePersonEndPoint).Methods("POST")
//
//	router.HandleFunc("/people/{id}", DeletePersonEndPoint).Methods("DELETE")
//
//	//Escuchar al servidor , anteponiendo un log por si ocurre algun error en el servidor cuando inicie
//	log.Fatal(http.ListenAndServe(":3004", router))
//}
//
