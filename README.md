# go-micro-frame

go-micro-frame，极简、快速、零成本，积木式go微服务框架。

go-micro-frame，是一套开源组件组合而成的微服务框架。所有组件都可自行替换。

没有框架的强约束，没有学习上的成本。只需要搭建积木的方式组合自己的框架来快速开展业务。框架只保留了微服务的核心功能，使用者可以完全自主的进行改造和模块替换。目前最新版本是1.3.3， ci/cd集成, nacos做注册中心。

## 项目特点

* 极简
* 快速
* 零成本
* 积木式
* 横向扩展
* ci/cd

## 文档

文档参考：https://github.com/jettjia/go-micro-fram-doc



## 运行项目

```
0.https://github.com/jettjia/go-micro-fram-demos，参考这个项目使用
1.运行 grpc 服务，里面有代码说明和使用说明
2.运行 web-gin 服务，这里实现了调用 grpc服务，里面有代码说明和使用说明
3.文档参考 https://github.com/jettjia/go-micro-fram-doc
```

## 组件模块介绍

```
gorm		【orm】
gin		【web服务】
grpc、proto	【rpc微服务】
zap 		【日志】
viper		【配置读取】
consul 		【服务注册和发现,nacos也已经支持】
nacos		【配置中心】
grpc-lb 	【负载均衡】
es		【搜索】
分布式锁	        【redis实现】
幂等性		【grpc重试，分布式锁处理，token处理等】
jaeger		【链路追踪】
sentinel	【限流、熔断、降级】
kong		【网关】
amqp            【amqp，消息队列，比如：rabbitmq】
cron            【分布式定时任务;go:go-cron,java:xxl-job】
分布式事务	【方式1：rocketmq，事务消息方式；方式2：seata-golang,暂缺】
分布式mysql	【go: gaea分库分表; java: shardingsphere-proxy】
```

## 版本说明

```
v1.1.X 完成
已经实现了模块介绍中的功能
```

```
v1.2.X 完成
会把microframe.com里的模块全部实现，进行统一维护，方便更新和迭代替换。
nacos：已封装
logger: 已封装
jaeger: 已封装
mysql：已封装，连接池方式
consul-register: 已封装
consul-discovery: 已封装
snowflake: 已封装
rabbitmq: 已封装
elasticSearch: 已封装
sentinel
redis: 已封装:alone sentinel cluster 三种模式
```

```
v1.3.X
会增加k8s的自动化发布脚本，Prometheus、Grafana监控等
v1.3.3
增加了cicd自动化发布脚本,docker打包脚本等
这里使用的是：drone做的cicd，类似jenkins，但是比jenkins轻便很多。
代码仓库是： gogs; docker仓库的是:harbor

v1.3.4 开发中
cicd，自动化发布到k8s中
```

```
v1.4.X 规划
web层多实现
web-springboot: java版本的web, 利用springboot框架调用grpc服务
web-hyperf: php版本的web，利用 hyperf框架调用grpc服务
```

```
v2.0 规划
会改造到istio或者dapr的三代微服务方式中，会单独另起一个项目进行维护
```



## ci/cd

ci/cd-docker

```
drone	【cicd】
harbor	【docker仓库】
gogs	【code仓库】

这里用drone替代了 jenkins；drone采取pipeline的方式进行自动化构建和发布。只需要维护.drone.yml文件。按照yml的格式编写我们的项目就可以了。
代码提交的时候，会自动触发，自动构建，自动发布。大大的节约开发和运维的时间
```

 drone![image-20210928144848124](images/image-20210928144848124.png)

harbor ![image-20210928144916688](images/image-20210928144916688.png)

gogs![image-20210928144942719](images/image-20210928144942719.png)

## 项目演示

### 启动 grpc 服务

1. nacos后台配置

    配置内容参考, grpc里的配置示例

    <img src="images/image-20210921163646319.png" alt="image-20210921163646319" style="zoom:50%;" />

2. 启动grpc项目，会从nacos读取配置

```
[root@localhost grpc]# go run main.go
2021-09-21T15:38:26.457+0800    INFO    nacos_client/nacos_client.go:87 logDir:<tmp/nacos/log>   cacheDir:<tmp/nacos/cache>
2021-09-21T15:38:27.955+0800    INFO    nacos/nacos.go:26       从nacos读取到的全部配置如下：%!(EXTRA string={

```

​	consul会注册服务，这里可以启动多个grpc服务，已经实现了负载均衡的获取服务。

  <img src="images/image-20210921162747411.png" alt="image-20210921162747411" style="zoom:33%;" />



### 启动 web 服务

1. 启动web项目，会从nacos读取配置

   ```
   [root@localhost web-gin]# go run main.go
   2021-09-21T16:37:38.927+0800    INFO    nacos_client/nacos_client.go:87 logDir:<tmp/nacos/log>   cacheDir:<tmp/nacos/cache>
   2021-09-21T16:37:39.118+0800    INFO    nacos/nacos.go:26       从nacos读取到的全部配置如下：%!(EXTRA string={
   
   ```

2. 这里会从consul中负载均衡的获取到服务，也会把web注入到consul中。这样可以用nginx或者kong等来读web进行负载均衡

   

3. 请求用postman 访问web层的接口

      <img src="images/image-20210921164033597.png" alt="image-20210921164033597" style="zoom: 25%;" />

4. 链路追踪

​       <img src="images/image-20210921164122346.png" alt="image-20210921164122346" style="zoom: 25%;" />



### 熔断限流降级

参考 sentinel章节，在web层增加处理



### 网关

参考kong 章节，在konga后台进行配置
