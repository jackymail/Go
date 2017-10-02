package main

import (
      "encoding/json"
      "log"
      "net/http"

      "github.com/gorilla/mux"
)

type Address struct{
    City string `json:"city,omitempty"`
    State string `json:"state,omitempty"`
}


type Person struct{
  ID string `json:"id,omitempty"`
  Firstname string `json:"firstname,omitempty"`
  Lastname  string `json::"lastname,omitempty"`
  Address *Address  `json:address,omitempty`
}


var people []Person

func GetPersonEndPoint(w http.ResponseWriter,req *http.Request) {
      params := mux.Vars(req)
      for _, item := range people {
        if item.ID == params["id"]{
        json.NewEncoder(w).Encode(item)
        return
      }
      }
      json.NewEncoder(w).Encode(&Person{})
}
func GetPeopleEndPoint(w http.ResponseWriter,req *http.Request) {
    json.NewEncoder(w).Encode(people)

}
func CreatePersonEndPoint(w http.ResponseWriter,req *http.Request) {
      params := mux.Vars(req)
      var person Person

      _ = json.NewDecoder(req.Body).Decode(&person)
      person.ID = params["id"]
      people = append(people,person)
      json.NewEncoder(w).Encode(people)
}

func DeletePersonEndPoint(w http.ResponseWriter,req *http.Request) {
      params := mux.Vars(req)
      for index ,item := range people{
          if item.ID == params["id"]{
          people = append(people[:index],people[index+1:]...)
          break
        }
      }
      json.NewEncoder(w).Encode(people)

}

func main() {
   router := mux.NewRouter()
   people = append(people,Person{ID:"1",Firstname:"Nic",Lastname:"Robby",Address:&Address{City:"Dublin",State:"c"}})
   people = append(people,Person{ID:"2",Firstname:"Jerry",Lastname:"Michael",Address:&Address{City:"NewYork",State:"US"}})

   router.HandleFunc("/people",GetPeopleEndpoint).Methods("GET")
   router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
   router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
   router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")
   log.Fatal(http.ListenAndServe(":12345",router))

}
