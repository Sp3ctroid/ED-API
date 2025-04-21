package main

import (
	"net/http"

	db "github.com/Sp3ctroid/ED-API/database"
	"github.com/Sp3ctroid/ED-API/handlers"
	"github.com/Sp3ctroid/ED-API/storage"

	"github.com/gorilla/mux"
)

func main() {

	mux := mux.NewRouter()

	home := handlers.HomeHandler{}

	dbstore := storage.NewDBStore(db.OpenStorage())

	albumH := handlers.NewAlbumHandler(dbstore)
	usersH := handlers.NewUsersHandler(dbstore)

	mux.HandleFunc("/", home.ServeHTTP)
	mux.HandleFunc("/albums/{id}", albumH.FindById).Methods("GET")
	mux.HandleFunc("/albums", albumH.CreateAlbum).Methods("POST")
	mux.HandleFunc("/albums", albumH.GetAlbums).Methods("GET")
	mux.HandleFunc("/albums/{id}", albumH.PutAlbum).Methods("PUT")

	users := mux.PathPrefix("/users").Subrouter()

	users.HandleFunc("/login", usersH.Login).Methods("POST")
	users.HandleFunc("/register", usersH.Register).Methods("POST")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
