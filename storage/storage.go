package storage

import (
	"github.com/Sp3ctroid/ED-API/types"
)

type SliceStore struct {
	Items []types.Album
}

type AlbumStore interface {
	GetAlbumByID(ID int) (album types.Album)
	PostAlbum(album types.Album)
	GetAllAlbums() []types.Album
	ChangeAlbum(id int, album types.Album)
}

func (slicestore SliceStore) GetAlbumByID(id int) types.Album {
	return slicestore.Items[id]
}

func (slicestore *SliceStore) PostAlbum(album types.Album) {
	slicestore.Items = append(slicestore.Items, album)
}

func (slicestore SliceStore) GetAllAlbums() []types.Album {
	return slicestore.Items
}

func (slicestore *SliceStore) ChangeAlbum(id int, album types.Album) {
	slicestore.Items[id] = album
}

func NewSliceStore() *SliceStore {
	return &SliceStore{}
}
