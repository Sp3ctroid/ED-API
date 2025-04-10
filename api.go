package main

import (
	"fmt"
	"net/http"

	"github.com/Sp3ctroid/ED-API/handlers"
	"github.com/Sp3ctroid/ED-API/storage"
	"github.com/Sp3ctroid/ED-API/types"

	"github.com/gorilla/mux"
)

func main() {

	mux := mux.NewRouter()

	home := handlers.HomeHandler{}
	slicestore := storage.NewSliceStore()
	slicestore.Items = append(slicestore.Items, types.Album{0, "Doobi Do", "Kim 5+", 14.30})
	slicestore.Items = append(slicestore.Items, types.Album{1, "Cringe", "Poser", 42.10})
	albumH := handlers.NewAlbumHandler(slicestore)

	mux.HandleFunc("/", home.ServeHTTP)
	mux.HandleFunc("/albums/{id}", albumH.FindById).Methods("GET")
	mux.HandleFunc("/albums", albumH.CreateAlbum).Methods("POST")
	mux.HandleFunc("/albums", albumH.GetAlbums).Methods("GET")
	mux.HandleFunc("/albums/{id}", albumH.PutAlbum).Methods("PUT")

	fmt.Printf("%v", slicestore.GetAllAlbums())
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
