package health_check

import (
	"GoCRUDs/pkg/constants"
	"GoCRUDs/pkg/utils"
	"net/http"
)

func (h *Health) CheckHealth(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, constants.SuccessResponseMessage, constants.SuccessResponseCode, nil, "Alive")
}
