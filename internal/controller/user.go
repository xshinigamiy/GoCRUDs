package controller

import (
	"GoCRUDs/internal/handler"
	"GoCRUDs/pkg/constants"
	"GoCRUDs/pkg/request"
	"GoCRUDs/pkg/response"
	"GoCRUDs/pkg/utils"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var userHandler *handler.UserHandler

func init() {
	userHandler = handler.GetUserHandler()
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user request.User
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.RespondWithJSON(w, constants.FailureResponseMessage, constants.FailureResponseCode, err, nil)
		return
	}

	if jErr := json.Unmarshal(reqBody, &user); jErr != nil {
		utils.RespondWithJSON(w, constants.FailureResponseMessage, constants.FailureResponseCode, jErr, nil)
		return
	}

	err, Id := userHandler.CreateUser(user)
	if err != nil {
		utils.RespondWithJSON(w, constants.SuccessResponseMessage, constants.FailureResponseCode, err, nil)
		return
	}

	resp := response.UserResponse{
		Id: Id,
	}
	utils.RespondWithJSON(w, constants.SuccessResponseMessage, constants.SuccessResponseCode, nil, resp)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user request.User
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.RespondWithJSON(w, constants.FailureResponseMessage, constants.FailureResponseCode, err, nil)
		return
	}

	if jErr := json.Unmarshal(reqBody, &user); jErr != nil {
		utils.RespondWithJSON(w, constants.FailureResponseMessage, constants.FailureResponseCode, jErr, nil)
		return
	}

	err, Id := userHandler.CreateUser(user)
	if err != nil {
		utils.RespondWithJSON(w, constants.SuccessResponseMessage, constants.FailureResponseCode, err, nil)
		return
	}

	resp := response.UserResponse{
		Id: Id,
	}
	utils.RespondWithJSON(w, constants.SuccessResponseMessage, constants.SuccessResponseCode, nil, resp)
}
