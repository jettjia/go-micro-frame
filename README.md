# go-micro-frame

go-micro-frame，极简、快速、零成本，积木式go微服务框架。

go-micro-frame，是一套开源组件组合而成的微服务框架。所有组件都可自行替换。

没有框架的强约束，没有学习上的成本。只需要搭建积木的方式组合自己的框架来快速开展业务。框架只保留了微服务的核心功能，使用者可以完全自主的进行改造和模块替换， ci/cd集成, nacos做注册中心，也可以consul替换。

## 项目特点

* 极简
* 快速
* 零成本
* 积木式
* 横向扩展
* ci/cd


## 文档

文档参考：https://github.com/jettjia/go-micro-fram-doc

框架使用参考：https://github.com/jettjia/go-micro-frame-demos

## 运行项目

```
0.https://github.com/jettjia/go-micro-fram-demos，参考这个项目使用
1.运行 grpc 服务，里面有代码说明和使用说明
2.运行 web-gin 服务，这里实现了调用 grpc服务，里面有代码说明和使用说明
3.文档参考 https://github.com/jettjia/go-micro-fram-doc
```

## 组件模块介绍
核心模块
```
gorm		【orm】
gin		【web服务】
grpc、proto	【rpc微服务】
zap 		【日志】
viper		【配置读取】
consul 		【服务注册和发现,nacos也已经支持】
nacos		【配置中心】
grpc-lb 	【负载均衡】
jaeger		【链路追踪】
sentinel	【限流、熔断、降级】
apisix		【网关、限流、熔断】
```

扩展模块
```
es		【搜索】
分布式锁	        【redis实现】
redis           【alone, cluster, cluster三种模式】
幂等性		【grpc重试，分布式锁处理，token处理等】
amqp            【amqp，消息队列，比如：rabbitmq】
cron            【分布式定时任务;go:go-cron,java:xxl-job】
分布式事务	【方式1：rocketmq，事务消息方式；方式2：seata-golang,暂缺】
分布式mysql	【go: gaea分库分表; java: shardingsphere-proxy】
OSS             【本地存储，阿里云存储，腾讯云存储，七牛存储】
SMS             【阿里云 sms】
zip             【压缩和解压】
xml2map         【xml2map, map2xml】
pay             【微信支付：小程序支付、jsapi支付,app支付,native,h5支付;支付宝：手机网站支付,app支付,电脑网站支付】
```

## 版本说明

```
v1.1.x ~ v.1.2.X 完成
已经实现了模块介绍中的功能
```

```
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
nacos, consul注册中心合并，可以在配置中任意切换
apisix,网关

增加了cicd自动化发布脚本,docker打包脚本等
这里使用的是：drone做的cicd，类似jenkins，但是比jenkins轻便很多。
代码仓库是： gogs; docker仓库的是:harbor
```

```
v1.2.9 当前版本
```

```
v1.3.X
会增加k8s的自动化发布脚本，Prometheus、Grafana监控等

v1.4.X 开发中
cicd，自动化发布到k8s中
```

```
v2.0 规划
会改造到istio或者dapr的三代微服务方式中，会单独另起一个项目进行维护
```



## 快速开始

安装 cli

```shell
git clone git@github.com:jettjia/go-micro-frame-cli.git
cd go-micro-frame-cli 

go build

go-micro-frame-cli install
```

初始化项目快速开始业务

```shell
$ go-micro-frame-cli init you_project_name
```

更多命令和功能

已支持：一键安装 mysql, redis, rabbitmq, es, go的开发环境

安装微服务需要的环境：nacos, jaeger,konga...

安装cicd自动化运维环境：gogs, drone, harbor

一键生成cicd运行的.drone.yaml配置模板

一键生成Dockerfile的运行模板

未来补充：k8s...

```shell
$ go-micro-frame-cli
Usage:
  go-micro-frame-cli [command]

Available Commands:
  build       build go project
  completion  generate the autocompletion script for the specified shell
  docker      create a docker image for current project
  drone       create .drone for ci/cd
  env         Print go-micro-frame version and environment info
  gen         automatically generate go files for ORM model,service, repository, handler, pb
  gofmt       gofmt your project
  help        Help about any command
  init        create and initialize an empty project
  install     install gf binary to system (might need root/admin permission)
  run         Install common service, like go-micro-frame-cli run mysql
  version     Show current binary version info

Flags:
      --config string   config file (default is $HOME/.go-micro-frame-cli.yaml)
  -h, --help            help for go-micro-frame-cli
  -t, --toggle          Help message for toggle

Use "go-micro-frame-cli [command] --help" for more information about a command.
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

 drone

![](https://cdn.nlark.com/yuque/0/2021/png/12759381/1635637765185-assets/web-upload/61a7da72-aef9-489f-84b5-50ae0eacb411.png)

harbor 

![](https://cdn.nlark.com/yuque/0/2021/png/12759381/1635637765209-assets/web-upload/98576ffd-9f1c-4dff-ad2c-3be7ef1717f2.png)

gogs

![](https://cdn.nlark.com/yuque/0/2021/png/12759381/1635637765170-assets/web-upload/7edd61a4-598e-407e-bb5e-bfc995c243c6.png)

## 项目演示

### 启动 grpc 服务

1. nacos后台配置

    配置内容参考, grpc里的配置示例

    ![https://cdn.nlark.com/yuque/0/2021/png/12759381/1635637765255-assets/web-upload/e26d6e6d-1714-46a0-abe4-e2e64bef9746.png](https://cdn.nlark.com/yuque/0/2021/png/12759381/1635637765255-assets/web-upload/e26d6e6d-1714-46a0-abe4-e2e64bef9746.png)

2. 启动grpc项目，会从nacos读取配置

```
[root@localhost grpc]# go run main.go
2021-09-21T15:38:26.457+0800    INFO    nacos_client/nacos_client.go:87 logDir:<tmp/nacos/log>   cacheDir:<tmp/nacos/cache>
2021-09-21T15:38:27.955+0800    INFO    nacos/nacos.go:26       从nacos读取到的全部配置如下：%!(EXTRA string={

```

​	consul会注册服务，这里可以启动多个grpc服务，已经实现了负载均衡的获取服务。

  ![](https://cdn.nlark.com/yuque/0/2021/png/12759381/1635637765137-assets/web-upload/9ec00855-6ff3-46bc-911a-1f479155531e.png)



### 启动 web 服务

1. 启动web项目，会从nacos读取配置

   ```
   [root@localhost web-gin]# go run main.go
   2021-09-21T16:37:38.927+0800    INFO    nacos_client/nacos_client.go:87 logDir:<tmp/nacos/log>   cacheDir:<tmp/nacos/cache>
   2021-09-21T16:37:39.118+0800    INFO    nacos/nacos.go:26       从nacos读取到的全部配置如下：%!(EXTRA string={
   
   ```

2. 这里会从consul中负载均衡的获取到服务，也会把web注入到consul中。这样可以用nginx或者kong等来读web进行负载均衡

   

3. 请求用postman 访问web层的接口

      ![](https://cdn.nlark.com/yuque/0/2021/png/12759381/1635637765258-assets/web-upload/6caee327-0679-4c51-a287-86cecff7b3e9.png)

4. 链路追踪

​       ![](https://cdn.nlark.com/yuque/0/2021/png/12759381/1635637765266-assets/web-upload/d4e43375-7ef2-4e42-b089-608710caf671.png)



### 网关 熔断 限流降级
apisix

