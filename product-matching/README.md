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
├── domain/             // 领域层：核心业务模型+规则（仅依赖顶级接口层）
│   ├── model/          // 领域模型
│   │   ├── channel.go
│   │   ├── product.go
│   │   └── user.go
│   └── service/        // 领域服务（依赖顶级接口层）
│       └── product_filter.go
├── infra/     // 基础设施层：实现顶级接口层的端口
│   ├── repository/     // 仓储适配器：实现 api/repository 接口
│   │   ├── mock_channel_repo.go
│   │   └── mock_product_repo.go
│   └── service/        // 外部服务适配器：实现 api/service 接口
│       └── remote_api.go
└── main.go             // 测试入口
└── README.md           // 项目文档

```

## 架构核心逻辑
- api 层（端口层）：定义所有 “外部交互的契约”，不包含任何实现，仅声明接口（如 “产品仓储需要支持查所有产品”“远程校验需要支持 MD5 检查”）；
- domain 层：核心业务逻辑，依赖 api 层的接口（而非具体实现），专注 “做什么”；
- infrastructure 层：适配器，实现 api 层的接口，专注 “怎么做”（如内存 Mock、MySQL、HTTP API）；
- application 层：编排流程，依赖 domain 层和 api 层。