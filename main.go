package main

import (
	"chaos"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir("./Frontend")))

	//獲取
	http.HandleFunc("/tinkerbell", tinkerbell)
	http.HandleFunc("/duffing", duffing)
	http.HandleFunc("/henon", henon)

	fmt.Println("Server start on 127.0.0.1:8000")
	http.ListenAndServe(":8000", nil)
}

func tinkerbell(w http.ResponseWriter, r *http.Request) {
	l := chaos.NewList()
	l.TinkerBell()

	data, err := json.Marshal(l.GetData())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func duffing(w http.ResponseWriter, r *http.Request) {
	l := chaos.NewList()
	l.Duffing()

	data, err := json.Marshal(l.GetData())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func henon(w http.ResponseWriter, r *http.Request) {
	l := chaos.NewList()
	l.Henon()

	data, err := json.Marshal(l.GetData())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
