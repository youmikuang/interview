# 面试题代码

## 1. IP 地址查询
目录：ipquery 中，具体实现方式可参考 `ipquery/README.md`

线上部署地址：http://82.158.225.153:9527/query?ip=115.191.200.34

## 2. 产品列表
目录：product-matching 中，具体实现方式可参考 `product-matching/README.md`
线上部署地址：http://82.158.225.153:9528

```bash
# 测试数据
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
