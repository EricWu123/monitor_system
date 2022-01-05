# monitor_system

- config 配置
- global 全局变量
- scripts 相关脚本
- internal 内部相关模块，包括model， router等
- docs 作业需求等相关文档

## 如何运行
### depend
go version go1.17.1
mysql  Ver 14.14 Distrib 5.7.21

### run
1、运行前先创建好mysql数据库，并将docs/mysql中的表导入数据库中。

2、运行服务端

```
cd server
go run main.go
```

3、运行客户端

```
cd client
go run main.go system_info.go
```

4、模拟用户登录并查询系统信息

```
cd server/scripts
./run.sh
```
### 单元测试
```
cd server/scripts
./test.sh
```
生成的单元测试报告在：server/conver.out  server/converage.html
