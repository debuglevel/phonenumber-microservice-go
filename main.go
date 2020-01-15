package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/nyaruka/phonenumbers"
)

// func post(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusCreated)
// 	w.Write([]byte(`{"message": "post called"}`))
// }

type Phonenumber struct {
	Phonenumber string
}

type FormattedPhonenumber struct {
	Phonenumber          string
	FormattedPhonenumber string
}

func post(w http.ResponseWriter, r *http.Request) {
	log.Println("Got POST request")

	var phonenumber Phonenumber

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&phonenumber)
	if err != nil {
		log.Println("Decoding HTTP body failed")

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println("Parsing and Formatting " + phonenumber.Phonenumber)

	// parse our phone number
	num, err := phonenumbers.Parse(phonenumber.Phonenumber, "DE")

	formattedNum := addBrackets(phonenumbers.Format(num, phonenumbers.INTERNATIONAL))

	formattedPhonenumber := FormattedPhonenumber{phonenumber.Phonenumber, formattedNum}

	log.Println("Returning JSON")
	// Do something with the Person struct...
	json.NewEncoder(w).Encode(formattedPhonenumber)
	//fmt.Fprintf(w, "%+v", formattedPhonenumber)
}

func addBrackets(phonenumber string) string {
	splits := strings.Split(phonenumber, " ")

	bracketed := splits[0] + " (" + splits[1] + ") " + splits[2]

	return bracketed
}

// func params(w http.ResponseWriter, r *http.Request) {
// 	pathParams := mux.Vars(r)
// 	w.Header().Set("Content-Type", "application/json")

// 	userID := -1
// 	var err error
// 	if val, ok := pathParams["userID"]; ok {
// 		userID, err = strconv.Atoi(val)
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			w.Write([]byte(`{"message": "need a number"}`))
// 			return
// 		}
// 	}

// 	commentID := -1
// 	if val, ok := pathParams["commentID"]; ok {
// 		commentID, err = strconv.Atoi(val)
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			w.Write([]byte(`{"message": "need a number"}`))
// 			return
// 		}
// 	}

// 	query := r.URL.Query()
// 	location := query.Get("location")

// 	w.Write([]byte(fmt.Sprintf(`{"userID": %d, "commentID": %d, "location": "%s" }`, userID, commentID, location)))
// }

func main() {
	log.Println("Starting Golang phonenumber-microservice...")

	router := mux.NewRouter()

	formatRouter := router.PathPrefix("/format").Subrouter()
	formatRouter.HandleFunc("/", post).Methods(http.MethodPost)

	//api.HandleFunc("/user/{userID}/comment/{commentID}", params).Methods(http.MethodGet)

	err := http.ListenAndServe(":80", router)

	log.Fatal(err)
}
