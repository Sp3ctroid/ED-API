package storage

import (
	"database/sql"

	"github.com/Sp3ctroid/ED-API/types"
)

type SliceStore struct {
	Items []types.Album
}

type DBStore struct {
	storage *sql.DB
}

type UserStore interface {
	LoginUser(username string, password string) (user types.User, err error)
}
type AlbumStore interface {
	GetAlbumByID(ID int) (album types.Album)
	PostAlbum(album types.Album)
	GetAllAlbums() []types.Album
	ChangeAlbum(id int, album types.Album)
}

func NewDBStore(db *sql.DB) *DBStore {
	return &DBStore{storage: db}
}

func (db DBStore) LoginUser(username string, password string) (types.User, error) {
	user := types.User{}
	row := db.storage.QueryRow("SELECT * FROM users WHERE username = ? AND passw = ?", username, password)

	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Is_admin)

	if err != nil {
		return types.User{}, err
	}

	return user, nil

}

func (db DBStore) GetAlbumByID(ID int) types.Album {
	rows, err := db.storage.Query("SELECT * FROM ALBUMS WHERE id = ?", ID)
	if err != nil {
		panic(err)
	}

	album := types.Album{}

	for rows.Next() {
		err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price)
		if err != nil {
			panic(err)
		}
	}
	return album
}

func (db *DBStore) PostAlbum(album types.Album) {
	_, err := db.storage.Exec("INSERT INTO ALBUMS (title, artist, price) VALUES (?, ?, ?)", album.Title, album.Artist, album.Price)
	if err != nil {
		panic(err)
	}
}

func (db DBStore) GetAllAlbums() []types.Album {
	rows, err := db.storage.Query("SELECT * FROM ALBUMS")
	if err != nil {
		panic(err)
	}

	albums := []types.Album{}

	for rows.Next() {
		album := types.Album{}
		err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price)
		if err != nil {
			panic(err)
		}
		albums = append(albums, album)
	}
	return albums
}

func (db *DBStore) ChangeAlbum(id int, album types.Album) {
	_, err := db.storage.Exec("UPDATE ALBUMS SET title = ?, artist = ?, price = ? WHERE id = ?", album.Title, album.Artist, album.Price, id)
	if err != nil {
		panic(err)
	}
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
