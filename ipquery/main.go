package main

import (
	"net/http"
	"time"

	"ipquery/application"
	"ipquery/iface"
	"ipquery/infra"
)

func main() {
	// 初始化依赖
	repo := infra.NewHTTPRepo()
	app := application.NewIPService(repo, 5*time.Second)
	handler := iface.NewIPHandler(app)

	// 注册路由
	http.HandleFunc("/query", handler.QueryIPHandler)

	println("server start at :8080")
	http.ListenAndServe(":8080", nil)
}
