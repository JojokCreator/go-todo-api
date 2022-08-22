package router

import (
	"go-todo/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/task", middleware.GetAllTasks).Methods("GET", "OPTIONS")
	router.HandleFunc("/task", middleware.CreateTask).Methods("POST", "OPTIONS")
	router.HandleFunc("/task/{id}", middleware.TaskComplete).Methods("PUT", "OPTIONS")
	router.HandleFunc("/undoTask/{id}", middleware.UndoTask).Methods("PUT", "OPTIONS")
	router.HandleFunc("/deleteTask/{id}", middleware.DeleteTask).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/deleteAll/", middleware.DeleteAllTasks).Methods("DELETE", "OPTIONS")
	return router
}
