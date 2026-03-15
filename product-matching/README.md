# 产品列表

> 根据用户填写的信息进行产品的筛选。
1. 渠道可以通过配置进行过滤，配置在 json 中
2. 特殊过滤器，通过 md5 进行过滤


## 目录结构

```plaintext
product-matching/
├── api/                // 顶级接口层（端口层）：定义所有外部交互的契约
│   ├── repository/     // 仓储端口：定义数据访问能力
│   │   ├── channel_repo.go
│   │   └── product_repo.go
│   └── service/        // 外部服务端口：定义远程调用能力
│       └── remote_checker.go
├── application/        // 应用层：流程编排
│   └── product_matcher.go
├── config/        // 配置层
│   └── config.go
│   └── config.json
├── domain/             // 领域层：核心业务模型+规则（仅依赖顶级接口层）
│   ├── model/          // 领域模型
│   │   ├── channel.go
│   │   ├── product.go
│   │   └── user.go
│   └── service/        // 领域服务（依赖顶级接口层）
│       └── product_filter.go
├── infra/     // 基础设施层：实现顶级接口层的端口
│   ├── handler/        // HTTP 处理器：入站适配器
│   │   └── match_handler.go
│   │   └── match_handler_test.go
│   ├── repository/     // 仓储适配器：实现 api/repository 接口
│   │   ├── mock_channel_repo.go
│   │   └── mock_product_repo.go
│   └── service/        // 外部服务适配器：实现 api/service 接口
│       └── remote_api.go
└── main.go             // 入口：依赖组装 + 路由注册 + 启动服务
└── README.md           // 项目文档

```

## 架构核心逻辑
- api 层（端口层）：定义所有 “外部交互的契约”，不包含任何实现，仅声明接口（如 “产品仓储需要支持查所有产品”“远程校验需要支持 MD5 检查”）；
- domain 层：核心业务逻辑，依赖 api 层的接口（而非具体实现），专注 “做什么”；
- infrastructure 层：适配器，实现 api 层的接口，专注 “怎么做”（如内存 Mock、MySQL、HTTP API）；
- application 层：编排流程，依赖 domain 层和 api 层。

## 启动方式

```bash
go run .
```

服务启动后监听 `http://localhost:9528`。

## 运行测试

```bash
go test -v ./...
```

## API 调用示例

**匹配成功用户：**

```bash
curl -X POST 'http://82.158.225.153:9528/match?channel_id=C001' \
  -H 'Content-Type: application/json' \
  -d '{
    "phone": "123456",
    "name": "张三",
    "age": 30,
    "gender": "男",
    "region": "北京",
    "hasHouse": true,
    "hasCar": true,
    "hasSocial": true
  }'
```

**不匹配用户：**

```bash
curl -X POST 'http://82.158.225.153:9528/match?channel_id=C001' \
  -H 'Content-Type: application/json' \
  -d '{
    "phone": "654321",
    "name": "李四",
    "age": 18,
    "gender": "女",
    "region": "上海",
    "hasHouse": false,
    "hasCar": false,
    "hasSocial": true
  }'
```