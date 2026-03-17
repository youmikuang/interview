package main

import (
	"ip/application"
	"ip/iface"
	"ip/infra"
	"net/http"
	"strings"
	"time"
)

func cors(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if strings.HasPrefix(origin, "http://localhost") || strings.HasPrefix(origin, "https://localhost") {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		}
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next(w, r)
	}
}

func main() {
	repo := infra.NewHttpRepo()
	app := application.NewIpService(repo, 5*time.Second)
	handler := iface.NewIPHandler(app)

	http.HandleFunc("/query", cors(handler.QueryIPHandler)) // 处理的方法

	println("server start at :9525")
	http.ListenAndServe(":9525", nil)

}
