package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"merchandise_control_system/config"
	"merchandise_control_system/models"
	"net/http"
	"strconv"
)

type DeleteResponse struct {
	Id string `json:"id"`
}

func rootPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Go Api Server")
	fmt.Println("Root endpoint is hooked!")
}

func fetchAllItems(w http.ResponseWriter, r *http.Request) {
	var items []models.Item
	models.GetAllItems(&items)
	responseBody, err := json.Marshal(items)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8001")
	if r.Method == http.MethodOptions {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}

func fetchSingleItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var item models.Item
	models.GetSingleItem(&item, id)
	responseBody, err := json.Marshal(item)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8001")
	if r.Method == http.MethodOptions {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}

func createItem(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var item models.Item
	if err := json.Unmarshal(reqBody, &item); err != nil {
		log.Fatal(err)
	}
	models.InsertItem(&item)
	responseBody, err := json.Marshal(item)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8001")
	if r.Method == http.MethodOptions {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	models.DeleteItem(id)
	responseBody, err := json.Marshal(DeleteResponse{Id: id})
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8001")
	if r.Method == http.MethodOptions {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	reqBody, _ := ioutil.ReadAll(r.Body)

	var updateItem models.Item
	if err := json.Unmarshal(reqBody, &updateItem); err != nil {
		log.Fatal(err)
	}
	models.UpdateItem(&updateItem, id)
	convertUintId, _ := strconv.ParseUint(id, 10, 64)
	updateItem.Model.ID = uint(convertUintId)
	responseBody, err := json.Marshal(updateItem)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8001")
	if r.Method == http.MethodOptions {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}

func StartWebServer() error {
	fmt.Println("Rest API with Mux Routers")
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", rootPage)
	r.HandleFunc("/items", fetchAllItems).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/item/{id}", fetchSingleItem).Methods(http.MethodGet, http.MethodOptions)

	r.HandleFunc("/item", createItem).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/item/{id}", deleteItem).Methods(http.MethodDelete, http.MethodOptions)
	r.HandleFunc("/item/{id}", updateItem).Methods(http.MethodPut, http.MethodOptions)
	r.Use(mux.CORSMethodMiddleware(r))

	return http.ListenAndServe(fmt.Sprintf(":%d", config.Config.ServerPort), r)
}
