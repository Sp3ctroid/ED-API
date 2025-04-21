package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	jwt_auth "github.com/Sp3ctroid/ED-API/jwt"
	"github.com/Sp3ctroid/ED-API/storage"
	"github.com/Sp3ctroid/ED-API/types"
	"github.com/golang-jwt/jwt/v5"

	"github.com/gorilla/mux"
)

type HomeHandler struct {
}

type AlbumHandler struct {
	store storage.AlbumStore
}

type UserHandler struct {
	store storage.UserStore
}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello World"))
}

func NewUsersHandler(storage storage.UserStore) *UserHandler {
	return &UserHandler{store: storage}
}

func NewAlbumHandler(storage storage.AlbumStore) *AlbumHandler {
	return &AlbumHandler{store: storage}
}

func (h *AlbumHandler) FindById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	w.Header().Set("Content-Type", "application/json")
	id_int, _ := strconv.Atoi(id)
	got_album, err := json.Marshal(h.store.GetAlbumByID(id_int))
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Write(got_album)
	}

}

func (h *AlbumHandler) CreateAlbum(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	JSON_STATUS := types.JSON_Status{}
	req_token := r.Header.Get("login")
	token, err := jwt_auth.ValidateToken("secret", req_token)
	claims := token.Claims.(jwt.MapClaims)
	subject := claims["sub"].(bool)
	if err != nil || !subject {
		JSON_STATUS.Response("Failed", "Invalid Token or not enough rights", req_token)
		JSON_STATUS_JSON, _ := json.Marshal(JSON_STATUS)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(JSON_STATUS_JSON)
		return
	}

	album := types.Album{}
	err = json.NewDecoder(r.Body).Decode(&album)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	h.store.PostAlbum(album)

	w.WriteHeader(http.StatusCreated)
	JSON_STATUS.Response("Success", "Album Created", album)
	JSON_STATUS_JSON, _ := json.Marshal(JSON_STATUS)
	w.Write(JSON_STATUS_JSON)
}

func (h *AlbumHandler) GetAlbums(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	all_albums := h.store.GetAllAlbums()
	JSON_Albums, _ := json.Marshal(all_albums)
	w.Write(JSON_Albums)
}

func (h *AlbumHandler) PutAlbum(w http.ResponseWriter, r *http.Request) {
	album := types.Album{}
	err := json.NewDecoder(r.Body).Decode(&album)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	JSON_STATUS := types.JSON_Status{}
	req_token := r.Header.Get("login")

	token, err := jwt_auth.ValidateToken("secret", req_token)
	claims := token.Claims.(jwt.MapClaims)
	subject := claims["sub"].(bool)

	if err != nil || !subject {
		JSON_STATUS.Response("Failed", "Invalid Token or not enough rights", req_token)
		JSON_STATUS_JSON, _ := json.Marshal(JSON_STATUS)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(JSON_STATUS_JSON)
		return
	}

	id := mux.Vars(r)["id"]
	id_int, _ := strconv.Atoi(id)
	h.store.ChangeAlbum(id_int, album)
	album.ID = id_int
	w.WriteHeader(http.StatusCreated)
	JSON_STATUS.Response("Success", "Album Updated", album)
	JSON_STATUS_JSON, _ := json.Marshal(JSON_STATUS)
	w.Write(JSON_STATUS_JSON)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {

	user := types.User{}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	got_user, err := h.store.LoginUser(user.Username, user.Password)
	status := types.JSON_Status{}
	if err != nil {
		status.Response("Failed", "Password or Username is incorrect, or User doesn't exist. Login Failed", user)
		status_JSON, _ := json.Marshal(status)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(status_JSON)
		return
	}

	token, err := jwt_auth.CreateToken("secret", got_user)
	if err != nil {
		status.Response("Failed", "Token Creation Failed", user)
		status_JSON, _ := json.Marshal(status)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(status_JSON)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	status.Response("Success", "Login Successful", map[string]string{"token": token})
	status_JSON, _ := json.Marshal(status)
	w.Write(status_JSON)
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	user := types.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	status := types.JSON_Status{}
	if err != nil {
		status.Response("Failed", "Bad Request", user)
		status_JSON, _ := json.Marshal(status)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(status_JSON)
		return
	}

	if h.store.CheckUser(user.Username) {
		status.Response("Failed", "User Already Exists", user)
		status_JSON, _ := json.Marshal(status)
		w.WriteHeader(http.StatusConflict)
		w.Write(status_JSON)
		return
	}

	err = h.store.RegisterUser(user.Username, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		status.Response("Failed", "Something Went Wrong", user)
		status_JSON, _ := json.Marshal(status)
		w.Write(status_JSON)
		return
	}

	status.Response("Success", "User Created", user)
	status_JSON, _ := json.Marshal(status)
	w.WriteHeader(http.StatusOK)
	w.Write(status_JSON)

}
