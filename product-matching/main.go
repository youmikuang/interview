package main

import (
	"log"
	"net/http"

	"product-matching/application"
	domain_service "product-matching/domain/service"
	infra_handler "product-matching/infra/handler"
	infra_repository "product-matching/infra/repository"
	infra_service "product-matching/infra/service"
)

func main() {
	// 1. 初始化适配器
	remoteChecker := infra_service.NewRemoteAPI("https://mock-api.com/check-md5")
	productRepo := &infra_repository.MockProductRepo{}
	channelRepo := &infra_repository.MockChannelRepo{}

	// 2. 注入适配器到领域服务
	filterService := domain_service.NewProductFilter(remoteChecker, productRepo, channelRepo)
	matchApp := application.NewProductMatchingApp(filterService)

	// 3. 注册路由
	http.HandleFunc("/match", infra_handler.NewMatchHandler(matchApp))

	// 4. 启动服务
	log.Println("server starting on :9528")
	if err := http.ListenAndServe(":9528", nil); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
