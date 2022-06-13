#### 项目结构
- application.go : 项目入口
- config.yaml : 配置文件
- go.mod: 项目依赖
- handler: 业务逻辑处理文件夹
- config : 配置读取 .yaml 配置文件
- logger : 日志配置
- mysql : 数据库配置
- redis : 缓存配置
- router : 路由配置
- wrapper :出入参数封装



#### 项目创建流程

##### 1、创建流程

- 创建 go.mod

##### 	cd 到目录文件下面

````
go mod init
````

- 创建配置文件  config.yaml

- 创建配置文件夹和解析go文件

  ````
  config\config.go
  ````

  

##### 2、项目依赖

- gin : Web容器
- graceful: Gin间接依赖，负责端口监听
- pflag : 替换原生 flag
- viper : 配置库 用于读取配置文件
- fsnotify: 文件系统监控工具
- gorm : 数据库 ORM 关系映射
- redis : Redis 配置与 CURD 操作
- zap : 日志管理
- govalidator: 参数校验

##### 3、配置

###### 3.1. yaml

----

`config.go` 读取配置文件需要符合以下几种需求：

1. 无任何启动参数设定默认文件。
2. 不同应用环境可以切换不同文件。
3. 项目运行时可动态修改配置。

###### 3.2. zap

----

  日志配置通常可以满足：按日期切割、按类别切割、保存周期等等。Go标准库包含Log相关操作，但在上述需求上还是有所欠缺，所以选择了第三方库替代原生Log，技术选型方面除了 `zap` 也配置 `logrus` 但最终也未实现想要的效果，果断放弃~

 



##### 4.启动类与配置文件

application.go

 首先使配置文件生效，然后配置MySQL、Redis、Log，最后启动Web容器，并监听端口。启动之后可以做健康检查，系统崩溃警报等等操作，简单的Golang Web结束~