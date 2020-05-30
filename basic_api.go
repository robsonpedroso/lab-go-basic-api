package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Account struct {
	Id        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Document  string   `json:"document,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City       string `json:"city,omitempty"`
	State      string `json:"state,omitempty"`
	ZipCode    string `json:"zip_code,omitempty"`
	Street     string `json:"street,omitempty"`
	District   string `json:"district,omitempty"`
	Complement string `json:"complement,omitempty"`
	Number     string `json:"number,omitempty"`
}

var accounts []Account

func main() {
	router := mux.NewRouter()

	initMoqs()

	router.HandleFunc("/account", GetAccount).Methods("GET")
	router.HandleFunc("/account/{id}", GetAccountById).Methods("GET")
	router.HandleFunc("/account/{id}", CreateAccount).Methods("POST")
	router.HandleFunc("/account/{id}", RemoveAccount).Methods("DELETE")

	port := ":3000"

	log.Println(fmt.Sprintf("Sever in http://localhost%s", port))
	log.Println("Routers")

	err := router.Walk(gorillaWalkFn)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.ListenAndServe(port, router))
}

func initMoqs() {
	accounts = append(accounts,
		Account{
			Id:        "1",
			Firstname: "Robson",
			Lastname:  "Pedroso",
			Document:  "00000000191",
			Address: &Address{
				City:     "Itatiba",
				State:    "SP",
				Street:   "Rua Apenas para Teste",
				District: "Centro",
				Number:   "123",
				ZipCode:  "12345678",
			},
		})

	accounts = append(accounts,
		Account{
			Id:        "2",
			Firstname: "Usuário",
			Lastname:  "Teste",
			Document:  "00000000272",
			Address: &Address{
				City:     "São Paulo",
				State:    "SP",
				Street:   "Rua Outro Teste",
				District: "Centro",
				Number:   "321",
				ZipCode:  "12345100",
			},
		})
}

func GetAccount(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(accounts)
}

func GetAccountById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range accounts {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Account{})
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var account Account
	_ = json.NewDecoder(r.Body).Decode(&account)
	account.Id = params["id"]
	accounts = append(accounts, account)
	json.NewEncoder(w).Encode(accounts)
}

func RemoveAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range accounts {
		if item.Id == params["id"] {
			accounts = append(accounts[:index], accounts[index+1:]...)
			break
		}

		json.NewEncoder(w).Encode(accounts)
	}
}

func gorillaWalkFn(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
	path, _ := route.GetPathTemplate()
	method, _ := route.GetMethods()

	log.Println(method, path)
	return nil
}
