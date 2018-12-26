基于go micro框架搭建的简单微服务例子，包含微服务组件网关、配置中心、熔断器以及具体服务。

## 需要安装的工具
* Docker
* Golang、项目依赖包
* protoc-gen-micro
* make

## 运行方式
1. `make`
2. `docker-compose build`
2. `docker-compose up --force-recreate`

## 整体架构
http://www.wangtianyi.top/blog/2018/09/28/ji-yu-go-microde-wei-fu-wu-jia-gou-ben-di-shi-zhan/

以 greeter/hello 为例，请求流程图如下：

![image](https://wx2.sinaimg.cn/mw1024/bea16acagy1fvrqdjvp9lj20na0hp3zg.jpg)

1. 访问greeter/hello
2. Micro Api解析请求，验证身份是否有效，决定请求是否继续传递，如果有效，则将解析出来的用户信息写到header中
3. GreeterService接收到gRPC请求并路由到Greeter.Hello方法中，调用其中的逻辑，尝试发送请求到UserService得到用户的其他信息，header也转发过去
4. UserService根据header中得到的id查询数据库得到具体的用户信息并返回

在服务间请求调用过程中，会有Hystrix来提供服务容错机制。在所有服务启动之前，会请求Config Service来获得对应服务的对应环境的配置信息。
