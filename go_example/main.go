package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

var tasks = []Task{
	{ID: 1, Title: "買牛奶", Description: "去超市買一瓶牛奶", Done: false},
	{ID: 2, Title: "學習 Go", Description: "完成 RESTful API 教學", Done: false},
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskID, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, task := range tasks {
		if task.ID == taskID {
			json.NewEncoder(w).Encode(task)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/tasks/{id}", GetTask).Methods("GET")

	fmt.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
