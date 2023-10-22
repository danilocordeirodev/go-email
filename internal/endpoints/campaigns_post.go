package endpoints

import (
	"errors"
	"net/http"

	"github.com/danilocordeirodev/go-email/internal/contract"
	internalerrors "github.com/danilocordeirodev/go-email/internal/internalErrors"
	"github.com/go-chi/render"
)

func(h *Handler) CampaignPost(w http.ResponseWriter, r *http.Request) {
	var request contract.NewCampaignReq
	render.DecodeJSON(r.Body, &request)
	id, err := h.CampaignService.Create(request)

	if err != nil {
		if errors.Is(err, internalerrors.ErrInternal) {
			render.Status(r, 500)
		} else {
			render.Status(r, 400)
		}
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}

	render.Status(r, 201)
	render.JSON(w, r, map[string]string{"id": id})
}