package user

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/fiscaluno/mu/db"
	"github.com/fiscaluno/pandorabox"

	"github.com/gorilla/mux"
)

// TokenDetails JWT Infos
func TokenDetails(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	jwt := vars["jwt"]

	info, err := TokenInfos(jwt)
	if err != nil {
		pandorabox.RespondWithJSON(w, http.StatusInternalServerError, info)
		return
	}

	pandorabox.RespondWithJSON(w, http.StatusOK, info)

}

// FindAll Users
func FindAll(w http.ResponseWriter, r *http.Request) {
	users := GetAll()
	pandorabox.RespondWithJSON(w, http.StatusOK, users)
}

// FindByFacebookID find a User by FacebookID
func FindByFacebookID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	FacebookID := vars["id"]
	users := GetByQuery("facebook_id = ?", FacebookID)

	if len(users) >= 0 {
		pandorabox.RespondWithJSON(w, http.StatusOK, users)
		return
	}

	msg := pandorabox.Message{
		Content: "Not exist this user",
		Status:  "ERROR",
		Body:    nil,
	}
	pandorabox.RespondWithJSON(w, http.StatusOK, msg)

}

// FindByID find a User by ID
func FindByID(w http.ResponseWriter, r *http.Request) {

	var msg pandorabox.Message

	msg = pandorabox.Message{
		Content: "Invalid ID, not a int",
		Status:  "ERROR",
		Body:    nil,
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		pandorabox.RespondWithJSON(w, http.StatusOK, msg)
		return
	}

	user := GetByID(id)

	if user.ID != 0 {
		pandorabox.RespondWithJSON(w, http.StatusOK, user)
		return
	}

	msg = pandorabox.Message{
		Content: "Not exist this user",
		Status:  "ERROR",
		Body:    nil,
	}
	pandorabox.RespondWithJSON(w, http.StatusOK, msg)

}

// Add a User
func Add(w http.ResponseWriter, r *http.Request) {

	var user User
	var msg pandorabox.Message

	msg = pandorabox.Message{
		Content: "Invalid request payload",
		Status:  "ERROR",
		Body:    nil,
	}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		pandorabox.RespondWithJSON(w, http.StatusBadRequest, msg)
		return
	}

	userx, err := user.AddWithVerification()
	if err != nil {
		msg = pandorabox.Message{
			Content: err.Error(),
			Status:  "ERROR",
			Body:    nil,
		}
		pandorabox.RespondWithJSON(w, http.StatusInternalServerError, msg)
		return
	}

	msg = pandorabox.Message{
		Content: "New user successfully added or already existed",
		Status:  "OK",
		Body:    userx,
	}
	pandorabox.RespondWithJSON(w, http.StatusCreated, msg)

}

// Addx a User
func Addx(w http.ResponseWriter, r *http.Request) {
	db := db.Conn()
	defer db.Close()

	var user User
	var msg pandorabox.Message

	msg = pandorabox.Message{
		Content: "Invalid request payload",
		Status:  "ERROR",
		Body:    nil,
	}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		pandorabox.RespondWithJSON(w, http.StatusBadRequest, msg)
		return
	}

	// user.Password = pandorabox.MD5(user.Password)

	db.Create(&user)

	msg = pandorabox.Message{
		Content: "New user successfully added",
		Status:  "OK",
		Body:    user,
	}
	pandorabox.RespondWithJSON(w, http.StatusCreated, msg)

}

// DeleteByID find a User by ID
func DeleteByID(w http.ResponseWriter, r *http.Request) {
	db := db.Conn()
	defer db.Close()

	var user User
	var msg pandorabox.Message

	msg = pandorabox.Message{
		Content: "Invalid ID, not a int",
		Status:  "ERROR",
		Body:    nil,
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		pandorabox.RespondWithJSON(w, http.StatusOK, msg)
		return
	}

	db.Find(&user, id)
	if user.ID != 0 {
		db.Delete(&user)
		msg = pandorabox.Message{
			Content: "Deleted user successfully",
			Status:  "OK",
			Body:    user,
		}

		pandorabox.RespondWithJSON(w, http.StatusAccepted, msg)
		return
	}

	msg = pandorabox.Message{
		Content: "Not exist this user",
		Status:  "ERROR",
		Body:    nil,
	}
	pandorabox.RespondWithJSON(w, http.StatusOK, msg)

}

// UpdateByID find a User by ID
func UpdateByID(w http.ResponseWriter, r *http.Request) {
	db := db.Conn()
	defer db.Close()

	var user User
	var msg pandorabox.Message

	msg = pandorabox.Message{
		Content: "Invalid ID, not a int",
		Status:  "ERROR",
		Body:    nil,
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		pandorabox.RespondWithJSON(w, http.StatusOK, msg)
		return
	}

	msg = pandorabox.Message{
		Content: "Invalid request payload",
		Status:  "ERROR",
		Body:    nil,
	}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		pandorabox.RespondWithJSON(w, http.StatusBadRequest, msg)
		return
	}

	if id >= 0 {
		user.ID = uint(id)
		db.Save(&user)

		msg = pandorabox.Message{
			Content: "Update user successfully ",
			Status:  "OK",
			Body:    user,
		}
		pandorabox.RespondWithJSON(w, http.StatusAccepted, msg)
		return
	}

	msg = pandorabox.Message{
		Content: "Not exist this user",
		Status:  "ERROR",
		Body:    nil,
	}
	pandorabox.RespondWithJSON(w, http.StatusOK, msg)

}
