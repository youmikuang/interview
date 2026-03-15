package main

import (
	"fmt"
	"product-matching/application"
	"product-matching/domain/model"
	domain_service "product-matching/domain/service"
	infra_repository "product-matching/infra/repository"
	infra_service "product-matching/infra/service"
)

func main() {
	// 1. 初始化适配器（实现顶级接口）
	remoteChecker := infra_service.NewRemoteAPI("https://mock-api.com/check-md5")
	productRepo := &infra_repository.MockProductRepo{}
	channelRepo := &infra_repository.MockChannelRepo{}

	// 2. 注入适配器到领域服务
	filterService := domain_service.NewProductFilter(remoteChecker, productRepo, channelRepo)
	matchApp := application.NewProductMatchingApp(filterService)

	// 3. 测试用户
	user1 := &model.User{
		Phone:     "123456",
		Name:      "张三",
		Age:       30,
		Gender:    "男",
		Region:    "北京",
		HasHouse:  true,
		HasCar:    true,
		HasSocial: true,
	}

	user2 := &model.User{
		Phone:     "654321",
		Name:      "李四",
		Age:       18,
		Gender:    "女",
		Region:    "上海",
		HasHouse:  false,
		HasCar:    false,
		HasSocial: true,
	}

	// 4. 执行匹配
	fmt.Println("===== 测试用户1（符合规则） =====")
	products1, err := matchApp.MatchProducts(user1, "C001")
	if err != nil {
		fmt.Printf("匹配失败：%v\n", err)
	} else {
		fmt.Printf("匹配到%d个产品：\n", len(products1))
		for _, p := range products1 {
			fmt.Printf("- %s（%s）\n", p.ID, p.Name)
		}
	}

	fmt.Println("\n===== 测试用户2（不符合规则） =====")
	products2, err := matchApp.MatchProducts(user2, "C001")
	if err != nil {
		fmt.Printf("匹配失败：%v\n", err)
	} else {
		fmt.Printf("匹配到%d个产品：\n", len(products2))
		for _, p := range products2 {
			fmt.Printf("- %s（%s）\n", p.ID, p.Name)
		}
	}
}
