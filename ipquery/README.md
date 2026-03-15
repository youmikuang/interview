# IP Query 

> 按照 DDD 的设计思路，需要对代码进行分层处理
> 
> 具体分为：Domain、Application、Infrastructure、Interface 四层

## DDD 项目分层构造
依赖方向：Interface → Application → Domain ← Infrastructure
- Domain：我要干什么（业务）
- Application：按什么顺序干（流程）
- Infrastructure：具体怎么干（技术）
- Interface：给外部提供入口（接入）

### 1. Domain 领域层（核心）
- 实体、值对象、领域服务
- 仓库接口（只定义，不实现）
- 纯业务，无任何技术代码

### 2. Application 应用层
- 业务流程编排
- 并发、事务、超时、权限控制
- 只调度，不写业务规则

### 3. Infrastructure 基础设施层
- 实现领域仓库接口
- DB、Redis、HTTP、第三方 API
- 只做技术实现，不掺业务

### 4. Interface 接口层
- HTTP/gRPC/ 消息入口
- 参数校验、格式封装
- 对外服务入口

## DDD vs MVC

### MVC 
> Model、View、Controller 的理念
> 
> 缺点：Model 只有数据，没有行为，业务全堆在 Controller 里

之前项目上关于 MVC 的布局，引入 Service 和 Repository。
1. Service 写具体的业务逻辑，和 Application 类似
2. Repository 做数据查询，和 Infrastructure 类似。
3. Interface 层写到了 Service 层里。


## 项目代码布局

```plaintext
ipquery/
├── domain/          # 领域：模型 + 接口
│   └── ipquery.go
├── application/     # 应用：流程编排
│   └── service.go
├── infra/           # 基础：DB/HTTP/第三方
│   └── ip_repo.go
├── iface/           # 接口：HTTP/gRPC 入口
│   └── handler.go
├── go.mod
├── main.go          # 依赖注入 & 启动
├── README.md
```

## IP 接口使用
1. [ipinfo.io](https://ipinfo.io/)
    - 示例：https://api.ipinfo.io/lite/125.37.212.152?token=ce146ea4171769
2. [ip-api.com](http://ip-api.com)，有次数限制
    - 示例：http://ip-api.com/json/125.37.212.152?lang=zh-CN
3. 两个不能获取数据的网站，分别模拟无法访问的数据和超时的设置

## 测试

```bash
# 解压后进入目录
cd ipquery

# 直接启动服务
go run main.go

# 访问接口测试
curl "http://127.0.0.1:8080/query?ip=103.62.49.170"
```

返回结果示例：
```json
{
    "A": {
        "error": "context deadline exceeded"
    },
    "B": {
        "error": "invalid status code: 403"
    },
    "C": {
        "province": "东京都",
        "city": "东京"
    },
    "D": {
        "province": "Japan",
        "city": "Japan"
    }
}
```

## 结果分析
- 4 渠道并发请求
- 5 秒全局超时
- DDD 四层架构最佳实践
- 失败隔离、不崩溃、不阻塞
