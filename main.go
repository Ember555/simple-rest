package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"simple-rest/route"
)

// type testData struct {
// 	Name string
// 	ID   string
// }

// type DBdata struct {
// 	User string
// }

// var data []testData

// func SearchAll(w http.ResponseWriter, r *http.Request) {
// 	json.NewEncoder(w).Encode(data)
// }

// func SearchByName(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	for _, item := range data {
// 		if item.Name == vars["name"] {
// 			json.NewEncoder(w).Encode(item)
// 			return
// 		}
// 	}
// 	json.NewEncoder(w).Encode("Name not match")
// }

// func SearchByID(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	for _, item := range data {
// 		if item.ID == vars["id"] {
// 			json.NewEncoder(w).Encode(item)
// 			return
// 		}
// 	}
// 	json.NewEncoder(w).Encode("ID not match")
// }

// func CreateData(w http.ResponseWriter, r *http.Request) {
// 	var newData testData
// 	_ = json.NewDecoder(r.Body).Decode(&newData)
// 	if newData.Name == "" || newData.ID == "" {
// 		json.NewEncoder(w).Encode("invalid input")
// 		return
// 	}
// 	data = append(data, newData)
// 	json.NewEncoder(w).Encode(data)
// }

// func RemoveDataByName(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	for index, item := range data {
// 		if item.Name == vars["name"] {
// 			data = append(data[:index], data[index+1:]...)
// 			json.NewEncoder(w).Encode(data)
// 			return
// 		}
// 	}
// 	json.NewEncoder(w).Encode("Name not match")
// }

func main() {

	// router := mux.NewRouter()
	// data = append(data, testData{Name: "initName", ID: "initID"})
	// data = append(data, testData{Name: "initName2", ID: "initID2"})
	// data = append(data, testData{Name: "initName3", ID: "initID3"})

	http.HandleFunc("/searchPar", route.SearchParticipant)
	// router.HandleFunc("/searchAll", SearchAll).Methods("GET")
	// router.HandleFunc("/search/name/{name}", SearchByName).Methods("GET")
	// router.HandleFunc("/search/id/{id}", SearchByID).Methods("GET")
	// router.HandleFunc("/create", CreateData).Methods("POST")
	// router.HandleFunc("/remove/{name}", RemoveDataByName).Methods("POST")

	log.Print("The service is ready to listen and serve.")
	log.Fatal(http.ListenAndServe(":8085", nil))
}
