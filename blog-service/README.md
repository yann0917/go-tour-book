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

### 2.7 模块开发：文章管理

### 2.8 上传图片和文件服务

### 2.9 API访问控制

### 2.10 常见应用中间件

### 2.11 链路追踪

### 2.12 应用配置问题

### 2.13 编译程序应用

### 2.14 优雅重启和停止

### 2.15 思考

---
