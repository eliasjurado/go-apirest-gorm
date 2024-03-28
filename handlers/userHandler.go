package handlers

import (
	"apirest-gorm/models"
	"apirest-gorm/repository"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {

	res := models.Response{}
	format := r.Header.Get("x-format")

	users, err := repository.GetAllUsers()
	if err != nil {
		res.Status = http.StatusInternalServerError
		Send(w, res, format)
		return
	}

	res.Data = users
	res.Status = http.StatusOK
	Send(w, res, format)
}

func GetOneUser(w http.ResponseWriter, r *http.Request) {
	res := models.Response{}
	vars := mux.Vars(r)
	format := r.Header.Get("x-format")
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		res.Status = http.StatusUnprocessableEntity
		Send(w, res, format)
		return
	}

	user, err := repository.GetOneUser(int64(id))
	if err != nil || user.Id != int64(id) {
		res.Status = http.StatusNotFound
		Send(w, res, format)
		return
	}

	res.Data = user
	res.Status = http.StatusOK
	Send(w, res, format)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	res := models.Response{}
	format := r.Header.Get("x-format")
	user := &models.User{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		res.Status = http.StatusUnprocessableEntity
		Send(w, res, format)
		return
	}
	user.Id = 0
	user, err = repository.SaveUser(user)
	if err != nil {
		res.Status = http.StatusInternalServerError
		Send(w, res, format)
		return
	}

	res.Data = user
	res.Status = http.StatusCreated
	Send(w, res, format)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	res := models.Response{}
	vars := mux.Vars(r)
	format := r.Header.Get("x-format")
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		res.Status = http.StatusUnprocessableEntity
		Send(w, res, format)
		return
	}

	dbUser, err := repository.GetOneUser(int64(id))
	if err != nil || dbUser.Id != int64(id) {
		res.Status = http.StatusNotFound
		Send(w, res, format)
		return
	}

	user := &models.User{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&user)
	if err != nil {
		res.Status = http.StatusUnprocessableEntity
		Send(w, res, format)
		return
	}

	dbUser.Username = user.Username
	dbUser.Password = user.Password
	dbUser.Email = user.Email

	user, err = repository.SaveUser(dbUser)
	if err != nil {
		res.Status = http.StatusInternalServerError
		Send(w, res, format)
		return
	}

	res.Data = user
	res.Status = http.StatusCreated
	Send(w, res, format)
}

func RemoveUser(w http.ResponseWriter, r *http.Request) {
	res := models.Response{}
	vars := mux.Vars(r)
	format := r.Header.Get("x-format")
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		res.Status = http.StatusUnprocessableEntity
		Send(w, res, format)
		return
	}

	user, err := repository.GetOneUser(int64(id))
	if err != nil || user.Id != int64(id) {
		res.Status = http.StatusNotFound
		Send(w, res, format)
		return
	}

	err = repository.RemoveUser(user)
	if err != nil {
		res.Status = http.StatusInternalServerError
		Send(w, res, format)
		return
	}

	res.Data = user
	res.Status = http.StatusOK
	Send(w, res, format)
}
