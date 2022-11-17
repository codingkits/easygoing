package demo

const (
	ProjectReadmeStr = `<h1 align="center">项目%s总览</h1>

## 项目介绍
* 【导读】：%s v1.0
* 项目工程示例

## 功能列表
* 项目工程示例

## 服务拆分
* Api
    * http层api服务
* Mq
    * 消息系统
* Auth
    * 鉴权
* Gateway
    * 网关系统

## 项目结构
>
>   ├── data                    # 数据挂载目录
>   ├── deploy                  # 部署
>   ├── docs                    # markdown
>   ├── service
>   │   ├── auth                # 鉴权服务
>   │   ├── api                 # HTTP
>   │   ├── mq                  # 消息
>   │   ├── gateway             # 网关服务
>   └── test

## 常见服务类型的目录结构
> xxsrv
>    ├── api         # http服务，业务需求实现
>    ├── task        # 定时任务，处理数据更新业务等
>    ├── model       # 数据操作
>    ├── mq          # 消息系统
>    ├── rpc         # rpc服务，给其他子系统提供基础数据访问
>    └── script      # 脚本，处理一些临时运营需求，临时数据修复

## 服务内部分层
> api/rpc/mq
>    ├── etc             # 服务配置
>    ├── internal
>    │   ├── config
>    │   ├── logic       # 业务逻辑
>    │   ├── server
>    │   └── svc         
>    └── pb              # pb文件

#### 项目参与者
- 版权所属：@copyright %s
- 维护者：@author %s  %s
- 项目创建日期 :  @date %s
- 其他：略。
	`
	GuideStr = `<h1 align="center">操作使用说明</h1>`

	ReadmeStr = `<h1 align="center">xxxx</h1>

## 服务介绍
* 

#### 项目参与者
- 维护者：
- 整理日期 :  @date 
- 其他：略。
`

	ApiStr = `package main

import (
	"%s/services/api/routers"
)

func main() {
	r := routers.SetRouter()
	r.Run(":2022")
}
`

	CmdStr = `package cmd`

	ControllerStr = `package controllers
`

	RouterStr = `package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetRouter() *gin.Engine {
	engine := gin.Default()
	engine.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "%s\n", "Hello Easygoing !")
		return
	})
	return engine
}
`

	ModelStr = `package models
`

	GoModuleStr = `module %s

go %s
`
)
