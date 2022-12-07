package router

import (
	"golang-react-todo/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router{
	router := mux.NewRouter()
     router.HandleFunc("/api/task", middleware.GetAllTasks).Methods("GET", "OPTIONS")
      router.HandleFunc("/api/task", middleware.CreateTasks).Methods("POST", "OPTIONS")
	  router.HandleFunc("/api/task/{id}", middleware.TaskComplete).Methods("PUT", "OPTIONS")
	  router.HandleFunc("/api/undoTask/{id}", middleware.undoTask).Methods("PUT", "OPTIONS")
	  router.HandleFunc("/api/deleteTask/{id}", middleware.deleteTask).Methods("DELETE", "OPTIONS")
	    
	  router.HandleFunc("/api/deleteAllTasks", middleware.deleteAllTasks).Methods("DELETE", "OPTIONS")
	    
	 	 
	return router
}

