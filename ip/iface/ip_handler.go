package iface

import (
	"encoding/json"
	"ip/application"
	"ip/domain"
	"net/http"
)

type IPHandler struct {
	app *application.IPService
}

func NewIPHandler(app *application.IPService) *IPHandler {
	return &IPHandler{app: app}
}

func (h *IPHandler) QueryIPHandler(w http.ResponseWriter, r *http.Request) {
	ip := r.URL.Query().Get("ip")
	channels := r.URL.Query()["channels"]
	if ip == "" || len(channels) == 0 {
		http.Error(w, "ip or channels is empty", http.StatusBadRequest)
		return
	}
	result := h.app.BatchQueryIP(r.Context(), domain.IPAddress(ip), channels)
	output := make(map[string]interface{})

	for ch, res := range result {
		if res.Err != nil {
			output[ch] = map[string]any{"error": res.Err.Error(), "cost": res.Cost}
		} else {
			output[ch] = map[string]any{
				"province": res.Location.Province,
				"city":     res.Location.City,
				"cost":     res.Cost,
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}
