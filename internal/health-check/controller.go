package health_check

import (
	"GoPractice/GoCRUDs/pkg/constants"
	"GoPractice/GoCRUDs/pkg/utils"
	"net/http"
)

func (h *Health) CheckHealth(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, constants.SuccessResponseMessage, constants.SuccessResponseCode, nil, "Alive")
}
