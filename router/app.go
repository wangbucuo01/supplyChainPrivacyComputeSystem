package router

import (
	"supplyChainPrivacyComputeSystem/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.GET("/", service.Hello)
	// 上传文件
	r.POST("/supplychain/file/upload", service.Upload)

	// 用户模块
	r.POST("/supplychain/user/create", service.CreateUser)
	r.POST("/supplychain/user/login", service.Login)

	// 数据模块
	// 用户上传数据->数据处理（加密）->数据存储
	r.POST("/supplychain/data/create", service.CreateData)
	// 查询数据列表（根据Uid、日期）
	r.GET("/supplychain/data/list", service.DataList)
	// 数据加密计算(按照日期)
	r.POST("/supplychain/data/compute", service.ComputeData)
	// 数据验证（根据DataID）
	r.POST("/supplychain/data/verify", service.DataVerify)

	return r
}
