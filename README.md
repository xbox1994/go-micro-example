基于go micro框架搭建的简单微服务例子，包含微服务组件网关、配置中心、熔断器以及具体服务。

注意：[必须配合参考这篇文章](https://www.wangtianyi.top/blog/2019/03/26/ji-yu-go-microde-wei-fu-wu-jia-gou-ben-di-shi-zhan/?utm_source=github&utm_medium=github)

## 依赖
* docker
* golang
* protoc、protoc-gen-go、protoc-gen-micro
* make

## 运行方式

仅Windows平台，其他平台请参考Makefile运行

1. `make build`
2. `make run`

## 整体架构
以 greeter/hello 为例，请求流程图如下：

![](https://github.com/xbox1994/go-micro-example/raw/master/index.png)

1. 访问greeter/hello
2. Micro Api解析请求，验证身份是否有效，决定请求是否继续传递，如果有效，则将解析出来的用户信息写到header中
3. GreeterService接收到gRPC请求并路由到Greeter.Hello方法中，调用其中的逻辑，尝试发送请求到UserService得到用户的其他信息，header也转发过去
4. UserService根据header中得到的id查询数据库得到具体的用户信息并返回

在服务间请求调用过程中，会有Hystrix来提供服务容错机制。在所有服务启动之前，会请求Config Service来获得对应服务的对应环境的配置信息。

## 请求方式
#### 登录
请求：
```bash
curl -X POST \
  http://localhost:8080/user/login \
  -H 'Content-Type: application/json' \
  -H 'cache-control: no-cache' \
  -d '{
	"id": "x1",
	"username": "wty",
	"password": "xxx"
}'
```

回应：
```json
{
    "code": 0,
    "message": "ok",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7ImlkIjoieDEiLCJ1c2VybmFtZSI6Ind0eSIsInBhc3N3b3JkIjoieHh4In0sImV4cCI6MTU1Mzc2MjA5OCwiaXNzIjoiZ28ubWljcm8uYXBpLnVzZXIifQ.mpUfLPGjHR7GCeDHgrUICbWuiK8fE_xZ5IfYRHyYBoE"
}
```

#### Hello
请求：
```bash
curl -X POST \
  http://localhost:8080/greeter/hello \
  -H 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7ImlkIjoieDEiLCJ1c2VybmFtZSI6Ind0eSIsInBhc3N3b3JkIjoieHh4In0sImV4cCI6MTU1Mzc2MjA5OCwiaXNzIjoiZ28ubWljcm8uYXBpLnVzZXIifQ.mpUfLPGjHR7GCeDHgrUICbWuiK8fE_xZ5IfYRHyYBoE' \
  -H 'Content-Type: application/json' \
  -H 'cache-control: no-cache' \
  -d '{
	"name": "wty"
}'
```

回应：
```json
{
    "code": 0,
    "id": "x1",
    "message": "ok",
    "password": "password from db",
    "setting_message": "Hi, here are greetings from test",
    "username": "wty"
}
```