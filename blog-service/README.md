# 博客之旅

> 源码[blog-service](https://github.com/go-programming-tour-book/blog-service)
>
## 技术选型

* web 框架 [`gin`](https://github.com/gin-gonic/gin), [gin-example](https://github.com/eddycjy/go-gin-example)

### 2.2 项目设计

> 对项目的目录结构、接口方案、路由注册、数据库等内容进行设计和开发。

* 目录结构参考，[Go 应用程序项目的基本布局](https://github.com/golang-standards/project-layout)
* 数据表
  * 文章表
  * 标签表
  * 文章标签关联表
* 接口选用 RESTful API, 参考[阮一峰 - RESTful API 设计指南](http://www.ruanyifeng.com/blog/2014/05/restful_api.html)

```sql
CREATE DATABASE
IF
    NOT EXISTS blog_service DEFAULT CHARACTER
    SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;

CREATE TABLE `blog_tag` (
`id` int(10) unsigned NOT NULL AUTO_INCREMENT,
`name` varchar(100) DEFAULT '' COMMENT '标签名称',
`state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
`created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
`created_by` varchar(100) DEFAULT '' COMMENT '创建人',`modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
`modified_by` varchar(100) DEFAULT '' COMMENT '修改人',`deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
`is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除 0为未删除、1为已删除',PRIMARY KEY (`id`))
ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='标签管理';

CREATE TABLE `blog_article` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(100) DEFAULT '' COMMENT '文章标题',
  `desc` varchar(255) DEFAULT'' COMMENT '文章简述',
  `cover_image_url` varchar(255) DEFAULT '' COMMENT '封面图片地址',
  `content` longtext COMMENT '文章内容',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT '新建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除： 0为未删除、1为已删除',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态：0为禁用、1为启用',  PRIMARY KEY (`id`))
  ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章管理';

CREATE TABLE `blog_article_tag` (
`id` int(10) unsigned NOT NULL AUTO_INCREMENT,
`article_id` int(11) NOT NULL COMMENT '文章ID',
`tag_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '标签ID',
`created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
`created_by` varchar(100) DEFAULT '' COMMENT '创建人',
`modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
`modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
`deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
`is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除： 0为未删除、1为已删除',
PRIMARY KEY (`id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章标签关联';

```

### 2.3 公共组件

* 错误码标准化
* 配置管理,使用[viper](https://github.com/spf13/viper)
* 数据库连接，使用[gorm-v1](https://github.com/go-gorm/gorm)
* 日志写入，使用[lumberjack](https://github.com/natefinch/lumberjack)
* 响应处理

### 2.4 接口文档

* 使用 [gin-swagger](https://github.com/swaggo/gin-swagger) 作为 API 文档框架

### 2.5 接口校验

### 2.6 模块开发：标签管理

### 2.7 上传图片和文件服务

* 文件服务使用 `gin.StaticFS` 提供静态资源站点

### 2.8 API访问控制

* [jwt-go](https://github.com/dgrijalva/jwt-go) Golang implementation of JSON Web Tokens (JWT)
* gin JWT 中间件

```sql
CREATE TABLE `blog_auth` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `app_key` varchar(20) DEFAULT '' COMMENT 'key',
  `app_secret` varchar(50) DEFAULT '' COMMENT 'secret',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除 0为未删除、1为已删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='jwt认证管理';

-- insert
INSERT INTO `blog_auth`(`id`, `app_key`, `app_secret`,`created_on`, `created_by`, `modified_on`, `modified_by`, `deleted_on`,`is_del`) 
VALUES (1, 'eddycjy', 'go-programming-tour-book', 0, 'eddycjy', 0,'', 0, 0);

```

### 2.9 常见应用中间件

* [gomail](https://github.com/go-gomail/gomail) SMTP服务发送电子邮件库
* [ratelimit](https://github.com/juju/ratelimit) 令牌桶实现的限流中间件

### 2.10 链路追踪

* [jaeger](https://www.jaegertracing.io/docs/1.22/)
* [opentracing-go](https://github.com/opentracing/opentracing-go)
* [jaeger-client-go](https://github.com/jaegertracing/jaeger-client-go)
* [opentracing-gorm](https://github.com/eddycjy/opentracing-gorm) SQL 追踪

### 2.11 应用配置问题

1. 可使用命令行参数写入配置
2. 写入系统环境变量, 使用 `os.Getenv("xxx")` 读取
3. 可借助第三方库将配置文件打包至二进制文件里，如 [go-bindata](https://github.com/go-bindata/go-bindata) ， 但是会导致文件增大
4. 集中式配置中心

* 配置热更新 [fsnotify](https://github.com/fsnotify/fsnotify)

### 2.12 编译程序应用

* Go语言的编译器默认支持并发编译和编译缓存，能够明显提升编译效率
* 缩小编译产生的二进制文件，如非必要情况，不建议压缩
    1. `go build -ldflags="-w -s"` 去掉DWARF调试信息和符号表信息
    2. 使用 [upx](https://github.com/upx/upx) 工具

### 2.13 优雅重启和停止

### 2.14 思考

---
