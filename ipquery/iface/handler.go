package iface

import (
	"encoding/json"
	"net/http"

	"ipquery/application"
	"ipquery/domain"
)

type IPHandler struct {
	app *application.IPService
}

func NewIPHandler(app *application.IPService) *IPHandler {
	return &IPHandler{app: app}
}

// QueryIPHandler HTTP接口
func (h *IPHandler) QueryIPHandler(w http.ResponseWriter, r *http.Request) {
	ip := r.URL.Query().Get("ip")
	if ip == "" {
		http.Error(w, "ip required", 400)
		return
	}

	result := h.app.BatchQueryIP(r.Context(), domain.IPAddress(ip))

	// 格式化输出
	output := make(map[string]interface{})
	for ch, res := range result {
		if res.Err != nil {
			output[ch] = map[string]string{"error": res.Err.Error()}
		} else {
			output[ch] = res.Location
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}
