package main

import (
	"fmt"
	"jukeboxCRUDapi/controllers"
	"jukeboxCRUDapi/database"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterSongRoutes(router *mux.Router) {
	router.HandleFunc("/api/songs", controllers.GetSongs).Methods("GET")
	router.HandleFunc("/api/songs/{id}", controllers.GetSongById).Methods("GET")
	router.HandleFunc("/api/songs/create", controllers.CreateSong).Methods("POST")
	router.HandleFunc("/api/songs/{id}", controllers.UpdateSong).Methods("PUT")
	router.HandleFunc("/api/songs/{id}", controllers.DeleteSong).Methods("DELETE")
}

func main() {
	// Load Configurations from config.json using Viper
	LoadAppConfig()

	// Initialize Database
	database.Connect(AppConfig.ConnectionString)
	database.Migrate()

	// Initialize the router
	router := mux.NewRouter().StrictSlash(true)

	// Register Routes
	RegisterSongRoutes(router)

	// Start the server
	log.Println(fmt.Sprintf("Starting Server on port %s", AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port), router))
}
