package controllers

import (
	"encoding/json"
	"jukeboxCRUDapi/database"
	"jukeboxCRUDapi/models"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Song models.Song
	json.NewDecoder(r.Body).Decode(&Song)
	database.Instance.Create(&Song)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Song)
}

func checkIfSongExists(SongId string) bool {
	var Song models.Song
	database.Instance.First(&Song, SongId)
	return Song.ID != 0
}

func GetSongById(w http.ResponseWriter, r *http.Request) {
	SongId := mux.Vars(r)["id"]
	if checkIfSongExists(SongId) == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Song Not Found!")
		return
	}
	var Song models.Song
	database.Instance.First(&Song, SongId)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Song)
}

func GetSongs(w http.ResponseWriter, r *http.Request) {
	var Songs []models.Song
	database.Instance.Find(&Songs)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Songs)
}

func UpdateSong(w http.ResponseWriter, r *http.Request) {
	SongId := mux.Vars(r)["id"]
	if checkIfSongExists(SongId) == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Song Not Found!")
		return
	}
	var Song models.Song
	database.Instance.First(&Song, SongId)
	json.NewDecoder(r.Body).Decode(&Song)
	database.Instance.Save(&Song)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Song)
}

func DeleteSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	SongId := mux.Vars(r)["id"]
	if checkIfSongExists(SongId) == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Song Not Found!")
		return
	}
	var Song models.Song
	database.Instance.Delete(&Song, SongId)
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode("Song Deleted Successfully!")
}
