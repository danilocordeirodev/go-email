package endpoints

import (
	"net/http"
)

func (h *Handler) CampaignGet(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {

	//campaigns, err := h.CampaignService.Repository.FindAll()
	//return campaigns, 200, err
	return nil, 200, nil
}
